package model

import "gorm.io/gorm"

type SubtopicCreate struct {
	TempID uint   `json:"tempId" validate:"required" binding:"required,numeric"`
	Name   string `json:"name" validate:"required" binding:"required,max=22,min=3"`
}

type Edge struct {
	From uint `json:"from" validate:"required" binding:"required,numeric"`
	To   uint `json:"to" validate:"required" binding:"required,numeric"`
}

type CreateSubtopicBulkRequest struct {
	SubjectID uint             `json:"subjectId" validate:"required" binding:"required,numeric,min=0"`
	Subtopics []SubtopicCreate `json:"subtopics" validate:"required" binding:"required"`
	Edges     []Edge           `json:"edges" validate:"required" binding:"required"`
}

func (c *DBConnector) GetTechTreeForSubject(subjectid uint, uid uint) ([]Subtopic, error) {
	return gorm.G[Subtopic](c.DB).Where("subject_id = ? AND user_id = ?", subjectid, uid).Preload("Children", nil).Find(ctx)
}

func (c *DBConnector) BulkCreateTechTreeForSubject(req CreateSubtopicBulkRequest)  {
	
}
