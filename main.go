package main

import (
	"fmt"
	// "log"
	// "net"
	"github.com/Anv3sh/Cache-Vault/app/service"
)

func main(){
	vault := service.Vault{}
	config := map[string]string{
		"HOST": "localhost",
		"PORT": "6932",
		"MACHINE": "Windows",
	}

	fmt.Println(vault.ConfigVault(&vault,config))
	fmt.Printf("The Cache Vault is runnnig on: \nport= %s host=%s machine=%s", vault.Port,vault.Host,vault.Machine)
}