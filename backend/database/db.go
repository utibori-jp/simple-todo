package database

import (
	"database/sql"
	"log"
	"simple-todo/backend/models"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	connStr := "user=user password=password host=db sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func GetTodos() ([]models.Todo, error) {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select id, task_name from todos order by id")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var todos []models.Todo
	for rows.Next() {
		var todo models.Todo
		if err := rows.Scan(&todo.No, &todo.Task); err != nil {
			log.Fatal(err)
			return nil, err
		}
		todos = append(todos, todo)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return todos, nil
}

func AddTodo(todo models.Todo) error {
	db, err := ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO todos (task_name) VALUES($1)", todo.Task)
	if err != nil {
		return err
	}
	return nil
}

func UpdateTodo(todo models.Todo) error {
	db, err := ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("UPDATE todos SET task_name = $1 WHERE id = $2", todo.Task, todo.No)
	if err != nil {
		return err
	}
	return nil
}

func DeleteTodo(todoID int8) error {
	db, err := ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM todos WHERE id = $1", todoID)
	if err != nil {
		return err
	}
	return nil
}
