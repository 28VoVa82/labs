package exceptionmgr

import "Lab_2/src/exchange"

type ErrMgr interface {
	GetErrorType() string
}

type LessThanZero struct{}
type GreaterThenTen struct{}
type ZeroException struct{}

type ErrorCounter struct {
	Sender* exchange.NetWorker
	CriticalErrorList []string
	Critical int
	Ordinary int
	BadSending int
}