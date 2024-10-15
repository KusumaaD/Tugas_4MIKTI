package handlers

import (
	"Tugas_4MIKTII/models"
	"database/sql"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func Login(db *sql.DB) echo.HandlerFunc {
  return func(c echo.Context) error {
    username := c.FormValue("username")
    password := c.FormValue("password")

    // Verify user credentials
    user, err := getUserByUsername(db, username)
    if err != nil {
      return echo.ErrUnauthorized
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
      return echo.ErrUnauthorized
    }

    // Create token
    token := jwt.New(jwt.SigningMethodHS256)

    // Set claims
    claims := token.Claims.(jwt.MapClaims)
    claims["username"] = user.Username
    claims["role"] = user.Role
    claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

    // Generate encoded token
    t, err := token.SignedString([]byte("secret"))
    if err != nil {
      return err
    }

    return c.JSON(http.StatusOK, map[string]string{
      "token": t,
    })
  }
}

func getUserByUsername(db *sql.DB, username string) (*models.User, error) {
    user := &models.User{}
    err := db.QueryRow("SELECT id, username, password, role FROM users WHERE username = ?", username).Scan(&user.ID, &user.Username, &user.Password, &user.Role)
    if err != nil {
        return nil, err
    }
    return user, nil
}
