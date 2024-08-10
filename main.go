package main

import (
	"log"
	// "time"
	// "net"
	"github.com/Anv3sh/Kioku/internals/constants"
	"github.com/Anv3sh/Kioku/internals/services"
	"github.com/Anv3sh/Kioku/internals/services/cmdutils"
	"github.com/Anv3sh/Kioku/internals/assests"

)

func main() {
	assests.PrintLogo()
	go cmdutils.CommandRegistry(&constants.REGCMDS, constants.COMMAND_LIST_PATH)
	constants.CONFIG.CreateConfig() // to create default config
	constants.CONFIG.SetConfig() // to set custom config settings
	constants.DICTIONARY.CreateDict(constants.CONFIG)
	constants.LFU_CACHE.CreateLFU(constants.CONFIG)
	kioku := services.NewKioku()
	aoffile,aof:=services.AOFFile(constants.CONFIG.AOFPolicy,constants.CONFIG.AOFPath)
	if aof{
		services.InMemSync(aoffile, &kioku, &constants.REGCMDS, &constants.DICTIONARY, &constants.LFU_CACHE, &constants.LRU_CACHE, constants.CONFIG)
		go services.AOFDiskWrite(aoffile,constants.CONFIG.AOFPolicy,&kioku)

	}
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
	log.Fatal(services.StartListening(&kioku))
}
