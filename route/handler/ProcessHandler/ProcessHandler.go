package ProcessHandler

import (
	"github.com/marqstree/gstep/model/dto"
	"github.com/marqstree/gstep/model/entity"
	"github.com/marqstree/gstep/service/ProcessService"
	"github.com/marqstree/gstep/util/db/DbUtil"
	"github.com/marqstree/gstep/util/db/dao"
	"github.com/marqstree/gstep/util/net/AjaxJson"
	"github.com/marqstree/gstep/util/net/RequestParsUtil"
	"net/http"
)

func Start(writer http.ResponseWriter, request *http.Request) {
	requestDto := dto.ProcessStartDto{}
	RequestParsUtil.Body2dto(request, &requestDto)

	tx := DbUtil.GetTx()
	dao.CheckById[entity.User](requestDto.StartUserId, tx)
	//创建流程及启动任务
	id := ProcessService.Start(&requestDto, tx)

	tx.Commit()

	AjaxJson.SuccessByData(id).Response(writer)
}
