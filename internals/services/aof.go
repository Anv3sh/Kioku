package services

import (
	"bufio"
	"log"
	"os"
	"strings"
	"time"
	"github.com/Anv3sh/Kioku/internals/config"
	"github.com/Anv3sh/Kioku/internals/services/cmdutils"
	"github.com/Anv3sh/Kioku/internals/storage"
	"github.com/Anv3sh/Kioku/internals/types"
)

const NO="NO"
const EVERYSEC="EVERYSEC"
const ALWAYS="ALWAYS"


var AOF_POLICY_CONSTANTS = map[string]int{NO:0,EVERYSEC:1}


func AOFFile(policy string, aoffilepath string) (*os.File, bool){
	policychk:=AOF_POLICY_CONSTANTS[strings.ToUpper(policy)]
	if policychk==0{
		return nil,false
	}
	file,err:=os.OpenFile(aoffilepath, os.O_APPEND|os.O_CREATE,0644)
	if err!=nil{
		log.Fatal("Error in reading AOF:",err)
		return nil, false
	}
	return file,true
}

func InMemSync(f *os.File, k *types.Kioku, regcmds *cmdutils.RegisteredCommands, dict *storage.Dict, lfu *storage.LFU, lru *storage.LRU, config config.Config){
	scanner:=bufio.NewScanner(f)
	for scanner.Scan(){
		_, wrcmd:= cmdutils.CommandChecker(strings.Split(scanner.Text(), " "),k,regcmds,dict,lfu,lru,config)
			if wrcmd{
				continue
			}
	}

	log.Println("AOF sync complete.")
}

func AOFDiskWrite(file *os.File, policy string, k *types.Kioku){
	
	defer file.Close()
	for op:=range k.Opch{
		appendToFile(op,file)
		time.Sleep(time.Duration(1)*time.Second)
	}
	
}

func appendToFile(op []string, f *os.File){
	_, err:=f.WriteString(strings.Join(op," ")+"\n")
	if err!=nil{
		log.Fatal("Error writing to file:", err)
        return
	}
}