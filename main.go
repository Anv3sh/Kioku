package main

import (
	"log"
	// "time"
	// "net"
	"github.com/Anv3sh/Kioku/internals/config"
	"github.com/Anv3sh/Kioku/internals/constants"
	"github.com/Anv3sh/Kioku/internals/services"
	"github.com/Anv3sh/Kioku/internals/services/cmdutils"
)

func main() {
	go cmdutils.CommandRegistry(&constants.REGCMDS, constants.COMMAND_LIST_PATH)
	config.SetConfig(&constants.CONFIG)
	constants.LFU_CACHE.CreateLFU(constants.CONFIG)
	kioku := services.NewKioku()

	go func() {
		for {
			conn := <-kioku.Connch
			msg := <-kioku.Msgch
			conn.Write(msg)
		}

	}()
	//ttl logic:
	// if constants.CONFIG.Eviction{
	// 	go func(){
	// 		for{
	// 			if len(constants.LFU_CACHE.MinHeap)>0{
	// 			time.Sleep(constants.CONFIG.TotalTimetoLive*time.Second)
	// 			constants.LFU_CACHE.DeleteLF()
	// 			}
	// 		}
	// 	}()
	// }
	log.Fatal(kioku.StartListening())
}
