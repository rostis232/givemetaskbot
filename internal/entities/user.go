package entities

import "github.com/rostis232/givemetaskbot/internal/messages"

type User struct {
	ChatId      int64             `db:"chat_id" binding:"required"`
	UserName    string            `db:"user_name" binding:"omitempty"`
	Language    messages.Language `db:"language,omitempty"`
	Status      int64             `db:"status,omitempty"`
	ActiveGroup int64             `db:"active_group,omitempty"`
	ActiveTask  int64             `db:"active_task,omitempty"`
}
