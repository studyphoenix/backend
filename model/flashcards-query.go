package model

import "gorm.io/gorm"

type FlashcardQuery struct {
	Name               *string `form:"name" binding:"omitempty,printascii"`
	Offset             *uint   `form:"offset" binding:"omitempty,min=5,max=50"` // OFFSET
	Limit              *uint   `form:"count" binding:"omitempty,min=5,max=50"`  // LIMIT
	SubtopicID         *uint   `form:"subtopicId" binding:"omitempty,min=0"`
	MinCompletionCount *uint   `form:"minCompletionCount" binding:"omitempty,numeric,min=0"`
	MaxCompletionCount *uint   `form:"maxCompletionCount" binding:"omitempty,numeric,min=0"`
}

func (c *DBConnector) QueryFlashcards(query FlashcardQuery, uid uint) ([]Flashcard, error) {
	filter := ""
	return gorm.G[Flashcard](c.DB).Where(filter).Find(ctx)
}
