package main

import (
	// "fmt"
	"log"
	// "net"
	"github.com/Anv3sh/Kioku/pkg/services"
)

func main(){
	config := map[string]string{
		"HOST": "localhost",
		"PORT": "6379",
	}
	kioku := services.NewKioku(config)
	log.Fatal(kioku.StartListening())
	log.Printf("The Cache Vault is started listening on: \nport= %s host=%s", kioku.Port,kioku.Host)
}