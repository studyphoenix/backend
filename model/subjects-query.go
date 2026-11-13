package model

import "gorm.io/gorm"

func (c *DBConnector) SubjectQuery(id uint) ([]Subject, error) {
	return gorm.G[Subject](c.DB).Where("id = ?", id).Find(ctx)
}
