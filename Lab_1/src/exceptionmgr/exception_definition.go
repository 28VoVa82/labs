package exceptionmgr

const (
	LessThanZeroT = iota
	ZeroExceptionsT
	GreaterThenTenT
)

type ErrMgr interface {
	GetErrorTypeStr() string
	GetErrorType() int
	isCritical() bool
}

type LessThanZero struct{}

func (e *LessThanZero) Error() string {
	return "Number is less than zero."
}

func (e *LessThanZero) isCritical() bool {
	return true
}

func (e *LessThanZero) GetErrorTypeStr() string {
	return "LessThanZero Error"
}

func (e *LessThanZero) GetErrorType() int {
	return LessThanZeroT
}

type GreaterThenTen struct{}

func (e *GreaterThenTen) Error() string {
	return "Number is greater then ten."
}

func (e *GreaterThenTen) isCritical() bool {
	return false
}

func (e *GreaterThenTen) GetErrorTypeStr() string {
	return "GreaterThenTen Error"
}

func (e *GreaterThenTen) GetErrorType() int {
	return GreaterThenTenT
}

type ZeroException struct{}

func (e *ZeroException) Error() string {
	return "Number is equal to zero."
}

func (e *ZeroException) isCritical() bool {
	return true
}

func (e *ZeroException) GetErrorTypeStr() string {
	return "ZeroException Error"
}

func (e *ZeroException) GetErrorType() int {
	return ZeroExceptionsT
}