package types

import(
	"net"
	"sync"
)

type Kioku struct {
	ServerHost     string
	ServerPort     string
	Ln             net.Listener
	Quitch         chan struct{}
	Maxconnections chan struct{} // to manage the max number of client connections
	Msgch          chan []byte
	Connch         chan net.Conn
	RWMut            sync.RWMutex //mutex to handle thread synchronization
	Mut				sync.Mutex	
	Opch			chan []string
}