package commands

import (
	"github.com/Anv3sh/Kioku/internals/storage"
)

func GetCommand(args []string, lfu *storage.LFU) []byte {
	_, exists := lfu.Store[args[1]]
	if !exists {
		return []byte("Key does not exists or evicted.\n")
	}
	val := lfu.Store[args[1]].Value
	lfu.Store[args[1]].Freq++
	return []byte("\"" + val + "\"\n")
}
