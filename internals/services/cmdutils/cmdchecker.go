package cmdutils

import (
	"fmt"
	"strings"

	"github.com/Anv3sh/Kioku/internals/commands"
	"github.com/Anv3sh/Kioku/internals/config"
	"github.com/Anv3sh/Kioku/internals/storage"
	"github.com/Anv3sh/Kioku/internals/types"

)

type CmdFunction func([]string, *types.Kioku, *storage.Dict, *storage.LFU, *storage.LRU, config.Config) []byte

const PING_FUNC = "PingCommand"
const SET_FUNC = "SetCommand"
const GET_FUNC = "GetCommand"

var CmdFunctions = map[string]CmdFunction{
	PING_FUNC: commands.PingCommand,
	SET_FUNC:  commands.SetCommand,
	GET_FUNC:  commands.GetCommand,
}

func CommandChecker(args []string,k *types.Kioku, regcmds *RegisteredCommands, dict *storage.Dict, lfu *storage.LFU, lru *storage.LRU, config config.Config) []byte {
	if len(args) == 0 {
		return []byte("")
	}
	cmd, exists := regcmds.Cmds[strings.ToUpper(args[0])]
	if !exists {
		return []byte("Command not found.\n")
	} else if len(args)-1 != cmd.TotalArgs {
		msg := fmt.Sprintf("Takes %d arguments but %d were given.\n", cmd.TotalArgs, len(args)-1)
		return []byte(msg)
	} else {
		cmdfunc := CmdFunctions[cmd.Function]
		msg := cmdfunc(args,k, dict, lfu, lru, config)
		return msg
	}

}
