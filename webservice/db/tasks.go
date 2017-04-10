package db

/*
Stores the database functions related to tasks like
GetTaskByID(id int)
GetTasks(status string)
DeleteAll()
*/

import (


	_ "../model"
)

var taskStatus map[string]int
var err error

//Database encapsulates database
type Database struct {
}

func init() {

}

//Close function closes this database connection
func Close() {
}




//TrashTask is used to delete the task
func TrashTask(username string, id int) error {
	return nil
}

//CompleteTask  is used to mark tasks as complete
func CompleteTask(username string, id int) error {
	return nil
}

//DeleteAll is used to empty the trash
func DeleteAll(username string) error {
	return nil
}

//RestoreTask is used to restore tasks from the Trash
func RestoreTask(username string, id int) error {
	return nil
}

//RestoreTaskFromComplete is used to restore tasks from the Trash
func RestoreTaskFromComplete(username string, id int) error {
	return nil
}

//DeleteTask is used to delete the task from the database
func DeleteTask(username string, id int) error {
	return nil
}

//AddTask is used to add the task in the database
//TODO: add dueDate feature later
func AddTask(title, content, category string, taskPriority int, username string, hidden int) error {

	return nil
}


//UpdateTask is used to update the tasks in the database
func UpdateTask(id int, title, content, category string, priority int, username string, hidden int) error {
	return nil
}

//taskQuery encapsulates running multiple queries which don't do much things
func taskQuery(sql string, args ...interface{}) error {

	return nil
}



//AddComments will be used to add comments in the database
func AddComments(username string, id int, comment string) error {

	return nil
}
