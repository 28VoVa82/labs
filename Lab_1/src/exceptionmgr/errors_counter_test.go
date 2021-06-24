package exceptionmgr

import "testing"

func TestNoErrors(i *testing.T)  {
	res := CheckError(nil)
	if res {
		i.Errorf("Expected `false` value. Actual value %t", res)
	}
}

func TestCriticalError(i *testing.T)  {
	criticalErrorsCnt = 0
	var err error = &ZeroException{}
	res := CheckError(err)
	if !res {
		i.Errorf("ZeroExeption: Expected `true` value. Actual value %t", res)
	}
	if criticalErrorsCnt != 1 {
		i.Errorf("ZeroExeption: Expected `1` value. Actual value %d", criticalErrorsCnt)
	}

	err = &LessThanZero{}
	res = CheckError(err)
	if !res {
		i.Errorf("LessThanZero: Expected `true` value. Actual value %t", res)
	}
	if criticalErrorsCnt != 2 {
		i.Errorf("LessThanZero: Expected `2` value. Actual value %d", criticalErrorsCnt)
	}
}

func TestStandardError(i *testing.T)  {
	standardErrorsCnt = 0
	var err error = &GreaterThenTen{}
	res := CheckError(err)
	if !res {
		i.Errorf("ZeroExeption: Expected `true` value. Actual value %t", res)
	}
	if standardErrorsCnt != 1 {
		i.Errorf("ZeroExeption: Expected `1` value. Actual value %d", standardErrorsCnt)
	}
}
