package entity

type Department struct {
	Id       string `json:"id" gorm:"primarykey"`
	ParentId string `json:"parentId"`
	Name     string `json:"name"`
}

func (e Department) TableName() string {
	return "department"
}

func (e Department) GetId() any {
	return e.Id
}
