package handlers

import (
	"Tugas_4MIKTII/models"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAllTodos(db *sql.DB) echo.HandlerFunc {
  return func(c echo.Context) error {
    todos, err := models.GetAllTodos(db)
    if err != nil {
      return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
    }
    return c.JSON(http.StatusOK, todos)
  }
}

func CreateTodo(db *sql.DB) echo.HandlerFunc {
  return func(c echo.Context) error {
    todo := new(models.Todo)
    if err := c.Bind(todo); err != nil {
      return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }
    
    id, err := models.CreateTodo(db, todo)
    if err != nil {
      return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
    }
    
    todo.ID = id
    return c.JSON(http.StatusCreated, todo)
  }
}

func UpdateTodo(db *sql.DB) echo.HandlerFunc {
  return func(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    todo := new(models.Todo)
    if err := c.Bind(todo); err != nil {
      return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }
    
    todo.ID = id
    if err := models.UpdateTodo(db, todo); err != nil {
      return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
    }
    
    return c.JSON(http.StatusOK, todo)
  }
}

func DeleteTodo(db *sql.DB) echo.HandlerFunc {
  return func(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    if err := models.DeleteTodo(db, id); err != nil {
      return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
    }
    return c.NoContent(http.StatusNoContent)
  }
}
