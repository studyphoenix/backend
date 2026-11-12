package model

import (
	"time"

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

func (c *DBConnector) CreateFlashcard(req CreateFlashcardRequest, userid uint) bool {
	flashcard := Flashcard{
		Name:      req.Name,
		Content:   req.Content,
		UserID:    userid,
		UpdatedAt: time.Now(),
	}

	if c.DB.Save(&flashcard).Error != nil {
		return false
	}
	return true
}

func (c *DBConnector) DeleteFlashcard(id uint) bool {
	if rows, err := gorm.G[Flashcard](c.DB).Where("id = ?", id).Delete(ctx); err != nil || rows == 0 {
		return false
	}
	return true
}
