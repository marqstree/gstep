package util_bizerror

import util_constant "github.com/marqstree/gstep/util/constant"

type BizError struct {
	Code int    `convert:"code"`
	Msg  string `convert:"msg"`
}

func New(msg string) *BizError {
	return &BizError{
		Code: util_constant.FAIL_CODE,
		Msg:  msg,
	}
}

func NewByCode(code int, msg string) *BizError {
	return &BizError{
		Code: code,
		Msg:  msg,
	}
}

func (e *BizError) Error() string {
	return e.Msg
}

type BizErrorData struct {
	Code int    `convert:"code"`
	Msg  string `convert:"msg"`
}

// 将自定义error转为对象
func (e *BizError) Data() *BizErrorData {
	return &BizErrorData{
		Code: e.Code,
		Msg:  e.Msg,
	}
}
