package controllers

import (
	"github.com/labstack/echo"
	"iserver/models"
	"net/http"
	"strconv"
)

func FindUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	u := &models.User{Id: id}
	err = u.Find()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, u)
}
