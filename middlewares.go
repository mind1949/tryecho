package main

import (
	"log"
	"github.com/labstack/echo"
)

func validateSession(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println("validateSession !")
		return next(c)
	}
}
