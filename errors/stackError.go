package error

import (
	"fmt"
	"runtime"
	"time"
)

type StackError struct {
	Stack string
	Err   error
	Time  time.Time
}

func (e StackError) Error() string {
	return fmt.Sprintf("[ Error ] time: %v \nmsg: %v\n ^^^^^^^^^^^^^^ \n%v ========================\n\n\n", e.Time, e.Err, e.Stack)
}

func GetStackError(err error) StackError {
	var buff [4096]byte
	n := runtime.Stack(buff[:], false)

	return StackError{
		Stack: string(buff[:n]),
		Err:   err,
		Time:  time.Now(),
	}
}
