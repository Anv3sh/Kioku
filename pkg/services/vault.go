package services

import (
	"fmt"
	"net"
	// "log"
	"github.com/Anv3sh/Cache-Vault/pkg/constants"
)

type Vault struct{
	Host string
	Port string
	ln  net.Listener
	quitch  chan struct{}
	maxconnections chan struct{} // to manage the max number of client connections
}


func NewVault(config map[string]string) Vault{
	return Vault{
		Host: config["HOST"],
		Port: config["PORT"],
		quitch: make(chan struct{}),
		maxconnections: make(chan struct{}, constants.ULIMIT),
	}
}

func (v *Vault) StartListening() error{
	ln, err := net.Listen("tcp", v.Host+":"+v.Port)
	if err != nil {
		return err
	}

	defer ln.Close()
	v.ln = ln
	go v.acceptLoop()
	<-v.quitch
	return nil
}

func (v *Vault) acceptLoop() {
	for{
		conn, err := v.ln.Accept()

		if err != nil{
			fmt.Println("accept error:", err)
			continue
		}

		// if reached max connections reject new connection else accept and start readloop
		select{
		case v.maxconnections<-struct{}{}:
			go v.readLoop(conn)
		default:
			conn.Close()
            fmt.Println("Connection limit reached. Rejecting new connection.")
		}

		
	}
}

func (v *Vault) readLoop(conn net.Conn){
	defer func() {
		conn.Close()
		<-v.maxconnections
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