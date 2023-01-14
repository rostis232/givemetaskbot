package entities

var UserStates []UserState

type UserState struct {
	ChatId        int64
	CreatingGroup bool
}
