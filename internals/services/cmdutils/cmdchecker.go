package cmdutils

import (
	"fmt"
	"github.com/Anv3sh/Kioku/internals/commands"
	"github.com/Anv3sh/Kioku/internals/storage"
	"strings"
)

type CmdFunction func([]string, *storage.LFU) []byte

const PING_FUNC = "PingCommand"
const SET_FUNC = "SetCommand"
const GET_FUNC = "GetCommand"

var CmdFunctions = map[string]CmdFunction{
	PING_FUNC: commands.PingCommand,
	SET_FUNC:  commands.SetCommand,
	GET_FUNC:  commands.GetCommand,
}

func CommandChecker(args []string, regcmds *RegisteredCommands, lfu *storage.LFU) []byte {
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
		msg := cmdfunc(args, lfu)
		return msg
	}

}
