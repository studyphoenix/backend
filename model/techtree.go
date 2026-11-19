package model

import (
	"errors"
	"strconv"

	"github.com/TwiN/logr"
	"gorm.io/gorm"
)

type SubtopicCreate struct {
	TempID uint    `json:"tempId" validate:"required" binding:"required,numeric"`
	Name   string  `json:"name" validate:"required" binding:"required,max=22,min=3"`
	Icon   *string `json:"icon" validate:"optional" binding:"omitempty,max=100,min=0,printascii"`
}

type Edge struct {
	From uint `json:"from" validate:"required" binding:"required,numeric"`
	To   uint `json:"to" validate:"required" binding:"required,numeric"`
}

type CreateSubtopicBulkRequest struct {
	SubjectID uint             `json:"subjectId" validate:"required" binding:"required,numeric,min=0"`
	Subtopics []SubtopicCreate `json:"subtopics" validate:"required" binding:"required,dive"`
	Edges     []Edge           `json:"edges" validate:"optional" binding:"omitempty,dive"`
}

func (c *DBConnector) GetTechTreeForSubject(subjectid uint, uid uint) ([]Subtopic, error) {
	return gorm.G[Subtopic](c.DB).Where("subject_id = ? AND user_id = ?", subjectid, uid).Preload("Children", nil).Find(ctx)
}

func (c *DBConnector) BulkCreateTechTreeForSubject(req CreateSubtopicBulkRequest, uid uint) bool {
	_, err := c.GetSubject(req.SubjectID, uid)

	if err != nil {
		return false
	}

	err = c.DB.Transaction(func(tx *gorm.DB) error {
		c := DBConnector{
			Config: c.Config,
			DB:     tx,
		}
		tempIdToInternalMapping := make(map[uint]uint)

		for _, v := range req.Subtopics {
			subtopicId, ok := c.CreateSubtopic(CreateSubtopicRequest{
				Progress: 0,
				Name:     v.Name,
				Icon:     dereferenceOrDefault(v.Icon, ""),
				Subject:  req.SubjectID,
			}, uid)

			if !ok {
				return errors.New("Failed to create subtopics with name: " + v.Name)
			}

			tempIdToInternalMapping[v.TempID] = subtopicId
		}

		for _, v := range req.Edges {
			from, ok := tempIdToInternalMapping[v.From]
			if !ok {
				return errors.New("FROM ID not found in map")
			}
			to, ok := tempIdToInternalMapping[v.To]
			if !ok {
				return errors.New("TO ID not found in map")
			}

			if !c.AssociateSubtopicToSubtopic(from, to, uid) {
				return errors.New("Failed to associate subtopic " +
					strconv.FormatUint(uint64(from), 10) + " to subtopic " +
					strconv.FormatUint(uint64(to), 10))
			}
		}
		return nil
	})

	if err != nil {
		logr.Errorf("%e", err)
		return false
	} else {
		return true
	}
}
