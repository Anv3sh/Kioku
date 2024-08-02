package constants

import (
	"github.com/Anv3sh/Kioku/internals/config"
	"github.com/Anv3sh/Kioku/internals/services/cmdutils"
	"github.com/Anv3sh/Kioku/internals/storage"
)

const (
	ULIMIT            = 4096
	COMMAND_LIST_PATH = "./internals/commands/cmdlist.json"
)

var (
	CONFIG    config.Config
	REGCMDS   cmdutils.RegisteredCommands
	LFU_CACHE storage.LFU
)
