package service

import (
	// "fmt"
)

type Vault struct{
	Host string
	Port string
	Machine string
}

// func (v Vault) StartService(vault Vault) string{
// 	fmt.Println("inside start service")
// 	return ("Vault session running...")
// }

func (v Vault) ConfigVault(vault *Vault, config map[string]string) string{
	vault.Host = config["HOST"]
	vault.Port = config["PORT"]
	vault.Machine = config["MACHINE"]

	return ("The vault is configured!!")
}