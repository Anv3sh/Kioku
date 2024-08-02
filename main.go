package main

import (
	"log"
	// "net"
	"github.com/Anv3sh/Kioku/internals/config"
	"github.com/Anv3sh/Kioku/internals/constants"
	"github.com/Anv3sh/Kioku/internals/services"
	"github.com/Anv3sh/Kioku/internals/services/cmdutils"
)


func main() {
	go cmdutils.CommandRegistry(&constants.REGCMDS, constants.COMMAND_LIST_PATH)
	config.SetConfig(&constants.CONFIG)
	kioku := services.NewKioku()

	go func() {
		for{
			conn:= <- kioku.Connch
			msg:= <- kioku.Msgch
			conn.Write(msg)
		}
		
	}()
	log.Fatal(kioku.StartListening())
}
