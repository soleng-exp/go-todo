package Handlers

import (
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"../Models"
	"strings"
)

type H map[string]interface{}

func GetTasks() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, Models.GetTasks())
	}
}

func PutTask() echo.HandlerFunc {
	return func(c echo.Context) error {

		var task Models.Task

		c.Bind(&task)

		if len(strings.TrimSpace(task.Name)) == 0 {
			return c.JSON(http.StatusBadRequest, H{
				"error": "name empty",
			})
		}

		id, err := Models.PutTask(task.Name)

		if err != nil {
			return err
		}

		return c.JSON(http.StatusCreated, H{
			"created": id,
		})
	}
}

func DeleteTask() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		nb, err := Models.DeleteTask(id)

		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, H{
			"deleted":  id,
			"affected": nb,
		})
	}
}
