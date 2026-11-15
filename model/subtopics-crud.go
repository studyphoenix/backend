package model

import (
	"time"

	"gorm.io/gorm"
)

func (c *DBConnector) GetSubtopic(id uint) (Subtopic, error) {
	return gorm.G[Subtopic](c.DB).Where("id = ?", id).Preload("Children", nil).Preload("Assets", nil).Preload("Chats", nil).First(ctx)
}

func (c *DBConnector) UpdateSubtopic(id uint, req UpdateSubtopicRequest) bool {
	subtopic, err := c.GetSubtopic(id)
	if err != nil {
		return false
	}

	subtopic.Icon = dereferenceOrDefault(req.Icon, subtopic.Icon)
	subtopic.Name = dereferenceOrDefault(req.Name, subtopic.Name)
	// FIXME: 9.3 severity vulnerability CVSS:4.0/AV:N/AC:L/AT:N/PR:N/UI:N/VC:H/VI:H/VA:N/SC:N/SI:N/SA:N
	// caused by not checking ownership of subject (also in all other retrieve, update, delete and associate endpoints)
	subtopic.SubjectID = dereferenceOrDefault(req.Subject, subtopic.SubjectID)
	subtopic.Progress = dereferenceOrDefault(req.Progress, subtopic.Progress)
	subtopic.UpdatedAt = time.Now()

	if c.DB.Save(&subtopic).Error != nil {
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
		return 0, false
	}
	return subtopic.ID, true
}

func (c *DBConnector) DeleteSubtopic(id uint) bool {
	// Delete assets (if possible)
	assets, err := gorm.G[uint](c.DB).Raw("SELECT asset_id FROM subtopics_assets WHERE subtopic_id = ?", id).Find(ctx)
	if err != nil {
		return false
	}
	for _, asset := range assets {
		c.DeleteAsset(asset)
	}

	// Delete subtopics (if possible)
	subtopics, err := gorm.G[uint](c.DB).Raw("SELECT child_id FROM subtopics_subtopics WHERE subtopic_id = ?", id).Find(ctx)
	if err != nil {
		return false
	}
	for _, subtopic := range subtopics {
		c.DeleteAsset(subtopic)
	}

	// Delete flashcards (if possible)
	flashcards, err := gorm.G[uint](c.DB).Raw("SELECT flashcard_id FROM subtopics_flashcards WHERE subtopic_id = ?", id).Find(ctx)
	if err != nil {
		return false
	}
	for _, flashcard := range flashcards {
		c.DeleteFlashcard(flashcard)
	}

	// Delete chats (if possible)
	chats, err := gorm.G[uint](c.DB).Raw("SELECT chat_id FROM subtopics_chats WHERE subtopic_id = ?", id).Find(ctx)
	if err != nil {
		return false
	}
	for _, chat := range chats {
		c.DeleteChat(chat)
	}

	// Actually delete the subtopic
	if rows, err := gorm.G[Subtopic](c.DB).Where("id = ?", id).Delete(ctx); err != nil || rows == 0 {
		return false
	}

	return true
}
