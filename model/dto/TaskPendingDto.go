package dto

import PageDto "github.com/marqstree/gstep/util/db/page"

type TaskPendingDto struct {
	PageDto.PageDto
	UserId string
}
