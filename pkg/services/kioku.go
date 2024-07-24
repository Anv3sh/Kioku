package services

import (
	"fmt"
	"net"
	// "log"
	"github.com/Anv3sh/Kioku/pkg/constants"
)

type Kioku struct{
	Host string
	Port string
	ln  net.Listener
	quitch  chan struct{}
	maxconnections chan struct{} // to manage the max number of client connections
}


func NewKioku(config map[string]string) Kioku{
	return Kioku{
		Host: config["HOST"],
		Port: config["PORT"],
		quitch: make(chan struct{}),
		maxconnections: make(chan struct{}, constants.ULIMIT),
	}
}

func (k *Kioku) StartListening() error{
	ln, err := net.Listen("tcp", k.Host+":"+k.Port)
	if err != nil {
		return err
	}

	defer ln.Close()
	k.ln = ln
	go k.acceptLoop()
	<-k.quitch
	return nil
}

func (k *Kioku) acceptLoop() {
	for{
		conn, err := k.ln.Accept()

		if err != nil{
			fmt.Println("accept error:", err)
			continue
		}
		// if reached max connections reject new connection else accept and start readloop
		select{
		case k.maxconnections<-struct{}{}:
			fmt.Println("Connected to:", conn.RemoteAddr())
			go k.readLoop(conn)
		default:
			conn.Close()
            fmt.Println("Connection limit reached. Rejecting new connection.")
		}

		
	}
}

func (k *Kioku) readLoop(conn net.Conn){
	defer func() {
		conn.Close()
		<-k.maxconnections
	}()

	buf := make([]byte, 2048)

	for{
		n,err:= conn.Read(buf)
		if err!=nil{
			fmt.Println("read error:", err)
			continue
		}
		msg:=buf[:n]
		fmt.Printf(string(msg))
	}
}