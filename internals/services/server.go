package services

import (
	"fmt"
	"log"
	"net"

	// "log"
	"bufio"
	"strings"
	"sync"

	"github.com/Anv3sh/Kioku/internals/constants"
	"github.com/Anv3sh/Kioku/internals/services/cmdutils"
	"github.com/Anv3sh/Kioku/internals/storage"
	"time"
)

type Kioku struct {
	ServerHost     string
	ServerPort     string
	Ln             net.Listener
	quitch         chan struct{}
	maxconnections chan struct{} // to manage the max number of client connections
	Msgch          chan []byte
	Connch         chan net.Conn
	mut            sync.RWMutex //mutex to handle thread synchronization
}

func NewKioku() Kioku {
	return Kioku{
		ServerHost:     constants.CONFIG.ServerHost,
		ServerPort:     constants.CONFIG.ServerPort,
		quitch:         make(chan struct{}),
		maxconnections: make(chan struct{}, constants.ULIMIT),
		Msgch:          make(chan []byte, 50),
		Connch:         make(chan net.Conn),
	}
}

func (k *Kioku) StartListening() error {
	ln, err := net.Listen("tcp", k.ServerHost+":"+k.ServerPort)
	log.Println("Kioku started listening on-> " + k.ServerHost + ":" + k.ServerPort)
	if err != nil {
		return err
	}

	defer ln.Close()
	k.Ln = ln
	go k.acceptLoop()
	<-k.quitch
	close(k.Msgch)
	close(k.Connch)
	return nil
}

func (k *Kioku) acceptLoop() {
	for {
		conn, err := k.Ln.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			continue
		}
		// if reached max connections reject new connection else accept and start readloop
		select {
		case k.maxconnections <- struct{}{}:
			fmt.Println("Connected to:", conn.RemoteAddr())

			go k.readLoop(conn)
		default:
			conn.Close()
			fmt.Println("Connection limit reached. Rejecting new connection.")
		}

	}
}

func (k *Kioku) readLoop(conn net.Conn) {
	defer func() {
		log.Println("Disconnected from: " + conn.RemoteAddr().String())
		conn.Close()
		<-k.maxconnections
	}()
	lfu := storage.CreateLFU(constants.CONFIG)
	// buf := make([]byte, 2048)
	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	for {
		conn.Write([]byte(k.Ln.Addr().String() + "> \r\n"))
		argv, err := rw.ReadString('\n')
		if err != nil {
			fmt.Println("read error:", err)
			continue
		}
		strings.TrimSpace(argv)
		args := strings.Fields(argv)
		msg := cmdutils.CommandChecker(args, &constants.REGCMDS, &lfu)
		k.Connch <- conn
		k.Msgch <- msg
		time.Sleep(500 * time.Millisecond)
	}
}
