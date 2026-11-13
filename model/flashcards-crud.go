package model

import (
	"fmt"
	"time"

	"github.com/TwiN/logr"
	"gorm.io/gorm"
)

func (c *DBConnector) GetFlashcard(id uint) (Flashcard, error) {
	return gorm.G[Flashcard](c.DB).Where("id = ?", id).First(ctx)
}

func (c *DBConnector) UpdateFlashcard(id uint, req UpdateFlashcardRequest) bool {
	flashcard, err := c.GetFlashcard(id)
	if err != nil {
		return false
	}

	flashcard.Content = dereferenceOrDefault(req.Content, flashcard.Content)
	flashcard.Name = dereferenceOrDefault(req.Name, flashcard.Name)
	flashcard.UpdatedAt = time.Now()

	if c.DB.Save(&flashcard).Error != nil {
		return false
	}
	return true
}

func (c *DBConnector) CreateFlashcard(req CreateFlashcardRequest, userid uint) (uint, bool) {
	flashcard := Flashcard{
		Name:      req.Name,
		Content:   req.Content,
		UserID:    userid,
		UpdatedAt: time.Now(),
	}

	result := gorm.WithResult()
	err := gorm.G[Flashcard](c.DB, result).Create(ctx, &flashcard)
	if err != nil {
		return 0, false
	}
	return flashcard.ID, true
}

func (c *DBConnector) DeleteFlashcard(id uint) bool {
	if rows, err := gorm.G[Flashcard](c.DB).Where("id = ?", id).Delete(ctx); err != nil || rows == 0 {
		logr.Debug(fmt.Sprint(err))
		return false
	}
	return true
}

func (c *DBConnector) CompleteFlashcard(id uint) bool {
	flashcard, err := c.GetFlashcard(id)
	if err != nil {
		return false
	}

	flashcard.CompletionCount += 1
	flashcard.UpdatedAt = time.Now()

	if err := c.DB.Save(&flashcard).Error; err != nil {
		logr.Debug(fmt.Sprint(err))
		return false
	}
	return true
}
