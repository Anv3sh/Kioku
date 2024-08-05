package storage

import(
	"time"
	"unsafe"
)
type Node struct{
	Key string
	Value string
	ValueArr []int64
	Freq int64
	LastUsedTime time.Time
	Next *Node
	Prev *Node
}

func CreateNode(key string, value string) *Node{
	return &Node{
		Key: key,
		Value: value,
		Freq: 1,
		LastUsedTime: time.Now(),
		Next: nil,
		Prev: nil,
	}
}


func (node *Node) getMemUsage() float64{
	//TODO: make the mem usage calculation accurate
	//approx mem usage
	totalsizeInBytes := int(unsafe.Sizeof(node.Key))+int(unsafe.Sizeof(node.Value))+int(unsafe.Sizeof(node.Freq))+int(unsafe.Sizeof(node.LastUsedTime))
	if node.Next!=nil{
		totalsizeInBytes = totalsizeInBytes*2
	}
	if node.Prev!=nil{
		totalsizeInBytes = totalsizeInBytes*2
	}
	return float64(totalsizeInBytes) / (1024 * 1024)
}

