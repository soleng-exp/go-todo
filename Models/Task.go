package Models

import (
	"../Database"
	_ "database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Task struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

type TaskCollection struct {
	Tasks []Task `json:"items"`
}

func GetTasks() TaskCollection{
	query := "SELECT * FROM tasks"

	rows, err := Database.Db.Query(query)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	result := TaskCollection{}

	for rows.Next() {
		task := Task{}
		err2 := rows.Scan(&task.Id, &task.Name)

		if err2 != nil {
			panic(err2)
		}
		result.Tasks = append(result.Tasks, task)
	}

	return result
}

func PutTask(name string) (int64, error) {
	query := "INSERT INTO tasks(name) VALUES(?)"

	// Create a prepared SQL statement
	stmt, err := Database.Db.Prepare(query)
	// Exit if we get an error
	if err != nil {
		panic(err)
	}
	// Make sure to cleanup after the program exits
	defer stmt.Close()

	// Replace the '?' in our prepared statement with 'name'
	result, err2 := stmt.Exec(name)
	// Exit if we get an error
	if err2 != nil {
		panic(err2)
	}

	return result.LastInsertId()
}

func DeleteTask (id int) (int64, error) {
	sql := "DELETE FROM tasks WHERE id = ?"

	// Create a prepared SQL statement
	stmt, err := Database.Db.Prepare(sql)
	// Exit if we get an error
	if err != nil {
		panic(err)
	}

	// Replace the '?' in our prepared statement with 'id'
	result, err2 := stmt.Exec(id)
	// Exit if we get an error
	if err2 != nil {
		panic(err2)
	}

	return result.RowsAffected()
}