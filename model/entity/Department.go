package entity

type Department struct {
	Id       int `gorm:"primarykey"`
	ParentId int
	Name     string
}

func (e Department) TableName() string {
	return "department"
}

func (e Department) GetId() any {
	return e.Id
}
