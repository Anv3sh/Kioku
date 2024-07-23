package services

import (
	// "fmt"
)

type Vault struct{
	Host string
	Port string
}


func NewVault(config map[string]string) Vault{
	return Vault{
		Host: config["HOST"],
		Port: config["PORT"],
	}
}