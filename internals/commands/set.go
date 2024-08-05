package commands

import (
	"github.com/Anv3sh/Kioku/internals/storage"
	"github.com/Anv3sh/Kioku/internals/config"
	"github.com/Anv3sh/Kioku/internals/types"
	"sync"
	
)

func storekey(k *types.Kioku, dict *storage.Dict, node *storage.Node, key string){
	k.Mut.Lock()
	defer k.Mut.Unlock()
	dict.Store[key] = node
}

func SetCommand(args []string, k *types.Kioku ,dict *storage.Dict, lfu *storage.LFU, lru *storage.LRU, config config.Config) []byte {
	var wg sync.WaitGroup
	node := storage.CreateNode(args[1],args[2])
	// eviction on the basis of cache size
	// if len(lfu.MinHeap)==cap(lfu.MinHeap){
	// 	log.Println("Eviction in process...")
	// 	lfu.DeleteLF()
	// }
	//eviction on the basis of memory used by dict and suitable eviction policy
	wg.Add(3)
	go dict.EvictKey(k, lfu, lru)
	wg.Done()
	go storekey(k,dict,node, args[1])
	wg.Done()
	if lfu.Eviction{
		go lfu.Insert(k, node)
	}else if lru.Eviction{
		go lru.Insert(k, node)
	}
	wg.Done()
	wg.Wait()
	node = nil
	return []byte("OK\n")
}
