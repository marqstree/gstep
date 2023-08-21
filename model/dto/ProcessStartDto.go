package dto

type ProcessStartDto struct {
	TemplateId  int
	Form        *map[string]any
	StartUserId string
}
