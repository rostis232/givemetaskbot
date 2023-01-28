package entities

type User struct {
	ChatId   int64  `db:"chat_id" binding:"required"`
	UserName string `db:"user_name"`
}
