package model

import (
	"time"

	"gorm.io/gorm"
)

func (c *DBConnector) GetSubject(id uint) (Subject, error) {
	return gorm.G[Subject](c.DB).Where("id = ?", id).First(ctx)
}

func (c *DBConnector) UpdateSubject(id uint, req UpdateSubjectRequest) bool {
	subject, err := c.GetSubject(id)
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

func (c *DBConnector) CreateSubject(req CreateSubjectRequest, userid uint) bool {
	subject := Subject{
		Name:      req.Name,
		Icon:      req.Icon,
		UserID:    userid,
		UpdatedAt: time.Now(),
	}

	if c.DB.Save(&subject).Error != nil {
		return false
	}
	return true
}

func (c *DBConnector) DeleteSubject(id uint) bool {
	if rows, err := gorm.G[Subject](c.DB).Where("id = ?", id).Delete(ctx); err != nil || rows == 0 {
		return false
	}
	// TODO: Delete anything that can't really be done via DB consistency
	return true
}
