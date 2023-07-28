package model_dto

type TemplateStepDto struct {
	Id            int
	Title         string
	Category      string
	Form          *map[string]any
	Expression    string
	AgreeSteps    []*TemplateStepDto
	RefuseStepIds []int
}
