package main

import (
	"Labs_1/src/exceptionmgr"
	"fmt"
	"math/rand"
)

func main() {
	sum := 0
	for i := 0; i < 200; i++ {
		genNum := rand.Intn(28) - 5
		newNum, err := modifyValue(genNum)
		if exceptionmgr.CheckError(err) {
			//log.Print(err.Error())
			continue
		}

		sum += newNum
	}

	fmt.Printf("Total sum %d\n", sum)
	criticalCnt, standardCnt := exceptionmgr.GetErrorCounters()
	fmt.Printf("Critical error count: %d\n", criticalCnt)
	fmt.Printf("Standard error count: %d\n", standardCnt)
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

