package storage

import(
	"github.com/Anv3sh/Kioku/internals/config"
	"github.com/Anv3sh/Kioku/internals/types"

)

type Dict struct{
	Store map[string]*Node
	MaxMem float64
}

func (d *Dict) CreateDict(config config.Config){
	d.Store=make(map[string]*Node)
	d.MaxMem=config.MaxMem
}

func (d *Dict) EvictKey(k *types.Kioku, lfu *LFU, lru *LRU){
	k.Mut.Lock()
	defer k.Mut.Unlock()
	if lfu.Eviction && d.getMemUsageLFU(lfu)>d.MaxMem{
		delete(d.Store,lfu.MinHeap[0].Key)
		lfu.Evict(k)
	}else if lru.Eviction && d.getMemUsageLRU(lru)>d.MaxMem{
		delete(d.Store,lru.Head.Key)
		lru.Evict(k)
	}else{
		return
	}
}


func (d *Dict) getMemUsageLFU(lfu *LFU) float64{
	sizeInBytes:=0
	if len(lfu.MinHeap)>0{
	sizeInBytes = len(d.Store) * int(lfu.MinHeap[0].getMemUsage())
	}
	return float64(sizeInBytes) / (1024 * 1024)
}


func (d *Dict) getMemUsageLRU(lru *LRU) float64{
	sizeInBytes:=0
	if lru.Head!=nil{
		sizeInBytes = len(d.Store) * int(lru.Head.getMemUsage())
	}
	return float64(sizeInBytes) / (1024 * 1024)
}