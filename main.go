package main

import (
	"fmt"
	"log"
	// "net"
	"github.com/Anv3sh/Kioku/internals/config"
	"github.com/Anv3sh/Kioku/internals/constants"
	"github.com/Anv3sh/Kioku/internals/services"
)

func main() {
	config.SetConfig(&constants.CONFIG)
	kioku := services.NewKioku()

	go func() {
		for msg := range kioku.Msgch {
			fmt.Println("recieved mssg from connection:", string(msg))
		}
	}()
	log.Fatal(kioku.StartListening())

	log.Printf("The Cache Vault is started listening on: \nport= %s host=%s", kioku.Port, kioku.Host)
}
