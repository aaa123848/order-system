package error

import "runtime"

type StdError struct {
	CallerFile     string
	CallerFuncLine int
	FuncName       string
	Line           int
	File           string
	Err            error
}

func (e StdError) Error() string {
	return e.Err.Error()
}

func GetStdError(err error) StdError {
	res := StdError{}
	p, f, l, _ := runtime.Caller(2)
	res.FuncName = runtime.FuncForPC(p).Name()
	res.File = f
	res.Line = l
	_, f, l, _ = runtime.Caller(3)
	res.CallerFile = f
	res.CallerFuncLine = l
	res.Err = err
	return res
}
