package controllers

import (
	"iserver/models"
)

func FindUser(req *struct{ ID *int}) (*models.User, error) {
	u := &models.User{Id: *req.ID}
	err := u.Find()
	if err != nil {
		return nil, err
	}
	return u, nil
}
