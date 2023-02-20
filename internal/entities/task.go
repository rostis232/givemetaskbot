package entities

type Task struct {
	Id              int64  `db:"task_id"`
	TaskName        string `db:"task_name"`
	TaskDescription string `db:"task_description"`
	GroupId         int64  `db:"group_id"`
	CreatingTime    string `db:"creating_time"`
}
