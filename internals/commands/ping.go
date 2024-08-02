package commands

import (
	"github.com/Anv3sh/Kioku/internals/storage"
)

func PingCommand(args []string, lfu *storage.LFU) []byte {
	return []byte("PONG!\n")
}
