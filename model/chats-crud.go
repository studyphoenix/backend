package model

import (
	"time"

	"gorm.io/gorm"
)

func (c *DBConnector) GetChat(id uint) (Chat, error) {
	return gorm.G[Chat](c.DB).Where("id = ?", id).Preload("ChatMessages", nil).First(ctx)
}

func (c *DBConnector) UpdateChat(id uint, req UpdateChatRequest) bool {
	chat, err := c.GetChat(id)
	if err != nil {
		return false
	}

	chat.Mode = dereferenceOrDefault(req.Mode, chat.Mode)
	chat.Name = dereferenceOrDefault(req.Name, chat.Name)
	chat.UpdatedAt = time.Now()

	if c.DB.Save(&chat).Error != nil {
		return false
	}
	return true
}

func (c *DBConnector) CreateChat(req CreateChatRequest, userid uint) (uint, bool) {
	chat := Chat{
		Name:      req.Name,
		Mode:      req.Mode,
		UserID:    userid,
		UpdatedAt: time.Now(),
	}

	result := gorm.WithResult()
	err := gorm.G[Chat](c.DB, result).Create(ctx, &chat)
	if err != nil {
		return 0, false
	}
	return chat.ID, true
}

func (c *DBConnector) DeleteChat(id uint) bool {
	if rows, err := gorm.G[Chat](c.DB).Where("id = ?", id).Delete(ctx); err != nil || rows == 0 {
		return false
	}
	return true
}
