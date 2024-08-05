package storage

import(
	"github.com/Anv3sh/Kioku/internals/config"
	"github.com/Anv3sh/Kioku/internals/types"

)

type LFU struct{
	MinHeap []*Node
	Eviction bool
}


// func (h *Heap) Heapify() error{

// }

func (l *LFU) CreateLFU(conf config.Config ){
	l.Eviction=conf.LFUEviction

}

func (l *LFU) Insert(k *types.Kioku, node *Node){
	k.Mut.Lock()
	defer k.Mut.Unlock()
	l.MinHeap = append(l.MinHeap, node)
	curr_idx := len(l.MinHeap)-1
	parent_idx:=int(curr_idx/2)
	for parent_idx>0 && l.MinHeap[curr_idx].Freq<l.MinHeap[parent_idx].Freq{
		child:=l.MinHeap[curr_idx]
		l.MinHeap[curr_idx]=l.MinHeap[parent_idx]
		l.MinHeap[parent_idx]=child
		curr_idx=parent_idx
		parent_idx=int(curr_idx/2)
	}
}

func(l *LFU) Evict(k *types.Kioku){
	l.MinHeap[0]=l.MinHeap[len(l.MinHeap)-1]
	l.MinHeap=l.MinHeap[:len(l.MinHeap)-1]
	curr_idx:=0
	for (len(l.MinHeap)>2 && (l.MinHeap[curr_idx].Freq>l.MinHeap[curr_idx*2].Freq || 
		l.MinHeap[curr_idx].Freq>l.MinHeap[(curr_idx*2)+1].Freq)){
			if(l.MinHeap[curr_idx].Freq>l.MinHeap[curr_idx*2].Freq){
				curr:=l.MinHeap[curr_idx]
				l.MinHeap[curr_idx]=l.MinHeap[curr_idx*2]
				l.MinHeap[curr_idx*2]=curr

				curr_idx=curr_idx*2
			}else{
				if(l.MinHeap[curr_idx].Freq>l.MinHeap[(curr_idx*2)+1].Freq){
					curr:=l.MinHeap[curr_idx]
					l.MinHeap[curr_idx]=l.MinHeap[(curr_idx*2)+1]
					l.MinHeap[(curr_idx*2)+1]=curr
	
					curr_idx=(curr_idx*2)+1
			}
		}
}
}

// func (l *LFU) GetMemUsage() float64{
// 	sizeInBytes := len(l.MinHeap) * int(l.MinHeap[0].GetMemUsage())
// 	return float64(sizeInBytes) / (1024 * 1024)
// }
//TTL:300sec

//root node min
//complete binary tree