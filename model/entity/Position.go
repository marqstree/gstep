package entity

type Position struct {
	Position string `gorm:"primarykey"`
}

func (e Position) TableName() string {
	return "user"
}

func (e Position) GetId() any {
	return e.Position
}
