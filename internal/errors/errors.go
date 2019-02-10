package errors

import "fmt"

//Error implements the error interface
type Error struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

func (e *Error) Error() string {
	if e == nil {
		return ""
	}
	return fmt.Sprintf("ERR_%.3d: %s", e.Code, e.Msg)
}
