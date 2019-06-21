package models

import (
	"iserver/db"
)

type User struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"mail"`
	Balance int    `json:"balance"`
}

func (u *User) Find() error {
	row := db.MysqlDB.QueryRow("select name, mail, balance from users where id = ?", u.Id)
	err := row.Scan(&u.Name, &u.Email, &u.Balance)
	if err != nil {
		return err
	}
	return nil
}
