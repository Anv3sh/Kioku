package services

import (
	"fmt"
	"net"
	// "log"
	"bufio"
	"github.com/Anv3sh/Kioku/internals/constants"
	"strings"
	"sync"
	// "github.com/Anv3sh/Kioku/internals/commands"
)

type Kioku struct {
	ServerHost     string
	ServerPort     string
	ln             net.Listener
	quitch         chan struct{}
	maxconnections chan struct{} // to manage the max number of client connections
	Msgch          chan []byte
	mut            sync.RWMutex //mutex to handle thread synchronization
}

func NewKioku() Kioku {
	return Kioku{
		ServerHost:     constants.CONFIG.ServerHost,
		ServerPort:     constants.CONFIG.ServerPort,
		quitch:         make(chan struct{}),
		maxconnections: make(chan struct{}, constants.ULIMIT),
		Msgch:          make(chan []byte, 10),
	}
}

func (k *Kioku) StartListening() error {
	ln, err := net.Listen("tcp", k.ServerHost+":"+k.ServerPort)
	if err != nil {
		return err
	}

	defer ln.Close()
	k.ln = ln
	go k.acceptLoop()
	<-k.quitch
	close(k.Msgch)
	return nil
}

func (k *Kioku) acceptLoop() {
	for {
		conn, err := k.ln.Accept()

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
		conn.Close()
		<-k.maxconnections
	}()
	// buf := make([]byte, 2048)
	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	for {
		conn.Write([]byte(k.ln.Addr().String() + "> \r\n"))
		argv, err := rw.ReadString('\n')
		if err != nil {
			fmt.Println("read error:", err)
			continue
		}
		strings.TrimSpace(argv)
		args := strings.Fields(argv)

		k.Msgch <- []byte(args[0])
	}
}
