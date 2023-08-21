package dto

import PageDto "github.com/marqstree/gstep/util/db/page"

type TemplateQueryDto struct {
	PageDto.PageDto
	VersionId int
}
