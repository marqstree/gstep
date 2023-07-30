package ServerError

import (
	"github.com/marqstree/gstep/util/CONSTANT"
)

type ServerError struct {
	Code int    `convert:"code"`
	Msg  string `convert:"msg"`
}

func New(msg string) *ServerError {
	return &ServerError{
		Code: CONSTANT.FAIL_CODE,
		Msg:  msg,
	}
}

func NewByCode(code int, msg string) *ServerError {
	return &ServerError{
		Code: code,
		Msg:  msg,
	}
}

func (e *ServerError) Error() string {
	return e.Msg
}

//type ServerErrorData struct {
//	Code int    `convert:"code"`
//	Msg  string `convert:"msg"`
//}
//
//// 将自定义error转为对象
//func (e *ServerError) ToData() *ServerErrorData {
//	return &ServerErrorData{
//		Code: e.Code,
//		Msg:  e.Msg,
//	}
//}
