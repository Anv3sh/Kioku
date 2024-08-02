package storage

import(
	"github.com/Anv3sh/Kioku/internals/config"
	"unsafe"
)

type LFU struct{
	Store map[string]*Node
	MinHeap []*Node
	MaxMem float64
	Eviction bool
}


// func (h *Heap) Heapify() error{

// }

func CreateLFU(conf config.Config ) LFU{
	return LFU{
		Store: make(map[string]*Node),
		MaxMem: conf.MaxMem,
		Eviction: conf.Eviction,
	}
}

func (l *LFU) Insert(node *Node){
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

func(l *LFU) DeleteLF(){
	delete(l.Store,l.MinHeap[0].Key)
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

func (l *LFU) Size() float64{
	sizeInBytes := len(l.MinHeap) * int(unsafe.Sizeof(l.MinHeap[0]))
	return float64(sizeInBytes) / (1024 * 1024)
}
//TTL:300sec

//root node min
//complete binary tree