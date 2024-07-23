package services

import (
	// "fmt"
	"fmt"
	"net"
)

type Vault struct{
	Host string
	Port string
	ln  net.Listener
	quitch  chan struct{}
}


func NewVault(config map[string]string) Vault{
	return Vault{
		Host: config["HOST"],
		Port: config["PORT"],
		quitch: make(chan struct{}, ULIMIT),
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
		go v.readLoop(conn)
	}
}

func (v *Vault) readLoop(conn net.Conn){
	defer conn.Close()
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