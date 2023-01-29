package entities

import (
	"github.com/rostis232/givemetaskbot/internal/messages"
)

type User struct {
	UserId      int64             `db:"user_id"`
	ChatId      int64             `db:"chat_id" binding:"required"`
	UserName    string            `db:"user_name,omitempty"`
	Language    messages.Language `db:"language,omitempty"`
	Status      int               `db:"status,omitempty"`
	ActiveGroup int64             `db:"active_group,omitempty"`
	ActiveTask  int64             `db:"active_task,omitempty"`
}
