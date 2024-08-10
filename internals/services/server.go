package services

import (
	"fmt"
	"log"
	"net"
	"io"
	// "log"
	"bufio"
	"strings"
	
	"github.com/Anv3sh/Kioku/internals/types"
	"github.com/Anv3sh/Kioku/internals/constants"
	"github.com/Anv3sh/Kioku/internals/services/cmdutils"
	"time"
)



func NewKioku() types.Kioku {
	return types.Kioku{
		ServerHost:     constants.CONFIG.ServerHost,
		ServerPort:     constants.CONFIG.ServerPort,
		Quitch:         make(chan struct{}),
		Maxconnections: make(chan struct{}, constants.ULIMIT),
		Msgch:          make(chan []byte, 50),
		Connch:         make(chan net.Conn),
		Opch:           make(chan []string, 50),
	}
}

func StartListening(k *types.Kioku) error {
	ln, err := net.Listen("tcp", k.ServerHost+":"+k.ServerPort)
	log.Println("Kioku started listening on-> " + k.ServerHost + ":" + k.ServerPort)
	if err != nil {
		return err
	}

	defer ln.Close()
	k.Ln = ln
	go acceptLoop(k)
	<-k.Quitch
	close(k.Msgch)
	close(k.Connch)
	return nil
}

func acceptLoop(k *types.Kioku) {
	for {
		conn, err := k.Ln.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			continue
		}
		// if reached max connections reject new connection else accept and start readloop
		select {
		case k.Maxconnections <- struct{}{}:
			log.Println("Connected to:", conn.RemoteAddr())

			go readLoop(k,conn)
		default:
			conn.Close()
			log.Println("Connection limit reached. Rejecting new connection.")
		}

	}
}

func readLoop(k *types.Kioku, conn net.Conn) {
	defer func() {
		log.Println(conn.RemoteAddr().String()+" disconnected.")
		conn.Close()
		<-k.Maxconnections
	}()
	// buf := make([]byte, 2048)
	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	for {
		conn.Write([]byte(k.Ln.Addr().String() + "> \r\n"))
		argv, err := rw.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return
			}
			// fmt.Println("read error:", err)
			return
		}
		strings.TrimSpace(argv)
		args := strings.Fields(argv)
		msg,wrcmd := cmdutils.CommandChecker(args, k, &constants.REGCMDS, &constants.DICTIONARY, &constants.LFU_CACHE, &constants.LRU_CACHE, constants.CONFIG)
		if wrcmd{
			k.Opch <- args
		}
		k.Connch <- conn
		k.Msgch <- msg
		time.Sleep(500 * time.Millisecond)
	}
}
