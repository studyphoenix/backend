package model

import (
	"strconv"

	"github.com/TwiN/logr"
	"gorm.io/gorm"
)

type SubtopicCreate struct {
	TempID uint   `json:"tempId" validate:"required" binding:"required,numeric"`
	Name   string `json:"name" validate:"required" binding:"required,max=22,min=3"`
	Icon   string `json:"icon" validate:"required" binding:"required,max=100,min=0,printascii"`
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

func (c *DBConnector) BulkCreateTechTreeForSubject(req CreateSubtopicBulkRequest, uid uint) bool {
	tempIdToInternalMapping := make(map[uint]uint)

	for _, v := range req.Subtopics {
		subtopicId, ok := c.CreateSubtopic(CreateSubtopicRequest{
			Progress: 0,
			Name:     v.Name,
			Icon:     v.Icon,
			Subject:  req.SubjectID,
		}, uid)

		if !ok {
			logr.Error("Failed to create subtopic with name: " + v.Name)
			return false
		}

		tempIdToInternalMapping[v.TempID] = subtopicId
	}

	for _, v := range req.Edges {
		if !c.AssociateSubtopicToSubtopic(tempIdToInternalMapping[v.From], tempIdToInternalMapping[v.To], uid) {
			logr.Error("Failed to associate subtopic " +
				strconv.FormatUint(uint64(tempIdToInternalMapping[v.From]), 10) + " to subtopic " +
				strconv.FormatUint(uint64(tempIdToInternalMapping[v.To]), 10))
		}
	}

	return true
}
