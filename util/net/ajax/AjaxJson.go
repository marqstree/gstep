package util_net_ajax

import (
	"encoding/json"
	"fmt"
	util_bizerror "github.com/marqstree/gstep/util/bizerror"
	util_constant "github.com/marqstree/gstep/util/constant"
	"net/http"
)

type AjaxJson struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
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
		Code: util_constant.SUCESS_CODE,
		Msg:  "成功",
		Data: data,
	}
}

func Success() *AjaxJson {
	return &AjaxJson{
		Code: util_constant.SUCESS_CODE,
		Msg:  "成功",
	}
}

func FailByError(err error) *AjaxJson {
	switch e := err.(type) {
	case *util_bizerror.BizError:
		return &AjaxJson{
			Code: e.Code,
			Msg:  e.Msg,
		}
	default:
		return &AjaxJson{
			Code: util_constant.FAIL_CODE,
			Msg:  e.Error(),
		}
	}
}

func Fail(msg string) *AjaxJson {
	return &AjaxJson{
		Code: util_constant.FAIL_CODE,
		Msg:  msg,
	}
}

func ResponseAjaxJson(writer http.ResponseWriter, aj *AjaxJson) {
	str, _ := json.Marshal(*aj)
	fmt.Fprintf(writer, "%s", str)
}
