package main

import (
	"Labs_1/src/exceptionmgr"
	"testing"
)

func TestModifyValue1(t *testing.T) {
	_, err := modifyValue(16)
	errMgr, isOk := err.(exceptionmgr.ErrMgr)
	if !isOk {
		t.Error("`err` not implement `exceptionmgr.ErrMgr`")
	}

	if eType := errMgr.GetErrorType(); eType != exceptionmgr.GreaterThenTenT {
		t.Errorf("Exceptions expected type GreaterThenTen(%d). Actual one %d", exceptionmgr.GreaterThenTenT, eType)
	}
}

func TestModifyValue2(t *testing.T) {
	_, err := modifyValue(0)
	errMgr, isOk := err.(exceptionmgr.ErrMgr)
	if !isOk {
		t.Error("`err` not implement `exceptionmgr.ErrMgr`")
	}

	if eType := errMgr.GetErrorType(); eType != exceptionmgr.ZeroExceptionsT {
		t.Errorf("Exceptions expected type ZeroExceptions(%d). Actual one %d", exceptionmgr.ZeroExceptionsT, eType)
	}

}
func TestModifyValue3(t *testing.T) {
	_, err := modifyValue(-1)
	errMgr, isOk := err.(exceptionmgr.ErrMgr)
	if !isOk {
		t.Error("`err` not implement `exceptionmgr.ErrMgr`")
	}

	if eType := errMgr.GetErrorType(); eType != exceptionmgr.LessThanZeroT {
		t.Errorf("Exceptions expected type LessThanZero(%d). Actual one %d", exceptionmgr.LessThanZeroT, eType)
	}

}
func TestModifyValue4(t *testing.T) {
	var a int
	a, err := modifyValue(5)
	if err != nil || a != 10 {
		t.Error("Expected val 10, nil, got", a, err)
	}

}
