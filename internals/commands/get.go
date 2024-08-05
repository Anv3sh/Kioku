package commands

import (
	"sync"

	"github.com/Anv3sh/Kioku/internals/config"
	"github.com/Anv3sh/Kioku/internals/storage"
	"github.com/Anv3sh/Kioku/internals/types"
)

func increasefreq(k *types.Kioku, key string, dict *storage.Dict){
	k.Mut.Lock()
	defer k.Mut.Unlock()
	dict.Store[key].Freq++
}

func GetCommand(args []string, k *types.Kioku, dict *storage.Dict, lfu *storage.LFU, lru *storage.LRU, config config.Config) []byte {
	var wg sync.WaitGroup
	_, exists := dict.Store[args[1]]
	if !exists {
		return []byte("Key does not exists or evicted.\n")
	}
	
	val := dict.Store[args[1]].Value
	wg.Add(1)
	go increasefreq(k ,args[1], dict)
	wg.Done()
	wg.Wait()
	return []byte("\"" + val + "\"\n")
}
