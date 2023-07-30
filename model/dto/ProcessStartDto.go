package dto

type ProcessStartDto struct {
	TemplateGroupId int
	Form            *map[string]any
	StartUserId     string
}
