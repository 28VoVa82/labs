package exceptionmgr

import (
	"testing"
)

func TestNoErrors(i *testing.T)  {
	var errCnt = NewErrorCounter(nil,nil)
	res := errCnt.CheckError(nil)
	if res {
		i.Errorf("Expected `false` value. Actual value %t", res)
	}
}

func TestCriticalError(i *testing.T)  {
	var errCnt = NewErrorCounter(nil, []string{"ZeroException", "LessThanZero"})

	var err error = &ZeroException{}
	res := errCnt.CheckError(err)
	if !res {
		i.Errorf("ZeroExeption: Expected `true` value. Actual value %t", res)
	}
	if errCnt.Critical != 1 {
		i.Errorf("ZeroExeption: Expected `1` value. Actual value %d", errCnt.Critical)
	}

	err = &LessThanZero{}
	res = errCnt.CheckError(err)
	if !res {
		i.Errorf("LessThanZero: Expected `true` value. Actual value %t", res)
	}
	if errCnt.Critical != 2 {
		i.Errorf("LessThanZero: Expected `2` value. Actual value %d", errCnt.Critical)
	}
}

func TestOrdinaryError(i *testing.T)  {
	var errCnt = NewErrorCounter(nil, []string{"ZeroException"})

	var err error = &GreaterThenTen{}
	res := errCnt.CheckError(err)
	if !res {
		i.Errorf("ZeroExeption: Expected `true` value. Actual value %t", res)
	}
	if errCnt.Ordinary != 1 {
		i.Errorf("ZeroExeption: Expected `1` value. Actual value %d", errCnt.Ordinary)
	}
}
