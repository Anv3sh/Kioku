package main

import (
	"fmt"
	// "log"
	// "net"
	"github.com/Anv3sh/Cache-Vault/pkg/services"
)

func main(){
	vault := services.Vault{}
	config := map[string]string{
		"HOST": "localhost",
		"PORT": "6932",
	}
	// vault.StartListening()
	fmt.Println(vault.ConfigVault(config))
	fmt.Printf("The Cache Vault is runnnig on: \nport= %s host=%s", vault.Port,vault.Host)
}