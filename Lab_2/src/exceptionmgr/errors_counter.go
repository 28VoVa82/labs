package exceptionmgr

import "Lab_2/src/exchange"

func NewErrorCounter(sender *exchange.NetWorker, criticalErrors []string) *ErrorCounter {
	return &ErrorCounter{
		Sender:				sender,
		CriticalErrorList:	criticalErrors,
		Critical:			0,
		Ordinary:			0,
		BadSending:			0,
	}
}

func (errCnt *ErrorCounter) CheckError(err error) bool {
	if err == nil {
		return false
	}

	errMgr, isOk := err.(ErrMgr)
	if !isOk {
		return false
	}

	if errCnt.isCritical(errMgr) {
		errCnt.Critical++
	} else {
		errCnt.Ordinary++
	}

	return true
}

func (errCnt *ErrorCounter) isCritical(mgr ErrMgr) bool {
	actualErrType := mgr.GetErrorType()
	for _, errType := range errCnt.CriticalErrorList {
		if errType == actualErrType {
			return true
		}
	}

	return false
}

