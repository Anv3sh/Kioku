package main

import (
	"fmt"
	"log"
	// "net"
	"github.com/Anv3sh/Kioku/internals/config"
	"github.com/Anv3sh/Kioku/internals/constants"
	"github.com/Anv3sh/Kioku/internals/services"
	"github.com/Anv3sh/Kioku/internals/services/cmdreg"
)

func main() {
	config.SetConfig(&constants.CONFIG)
	kioku := services.NewKioku()
	err := cmdreg.CommandRegistry(&constants.REGCMDS, constants.COMMAND_LIST_PATH)
	if err != nil {
		log.Fatal("Command registry failed.")
	}

	go func() {
		for msg := range kioku.Msgch {
			fmt.Println("recieved mssg from connection:", string(msg))
		}
	}()
	log.Fatal(kioku.StartListening())

	log.Printf("Kioku started listening on: \nport= %s host=%s", kioku.ServerPort, kioku.ServerHost)
}
