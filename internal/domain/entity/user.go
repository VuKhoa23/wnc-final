package entity

// mapping for database

type User struct {
	Id           uint64 `db:"id"`
	Username     string `db:"username" json:"username"`
	Password     string `db:"password"`
	RefreshToken string `db:"refresh_token"`
}
