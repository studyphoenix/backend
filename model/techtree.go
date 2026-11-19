package model

import "gorm.io/gorm"

type SubtopicCreate struct {
	TempID uint   `json:"tempId"`
	Name   string `json:"name" validate:"required" binding:"required,max=22,min=3"`
}

type Edge struct {
	From uint
	To   uint
}

type CreateSubtopicBulkRequest struct {
	SubjectID uint             `json:"subjectId" validate:"required" binding:"required,numeric,min=0"`
	Subtopics []SubtopicCreate `json:"subtopics" validate:"required"`
	Edges     []Edge           `json:"edges" validate:"required"`
}

func (c *DBConnector) GetTechTreeForSubject(subjectid uint, uid uint) ([]Subtopic, error) {
	return gorm.G[Subtopic](c.DB).Where("subject_id = ? AND user_id = ?", subjectid, uid).Preload("Children", nil).Find(ctx)
}
