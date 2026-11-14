package controller

import "github.com/studyphoenix/backend/model"

type Controller struct {
	db *model.DBConnector
}

func NewController(dbcon *model.DBConnector) *Controller {
	return &Controller{
		db: dbcon,
	}
}
