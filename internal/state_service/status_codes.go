package state_service

const (
	Undefined_status = iota
	Expecting_language
	Expecting_new_user_name
	Changing_user_name
	Expecting_new_group_name
	Changing_group_name
	Expecting_new_task_title
	Changging_task_title
	Expecting_new_task_description
	Changing_task_description
	Adding_employee_to_group
	Adding_employee_to_task
	MainMenu
	Expecting_time_zone //It`ll be need to be used when implementing time control
)
