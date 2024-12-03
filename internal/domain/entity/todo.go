package entity

type Todo struct {
	Id        int64  `gorm:"column:id;primaryKey" json:"id"`
	Content   string `gorm:"column:content" json:"content"`
	Completed bool   `gorm:"column:completed" json:"completed"`
	UserId    int64  `gorm:"column:user_id" json:"user_id"`
}

func (Todo) TableName() string {
	return "todos"
}
