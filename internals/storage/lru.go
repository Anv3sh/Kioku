package storage

import(
	"github.com/Anv3sh/Kioku/internals/config"
	"github.com/Anv3sh/Kioku/internals/types"

)


type LRU struct{
	Head *Node
	Eviction bool
}

func (l *LRU) CreateLRU(config *config.Config){
	l.Eviction=config.LRUEviction
}

func (l *LRU) Insert(k *types.Kioku, node *Node){
	k.Mut.Lock()
	defer k.Mut.Unlock()
	if l.Head==nil{
		l.Head=node
		return
	}
	curr := l.Head
	for curr.Next !=nil{
		curr = curr.Next
	}
	curr.Next=node
	node.Prev=curr
}

func (l *LRU) Evict(k *types.Kioku){
	l.Head=l.Head.Next
	l.Head.Prev=nil
}

func (l *LRU) GetMemUsage(){

}