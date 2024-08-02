package cmdutils

import (
	"fmt"
	"strings"
	// "github.com/Anv3sh/Kioku/internals/commands"
)


func CommandChecker(args []string, regcmds *RegisteredCommands)[]byte{
	cmd, exists := regcmds.Cmds[strings.ToUpper(args[0])]
	if !exists{
		return []byte("Command not found.\n")
	}else if(len(args)-1!=cmd.TotalArgs){
		msg:=fmt.Sprintf("Takes %d arguments but %d were given.\n",cmd.TotalArgs,len(args)-1)
		return []byte(msg)	
	}else{
		return []byte("OK\n")
	}
	
}