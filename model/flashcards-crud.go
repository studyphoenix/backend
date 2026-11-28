package model

import (
	"time"

	"github.com/TwiN/logr"
	"gorm.io/gorm"
)

func (c *DBConnector) GetFlashcard(id uint, uid uint) (Flashcard, error) {
	return gorm.G[Flashcard](c.DB).Where("id = ? AND user_id = ?", id, uid).First(ctx)
}

func (c *DBConnector) UpdateFlashcard(id uint, req UpdateFlashcardRequest, uid uint) bool {
	flashcard, err := c.GetFlashcard(id, uid)
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

func (c *DBConnector) DeleteFlashcard(id uint, uid uint) bool {
	_, err := c.GetFlashcard(id, uid)
	if err != nil {
		return false
	}

	err = c.DB.Transaction(func(tx *gorm.DB) error {
		c := DBConnector{
			Config: c.Config,
			DB:     tx,
		}

		return c.deleteFlashcardWithoutTransaction(id, uid)
	})

	if err != nil {
		logr.Errorf("[DELETE FLASHCARD] Delete transaction failed: %e", err)
		return false
	} else {
		return true
	}
}

func (c *DBConnector) deleteFlashcardWithoutTransaction(id uint, uid uint) error {
	c.DB.Exec("DELETE FROM `subtopics_flashcards` WHERE flashcard_id = ?", id)

	if rows, err := gorm.G[Flashcard](c.DB).Where("id = ?", id).Delete(ctx); err != nil || rows == 0 {
		return err
	}

	return nil
}

func (c *DBConnector) CompleteFlashcard(id uint, uid uint) bool {
	flashcard, err := c.GetFlashcard(id, uid)
	if err != nil {
		return false
	}

	flashcard.CompletionCount += 1
	flashcard.UpdatedAt = time.Now()

	if err := c.DB.Save(&flashcard).Error; err != nil {
		logr.Debugf("%e", err)
		return false
	}
	return true
}
