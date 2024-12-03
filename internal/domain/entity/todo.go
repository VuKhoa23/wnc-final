package entity

type Todo struct {
	Id        int    `db:"id"`
	Content   string `db:"content"`
	Completed bool   `db:"completed"`
	UserId    int    `db:"user_id"`
}
