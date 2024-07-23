package services

import (
	// "fmt"
)

type Vault struct{
	Host string
	Port string
}

// func (v Vault) StartService(vault Vault) string{
// 	fmt.Println("inside start service")
// 	return ("Vault session running...")
// }

func (v *Vault) ConfigVault(config map[string]string) string{
	v.Host = config["HOST"]
	v.Port = config["PORT"]

	return ("The vault is configured!!")
}

// func (v Vault) StartListening() error{
	
// }