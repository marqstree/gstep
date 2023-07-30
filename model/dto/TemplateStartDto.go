package dto

type TemplateStartDto struct {
	ProcessId   int
	Form        *map[string]any `json:"form,omitempty"`
	startUserId string
}
