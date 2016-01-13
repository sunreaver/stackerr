package stackerr

import (
	"fmt"
	"runtime"
	"strings"
)

type StackErr struct {
	Filename      string
	CallingMethod string
	Line          int
	ErrorMessage  string
	StackTrace    string
}

func New(err interface{}) *StackErr {
	var errMessage string
	switch t := err.(type) {
	case *StackErr:
		return t
	case string:
		errMessage = t
	case error:
		errMessage = t.Error()
	default:
		errMessage = fmt.Sprintf("%v", t)
	}
	stackErr := &StackErr{}

	stackErr.ErrorMessage = errMessage
	_, file, line, ok := runtime.Caller(1)
	if ok {
		stackErr.Line = line
		components := strings.Split(file, "/")
		stackErr.Filename = components[(len(components) - 1)]
	}

	const size = 1 << 12
	buf := make([]byte, size)
	n := runtime.Stack(buf, false)
	stackErr.StackTrace = string(buf[:n])

	return stackErr
}

func (this *StackErr) Error() string {
	return fmt.Sprintf("{%s:%d} %s", this.Filename, this.Line, this.ErrorMessage)
}

func (this *StackErr) Stack() string {
	return this.StackTrace
}

func (this *StackErr) Detail() string {
	return fmt.Sprintf("{%s:%d} %s\nStack Info:\n %s", this.Filename, this.Line, this.ErrorMessage, this.StackTrace)
}
