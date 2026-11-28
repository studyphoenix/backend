package model

import (
	"time"

	"github.com/TwiN/logr"
	"gorm.io/gorm"
)

func (c *DBConnector) GetSubtopic(id uint, uid uint) (Subtopic, error) {
	return gorm.G[Subtopic](c.DB).Where("id = ? AND user_id = ?", id, uid).Preload("Children", nil).Preload("Assets", nil).Preload("Chats", nil).First(ctx)
}

func (c *DBConnector) UpdateSubtopic(id uint, req UpdateSubtopicRequest, uid uint) bool {
	subtopic, err := c.GetSubtopic(id, uid)
	if err != nil {
		logr.Errorf("[UPDATE SUBTOPIC] Failed to retrieve subtopic: %w", err)
		return false
	}

	subtopic.Icon = dereferenceOrDefault(req.Icon, subtopic.Icon)
	subtopic.Name = dereferenceOrDefault(req.Name, subtopic.Name)
	subtopic.SubjectID = dereferenceOrDefault(req.Subject, subtopic.SubjectID)
	subtopic.Progress = dereferenceOrDefault(req.Progress, subtopic.Progress)
	subtopic.UpdatedAt = time.Now()

	if err := c.DB.Save(&subtopic).Error; err != nil {
		logr.Errorf("[UPDATE SUBTOPIC] Failed to save changes: %w", err)
		return false
	}
	return true
}

func (c *DBConnector) CreateSubtopic(req CreateSubtopicRequest, userid uint) (uint, bool) {
	subtopic := Subtopic{
		Name:      req.Name,
		Icon:      req.Icon,
		UserID:    userid,
		SubjectID: req.Subject,
		Progress:  req.Progress,
		UpdatedAt: time.Now(),
	}

	result := gorm.WithResult()
	err := gorm.G[Subtopic](c.DB, result).Create(ctx, &subtopic)
	if err != nil {
		logr.Errorf("[CREATE SUBTOPIC] Could not update: %w", err)
		return 0, false
	}
	return subtopic.ID, true
}

func (c *DBConnector) DeleteSubtopic(id uint, uid uint) bool {
	_, err := c.GetSubtopic(id, uid)
	if err != nil {
		logr.Errorf("[DELETE SUBTOPIC] Could not find subtopic to delete: %w", err)
		return false
	}

	err = c.DB.Transaction(func(tx *gorm.DB) error {
		c := DBConnector{
			Config: c.Config,
			DB:     tx,
		}

		return c.deleteSubtopicWithoutTransaction(id, uid)
	})

	if err != nil {
		logr.Errorf("[DELETE SUBTOPIC] Delete transaction failed: %w", err)
		return false
	} else {
		return true
	}
}

// Use this function to delete a subtopic without starting a new transaction.
// Useful to chain it into another transaction
func (c *DBConnector) deleteSubtopicWithoutTransaction(id uint, uid uint) error {
	// Delete assets (if possible)
	assets, err := gorm.G[uint](c.DB).Raw("SELECT asset_id FROM subtopics_assets WHERE subtopic_id = ?", id).Find(ctx)
	if err != nil {
		return err
	}
	for _, asset := range assets {
		if err := c.deleteAssetWithoutTransaction(asset, uid); err != nil {
			return err
		}
	}

	// Delete subtopics (if possible)
	subtopics, err := gorm.G[uint](c.DB).Raw("SELECT child_id FROM subtopics_subtopics WHERE subtopic_id = ?", id).Find(ctx)
	if err != nil {
		return err
	}
	for _, subtopic := range subtopics {
		if err := c.deleteSubtopicWithoutTransaction(subtopic, uid); err != nil {
			return err
		}
	}

	// Delete flashcards (if possible)
	flashcards, err := gorm.G[uint](c.DB).Raw("SELECT flashcard_id FROM subtopics_flashcards WHERE subtopic_id = ?", id).Find(ctx)
	if err != nil {
		return err
	}
	for _, flashcard := range flashcards {
		if err := c.deleteFlashcardWithoutTransaction(flashcard, uid); err != nil {
			return err
		}
	}

	// Delete chats (if possible)
	chats, err := gorm.G[uint](c.DB).Raw("SELECT chat_id FROM subtopics_chats WHERE subtopic_id = ?", id).Find(ctx)
	if err != nil {
		return err
	}
	for _, chat := range chats {
		if err := c.deleteChatWithoutTransaction(chat, uid); err != nil {
			return err
		}
	}

	// Actually delete the subtopic
	if rows, err := gorm.G[Subtopic](c.DB).Where("id = ?", id).Delete(ctx); err != nil || rows == 0 {
		return err
	}

	return nil
}
