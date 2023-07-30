package entity

type CommonEntity interface {
	TableName() string
	GetId() any
}
