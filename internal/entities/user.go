package entities

type User struct {
	Id       int64  `db:"user_id"`
	ChatId   int64  `db:"chat_id" binding:"required"`
	UserName string `db:"user_name" binding:"required"`
}
