package model

import (
	"errors"

	"gorm.io/gorm"
)

type FlashcardQuery struct {
	Name               *string `form:"q" binding:"omitempty,printascii"`
	ExactMatch         *bool   `form:"exactMatch" binding:"omitempty,boolean"`
	Offset             *uint   `form:"offset" binding:"omitempty,min=5,max=50"` // OFFSET
	Limit              *uint   `form:"count" binding:"omitempty,min=5,max=50"`  // LIMIT
	SubtopicID         *uint   `form:"subtopicId" binding:"omitempty,min=0"`
	MinCompletionCount *uint   `form:"minCompletionCount" binding:"omitempty,numeric,min=0"`
	MaxCompletionCount *uint   `form:"maxCompletionCount" binding:"omitempty,numeric,min=0"`
}

func (c *DBConnector) QueryFlashcards(query FlashcardQuery, uid uint) ([]Flashcard, error) {
	sql := gorm.G[Flashcard](c.DB).Where("user_id = ?", uid)

	// Name
	if query.Name != nil {
		if query.ExactMatch != nil && !*query.ExactMatch {
			sql = sql.Where("name = ?", *query.Name)
		} else {
			sql = sql.Where("name LIKE ?", "%"+*query.Name+"%")
		}
	}

	// CompletionCount
	if query.MaxCompletionCount != nil {
		sql = sql.Where("completion_count < ?", *query.MaxCompletionCount+1)
	}

	if query.MinCompletionCount != nil {
		sql = sql.Where("completion_count > ?", *query.MinCompletionCount-1)
	}

	// SubtopicID
	if query.SubtopicID != nil {
		sq := c.DB.Select("flashcard_id").Where("subtopic_id = ?", *query.SubtopicID).Table("subtopics_flashcards")
		sql = sql.Where("id IN (?)", sq)
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
