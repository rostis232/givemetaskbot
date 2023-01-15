package entities

type Group struct {
	Id          int64  `db:"group_id"`
	ChiefUserId int64  `db:"chief_user_id"`
	GroupName   string `db:"group_name"`
}
