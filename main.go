package main

import (
	"log"
	// "net"
	"github.com/Anv3sh/Kioku/internals/config"
	"github.com/Anv3sh/Kioku/internals/constants"
	"github.com/Anv3sh/Kioku/internals/services"
	"github.com/Anv3sh/Kioku/internals/services/cmdreg"
)


func main() {
	go cmdreg.CommandRegistry(&constants.REGCMDS, constants.COMMAND_LIST_PATH)
	config.SetConfig(&constants.CONFIG)
	kioku := services.NewKioku()

	go func() {
		for conn:= range kioku.Connch{
			for msg := range kioku.Msgch {
				conn.Write(msg)
			}
		}
		
	}()
	log.Fatal(kioku.StartListening())
}
