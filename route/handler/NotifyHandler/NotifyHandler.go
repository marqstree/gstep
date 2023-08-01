package NotifyHandler

import (
	"github.com/marqstree/gstep/model/dto"
	"github.com/marqstree/gstep/util/JsonUtil"
	"github.com/marqstree/gstep/util/net/AjaxJson"
	"github.com/marqstree/gstep/util/net/RequestParsUtil"
	"log"
	"net/http"
)

func TaskStateChange(writer http.ResponseWriter, request *http.Request) {
	dto := dto.NotifyTaskStateChangeDto{}
	RequestParsUtil.Body2dto(request, &dto)

	log.Println("接收到通知消息")
	log.Println(JsonUtil.Obj2PrettyJson(dto))

	AjaxJson.Success().Response(writer)
}
