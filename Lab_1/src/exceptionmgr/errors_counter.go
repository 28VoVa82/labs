package exceptionmgr

var (
	criticalErrorsCnt int = 0
	standardErrorsCnt int = 0
)

func CheckError(err error) bool {
	if err == nil {
		return false
	}

	errMgr, isOk := err.(ErrMgr)
	if !isOk {
		return false
	}

	if errMgr.isCritical() {
		criticalErrorsCnt++
	} else {
		standardErrorsCnt++
	}

	return true
}

func GetErrorCounters() (int, int) {
	return criticalErrorsCnt, standardErrorsCnt
}
