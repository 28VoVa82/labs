package main

import (
	"Lab_2/src/cfg"
	"Lab_2/src/exceptionmgr"
	"Lab_2/src/exchange"
	"fmt"
	"log"
	"math/rand"
)

var configFilePath = "./config/cli_cfg.toml"
var cliNetWorker exchange.NetWorker

func main() {
	cliCfg, cfgParseRes := cfg.ParseCliCfg(configFilePath)
	if cfgParseRes == false {
		log.Printf("Error while parsing client config file %s", configFilePath)
		return
	}

	sender := exchange.NewNetWorker(&cliCfg, &exchange.DefaultNetworkComm{})
	errCnt := exceptionmgr.NewErrorCounter(sender, cliCfg.CriticalErrorsList)

	sum := 0
	for i := 0; i < 200; i++ {
		genNum := rand.Intn(28) - 5
		newNum, err := modifyValue(genNum)
		if errCnt.CheckError(err) {
			errCnt.SendException(err)
			continue
		}

		sum += newNum
	}

	fmt.Printf("Total sum %d\n", sum)
	fmt.Printf("Critical error count: %d\n", errCnt.Critical)
	fmt.Printf("Ordinary error count: %d\n", errCnt.Ordinary)
	fmt.Printf("Server sending failes error count: %d\n", errCnt.BadSending)
}

func modifyValue(val int) (int, error) {
	if val < 0 {
		return 0, &exceptionmgr.LessThanZero{}
	}
	if val == 0 {
		return 0, &exceptionmgr.ZeroException{}
	}
	if val > 10 {
		return 0, &exceptionmgr.GreaterThenTen{}
	}

	return val * 2, nil
}
