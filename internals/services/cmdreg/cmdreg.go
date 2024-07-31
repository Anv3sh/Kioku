package cmdreg

import (
	"encoding/json"
	"log"
	"os"
)

type Cmd struct {
	Name      string   `json:"name"`
	Info      string   `json:"info"`
	TotalArgs int      `json:"total_arguments"`
	Args      []string `json:"arguments"`
}

type RegisteredCommands struct {
	Cmds []Cmd `json:"commands"`
}

func CommandRegistry(regCmds *RegisteredCommands, cmdsListPath string) error {
	content, err := os.ReadFile(cmdsListPath)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
		return err
	}

	err = json.Unmarshal(content, &regCmds)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
		return err
	}

	return nil
}
