package model

import (
	"time"

	"gorm.io/gorm"
)

func (c *DBConnector) GetSubject(id uint, uid uint) (Subject, error) {
	return gorm.G[Subject](c.DB).Where("id = ? AND user_id = ?", id, uid).First(ctx)
}

func (c *DBConnector) UpdateSubject(id uint, req UpdateSubjectRequest, uid uint) bool {
	subject, err := c.GetSubject(id, uid)
	if err != nil {
		return false
	}

	subject.Icon = dereferenceOrDefault(req.Icon, subject.Icon)
	subject.Name = dereferenceOrDefault(req.Name, subject.Name)
	subject.UpdatedAt = time.Now()

	if c.DB.Save(&subject).Error != nil {
		return false
	}
	return true
}

func (c *DBConnector) CreateSubject(req CreateSubjectRequest, userid uint) (uint, bool) {
	subject := Subject{
		Name:      req.Name,
		Icon:      req.Icon,
		UserID:    userid,
		UpdatedAt: time.Now(),
	}

	result := gorm.WithResult()
	err := gorm.G[Subject](c.DB, result).Create(ctx, &subject)
	if err != nil {
		return 0, false
	}
	return subject.ID, true
}

func (c *DBConnector) DeleteSubject(id uint, uid uint) bool {
	_, err := c.GetSubject(id, uid)
	if err != nil {
		return false
	}

	// Delete flashcards (if possible)
	assets, err := gorm.G[uint](c.DB).Raw("SELECT asset_id FROM subjects_assets WHERE subject_id = ?", id).Find(ctx)
	if err != nil {
		return false
	}
	for _, asset := range assets {
		c.DeleteAsset(asset, uid)
	}

	subtopics, err := gorm.G[uint](c.DB).Raw("SELECT id FROM subtopics WHERE subject_id = ?", id).Find(ctx)
	if err != nil {
		return false
	}
	for _, subtopic := range subtopics {
		c.DeleteSubtopic(subtopic, uid)
	}

	// Actually delete the subject
	if rows, err := gorm.G[Subject](c.DB).Where("id = ?", id).Delete(ctx); err != nil || rows == 0 {
		return false
	}

	return true
}
