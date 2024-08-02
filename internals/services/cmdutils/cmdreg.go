package cmdutils

import (
	"encoding/json"
	"log"
	"os"
	// "github.com/Anv3sh/Kioku/internals/commands"
)

type cmdFunction func(string)

type CmdDetails struct {
	Name      string   `json:"name"`
	Info      string   `json:"info"`
	TotalArgs int      `json:"total_arguments"`
	Args      []string `json:"arguments"`
	Function	cmdFunction `json:"funtion"`
}

type RegisteredCommands struct {
	Cmds map[string]CmdDetails `json:"commands"`
}

func CommandRegistry(regCmds *RegisteredCommands, cmdsListPath string){
	content, err := os.ReadFile(cmdsListPath)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	err = json.Unmarshal(content, &regCmds)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	log.Println("Command registry complete.")

}
