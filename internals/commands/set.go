package commands

import (
	"github.com/Anv3sh/Kioku/internals/storage"
)

func SetCommand(args []string, lfu *storage.LFU) []byte{
	node := &storage.Node{
		Key: args[1],
		Value: args[2],
		Freq: 1,
	}
	// eviction on the basis of cache size
	// if len(lfu.MinHeap)==cap(lfu.MinHeap){
	// 	log.Println("Eviction in process...")
	// 	lfu.DeleteLF()
	// }
	//eviction on the basis of memory used by cache
	if lfu.Size()>=lfu.MaxMem && lfu.Eviction{
		// log.Println("Eviction in process...")
		lfu.DeleteLF()
	}
	lfu.Store[args[1]]=node
	lfu.Insert(node)
	node=nil
	return []byte("OK\n")
}