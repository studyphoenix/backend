package model

import "gorm.io/gorm"

func (c *DBConnector) GetTechTreeForSubject(subjectid uint, uid uint) ([]Subtopic, error) {
	return gorm.G[Subtopic](c.DB).Where("subject_id = ? AND user_id = ?", subjectid, uid).Preload("Children", nil).Find(ctx)
}
