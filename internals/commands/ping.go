package commands

import (
	"github.com/Anv3sh/Kioku/internals/storage"
	"github.com/Anv3sh/Kioku/internals/config"
	"github.com/Anv3sh/Kioku/internals/types"

)

func PingCommand(args []string,k *types.Kioku, dict *storage.Dict, lfu *storage.LFU, lru *storage.LRU, config config.Config) []byte {
	return []byte("PONG!\n")
}
