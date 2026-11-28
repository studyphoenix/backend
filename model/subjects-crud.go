package model

import (
	"time"

	"github.com/TwiN/logr"
	"gorm.io/gorm"
)

func (c *DBConnector) GetSubject(id uint, uid uint) (Subject, error) {
	return gorm.G[Subject](c.DB).Where("id = ? AND user_id = ?", id, uid).Preload("Assets", nil).First(ctx)
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

	err = c.DB.Transaction(func(tx *gorm.DB) error {
		c := DBConnector{
			Config: c.Config,
			DB:     tx,
		}

		return c.deleteSubjectWithoutTransaction(id, uid)
	})

	if err != nil {
		logr.Errorf("[DELETE SUBTOPIC] Delete transaction failed: %w", err)
		return false
	} else {
		return true
	}
}

func (c *DBConnector) deleteSubjectWithoutTransaction(id uint, uid uint) error {
	// Delete flashcards (if possible)
	assets, err := gorm.G[uint](c.DB).Raw("SELECT asset_id FROM subjects_assets WHERE subject_id = ?", id).Find(ctx)
	if err != nil {
		return err
	}
	for _, asset := range assets {
		if err := c.deleteAssetWithoutTransaction(asset, uid); err != nil {
			return err
		}
	}

	subtopics, err := gorm.G[uint](c.DB).Raw("SELECT id FROM subtopics WHERE subject_id = ?", id).Find(ctx)
	if err != nil {
		return err
	}
	for _, subtopic := range subtopics {
		if err := c.deleteSubtopicWithoutTransaction(subtopic, uid); err != nil {
			return err
		}
	}

	// Actually delete the subject
	if rows, err := gorm.G[Subject](c.DB).Where("id = ?", id).Delete(ctx); err != nil || rows == 0 {
		return err
	}

	return nil
}
