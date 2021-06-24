package exceptionmgr

func (e *LessThanZero) Error() string {
	return "Number is less than zero."
}

func (e *LessThanZero) GetErrorType() string {
	return "LessThanZero"
}

func (e *GreaterThenTen) Error() string {
	return "Number is greater then ten."
}

func (e *GreaterThenTen) GetErrorType() string {
	return "GreaterThenTen"
}

func (e *ZeroException) Error() string {
	return "Number is equal to zero."
}

func (e *ZeroException) GetErrorType() string {
	return "ZeroException"
}
