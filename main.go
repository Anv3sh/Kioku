package main

import (
	// "fmt"
	"log"
	// "net"
	"github.com/Anv3sh/Cache-Vault/pkg/services"
)

func main(){
	config := map[string]string{
		"HOST": "localhost",
		"PORT": "6379",
	}
	vault := services.NewVault(config)
	log.Fatal(vault.StartListening())
	log.Printf("The Cache Vault is started listening on: \nport= %s host=%s", vault.Port,vault.Host)
}