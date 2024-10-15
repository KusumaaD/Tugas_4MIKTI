package models

import (
	"database/sql"
)

type User struct {
  ID       int    `json:"id"`
  Username string `json:"username"`
  Password string `json:"-"`
  Role     string `json:"role"`
}

func GetAllUsers(db *sql.DB) ([]User, error) {
  rows, err := db.Query("SELECT id, username, role FROM users")
  if err != nil {
    return nil, err
  }
  defer rows.Close()

  var users []User
  for rows.Next() {
    var u User
    if err := rows.Scan(&u.ID, &u.Username, &u.Role); err != nil {
      return nil, err
    }
    users = append(users, u)
  }
  return users, nil
}

func CreateUser(db *sql.DB, user *User) (int, error) {
  result, err := db.Exec("INSERT INTO users (username, password, role) VALUES (?, ?, ?)",
    user.Username, user.Password, user.Role)
  if err != nil {
    return 0, err
  }
  id, err := result.LastInsertId()
  return int(id), err
}

func UpdateUser(db *sql.DB, user *User) error {
  _, err := db.Exec("UPDATE users SET username = ?, role = ? WHERE id = ?",
    user.Username, user.Role, user.ID)
  return err
}

func DeleteUser(db *sql.DB, id int) error {
  _, err := db.Exec("DELETE FROM users WHERE id = ?", id)
  return err
}
