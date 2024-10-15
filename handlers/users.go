package handlers

import (
	"Tugas_4MIKTII/models"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func GetAllUsers(db *sql.DB) echo.HandlerFunc {
  return func(c echo.Context) error {
    users, err := models.GetAllUsers(db)
    if err != nil {
      return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
    }
    return c.JSON(http.StatusOK, users)
  }
}

func CreateUser(db *sql.DB) echo.HandlerFunc {
  return func(c echo.Context) error {
    user := new(models.User)
    if err := c.Bind(user); err != nil {
      return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }
    
    // Hash password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
      return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
    }
    user.Password = string(hashedPassword)
    
    id, err := models.CreateUser(db, user)
    if err != nil {
      return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
    }
    
    user.ID = id
    return c.JSON(http.StatusCreated, user)
  }
}

func UpdateUser(db *sql.DB) echo.HandlerFunc {
  return func(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    user := new(models.User)
    if err := c.Bind(user); err != nil {
      return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }
    
    user.ID = id
    if user.Password != "" {
      // Hash new password if provided
      hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
      if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
      }
      user.Password = string(hashedPassword)
    }
    
    if err := models.UpdateUser(db, user); err != nil {
      return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
    }
    
    return c.JSON(http.StatusOK, user)
  }
}

func DeleteUser(db *sql.DB) echo.HandlerFunc {
  return func(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    if err := models.DeleteUser(db, id); err != nil {
      return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
    }
    return c.NoContent(http.StatusNoContent)
  }
}
