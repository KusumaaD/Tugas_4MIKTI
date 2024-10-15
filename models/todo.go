package models

import (
	"database/sql"
)

type Todo struct {
  ID          int    `json:"id"`
  Title       string `json:"title"`
  Description string `json:"description"`
  Completed   bool   `json:"completed"`
}

func GetAllTodos(db *sql.DB) ([]Todo, error) {
  rows, err := db.Query("SELECT id, title, description, completed FROM todos")
  if err != nil {
    return nil, err
  }
  defer rows.Close()

  var todos []Todo
  for rows.Next() {
    var t Todo
    if err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Completed); err != nil {
      return nil, err
    }
    todos = append(todos, t)
  }
  return todos, nil
}

func CreateTodo(db *sql.DB, todo *Todo) (int, error) {
  result, err := db.Exec("INSERT INTO todos (title, description, completed) VALUES (?, ?, ?)",
    todo.Title, todo.Description, todo.Completed)
  if err != nil {
    return 0, err
  }
  id, err := result.LastInsertId()
  return int(id), err
}

func UpdateTodo(db *sql.DB, todo *Todo) error {
  _, err := db.Exec("UPDATE todos SET title = ?, description = ?, completed = ? WHERE id = ?",
    todo.Title, todo.Description, todo.Completed, todo.ID)
  return err
}

func DeleteTodo(db *sql.DB, id int) error {
  _, err := db.Exec("DELETE FROM todos WHERE id = ?", id)
  return err
}
