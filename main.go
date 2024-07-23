package main

import (
	"fmt"
	// "log"
	// "net"
	"github.com/Anv3sh/Cache-Vault/pkg/services"
)

func main(){
	config := map[string]string{
		"HOST": "localhost",
		"PORT": "6932",
	}
	vault := services.NewVault(config)
	// vault.StartListening()
	fmt.Printf("The Cache Vault is runnnig on: \nport= %s host=%s", vault.Port,vault.Host)
}