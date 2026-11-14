package model

import (
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type FlashcardQuery struct {
	Name               *string `form:"name" binding:"omitempty,printascii"`
	Offset             *uint   `form:"offset" binding:"omitempty,min=5,max=50"` // OFFSET
	Limit              *uint   `form:"count" binding:"omitempty,min=5,max=50"`  // LIMIT
	SubtopicID         *uint   `form:"subtopicId" binding:"omitempty,min=0"`
	MinCompletionCount *uint   `form:"minCompletionCount" binding:"omitempty,numeric,min=0"`
	MaxCompletionCount *uint   `form:"maxCompletionCount" binding:"omitempty,numeric,min=0"`
}

func (c *DBConnector) QueryFlashcards(query FlashcardQuery, uid uint) ([]Flashcard, error) {
	sql := gorm.G[Flashcard](c.DB).Where("uid = ?", uid)

	// Name
	if query.Name != nil {
		sql = sql.Where("name = ?", *query.Name)
	}

	// SubtopicID
	if query.Name != nil {
		// TEST: This is probably incorrect
		sql = sql.Joins(clause.InnerJoin.Association("subtopics_flashcards"), nil).Where("subtopic_id = ?", *query.SubtopicID)
	}

	// CompletionCount
	if query.MaxCompletionCount != nil {
		sql = sql.Where("completion_count < ?", *query.MaxCompletionCount+1)
	}

	if query.MinCompletionCount != nil {
		sql = sql.Where("completion_count > ?", *query.MinCompletionCount+1)
	}

	// Offset
	if query.Offset != nil {
		offset := int(*query.Offset)
		if offset < 0 {
			return nil, errors.New("invalid offset specified")
		}
		sql = sql.Offset(offset)
	}

	// Limit
	limit := 25
	if query.Limit != nil {
		limit = int(*query.Limit)
	}

	if limit < 5 || limit > 50 {
		limit = 25
	}

	return sql.Limit(limit).Find(ctx)
}
