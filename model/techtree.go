package model

import "gorm.io/gorm"

func (c *DBConnector) GetTechTreeForSubject(subjectid uint) ([]Subtopic, error) {
	return gorm.G[Subtopic](c.DB).Where("subject_id = ?", subjectid).Find(ctx)
}
