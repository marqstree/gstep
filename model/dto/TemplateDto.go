package model_dto

type TemplateDto struct {
	Id       int
	GroupId  int
	Title    string
	RootStep *TemplateStepDto
}
