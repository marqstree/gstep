package AjaxJson

import (
	"encoding/json"
	"fmt"
	"github.com/marqstree/gstep/util/CONSTANT"
	"github.com/marqstree/gstep/util/ServerError"
	"net/http"
)

type AjaxJson struct {
	Code int    `JsonUtil:"code"`
	Msg  string `JsonUtil:"msg"`
	Data any    `JsonUtil:"data"`
}

func New(code int, msg string, data any) *AjaxJson {
	return &AjaxJson{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

func SuccessByData(data any) *AjaxJson {
	return &AjaxJson{
		Code: CONSTANT.SUCESS_CODE,
		Msg:  "成功",
		Data: data,
	}
}

func Success() *AjaxJson {
	return &AjaxJson{
		Code: CONSTANT.SUCESS_CODE,
		Msg:  "成功",
	}
}

func FailByError(err error) *AjaxJson {
	switch e := err.(type) {
	case *ServerError.ServerError:
		return &AjaxJson{
			Code: e.Code,
			Msg:  e.Msg,
		}
	default:
		return &AjaxJson{
			Code: CONSTANT.FAIL_CODE,
			Msg:  e.Error(),
		}
	}
}

func Fail(msg string) *AjaxJson {
	return &AjaxJson{
		Code: CONSTANT.FAIL_CODE,
		Msg:  msg,
	}
}

func (a *AjaxJson) Response(writer http.ResponseWriter) {
	str, _ := json.Marshal(*a)
	fmt.Fprintf(writer, "%s", str)
}
