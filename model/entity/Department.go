package entity

type Department struct {
	Id       int `gorm:"primarykey"`
	ParentId int
	Name     string
}
