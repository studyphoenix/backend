package model

import (
	"time"

	"gorm.io/gorm"
)

func (c *DBConnector) GetSubtopic(id uint) (Subtopic, error) {
	return gorm.G[Subtopic](c.DB).Where("id = ?", id).Preload("Subtopics", nil).Preload("Assets", nil).Preload("Chats", nil).First(ctx)
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
	subtopic.UpdatedAt = time.Now()

	if c.DB.Save(&subtopic).Error != nil {
		return false
	}
	return true
}

func (c *DBConnector) CreateSubtopic(req CreateSubtopicRequest, userid uint) bool {
	subtopic := Subtopic{
		Name:      req.Name,
		Icon:      req.Icon,
		UserID:    userid,
		SubjectID: req.Subject,
		UpdatedAt: time.Now(),
	}

	if c.DB.Save(&subtopic).Error != nil {
		return false
	}
	return true
}

func (c *DBConnector) DeleteSubtopic(id uint) bool {
	if rows, err := gorm.G[Subtopic](c.DB).Where("id = ?", id).Delete(ctx); err != nil || rows == 0 {
		return false
	}
	// TODO: Delete anything that can't really be done via DB consistency
	return true
}
