package main

import (
	"github.com/labstack/echo"
	"iserver/db"
	"log"
	"net/http"
)

func validateSession(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sidCookie, err := c.Cookie(db.SidKey)
		if err != nil {
			return c.JSON(http.StatusForbidden, "no session id in cookie !")
		}
		session, err := db.RedisClient.Do("Get", sidCookie.Value).String()
		if err != nil {
			log.Printf("err: redisClient.Do('Get', %v) session=%q err=%v", sidCookie.Value, session, err)
			return c.JSON(http.StatusForbidden, "session verification failed !")
		}
		return next(c)
	}
}
