package services

import (
	"fmt"
	"log"
	"net"

	// "log"
	"bufio"
	"strings"
	"sync"
	// "time"

	"github.com/Anv3sh/Kioku/internals/constants"
	// "github.com/Anv3sh/Kioku/internals/commands"
)

type Kioku struct {
	ServerHost     string
	ServerPort     string
	ln             net.Listener
	quitch         chan struct{}
	maxconnections chan struct{} // to manage the max number of client connections
	Msgch          chan []byte
	Connch		   chan net.Conn
	mut            sync.RWMutex //mutex to handle thread synchronization
}

func NewKioku() Kioku {
	return Kioku{
		ServerHost:     constants.CONFIG.ServerHost,
		ServerPort:     constants.CONFIG.ServerPort,
		quitch:         make(chan struct{}),
		maxconnections: make(chan struct{}, constants.ULIMIT),
		Msgch:          make(chan []byte, 10),
		Connch:         make(chan net.Conn),
	}
}

func (k *Kioku) StartListening() error {
	ln, err := net.Listen("tcp", k.ServerHost+":"+k.ServerPort)
	log.Println("Kioku started listening on: \nport= "+k.ServerPort+" host= "+k.ServerHost)
	if err != nil {
		return err
	}

	defer ln.Close()
	k.ln = ln
	go k.acceptLoop()
	<-k.quitch
	close(k.Msgch)
	close(k.Connch)
	return net.ErrClosed
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
		argv, err := rw.ReadString('\n')
		if err != nil {
			fmt.Println("read error:", err)
			continue
		}
		strings.TrimSpace(argv)
		args := strings.Fields(argv)
		_, exists := constants.REGCMDS.Cmds[strings.ToUpper(args[0])]
		k.Connch <- conn
		if !exists{
			k.Msgch <- []byte("No command found!\n"+k.ln.Addr().String() + "> \r\n")
		}else{
			k.Msgch <- []byte("OK\n"+k.ln.Addr().String() + "> \r\n")	
		}
		// time.Sleep(1*time.Second)
	}
}
