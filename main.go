package main

import (
	// "fmt"
	"fmt"
	"log"
	// "net"
	"github.com/Anv3sh/Kioku/internals/services"
)

func main() {
	config := map[string]string{
		"HOST": "localhost",
		"PORT": "6379",
	}
	kioku := services.NewKioku(config)
	go func() {
		for msg := range kioku.Msgch {
			fmt.Println("recieved mssg from connection:", string(msg))
		}
	}()
	log.Fatal(kioku.StartListening())

	log.Printf("The Cache Vault is started listening on: \nport= %s host=%s", kioku.Port, kioku.Host)
}
