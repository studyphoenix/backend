package model

import (
	"time"

	"gorm.io/gorm"
)

func (c *DBConnector) GetChatMessage(id uint) (ChatMessage, error) {
	return gorm.G[ChatMessage](c.DB).Where("id = ?", id).First(ctx)
}

func (c *DBConnector) UpdateChatMessage(id uint, req UpdateChatMessageRequest) bool {
	msg, err := c.GetChatMessage(id)
	if err != nil {
		return false
	}

	msg.ModelMessage = dereferenceOrDefault(req.ModelMessage, msg.ModelMessage)
	msg.UserMessage = dereferenceOrDefault(req.UserMessage, msg.UserMessage)
	msg.Summary = dereferenceOrDefault(req.Summary, msg.Summary)
	msg.UpdatedAt = time.Now()

	if c.DB.Save(&msg).Error != nil {
		return false
	}
	return true
}

func (c *DBConnector) CreateChatMessage(req AddChatMessageRequest, chatid uint) (uint, bool) {
	msg := ChatMessage{
		UpdatedAt:    time.Now(),
		ChatID:       chatid,
		UserMessage:  req.UserMessage,
		ModelMessage: req.ModelMessage,
		Summary:      req.Summary,
	}

	result := gorm.WithResult()
	err := gorm.G[ChatMessage](c.DB, result).Create(ctx, &msg)
	if err != nil {
		return 0, false
	}
	return msg.ID, true
}

func (c *DBConnector) DeleteChatMessage(id uint) bool {
	if rows, err := gorm.G[ChatMessage](c.DB).Where("id = ?", id).Delete(ctx); err != nil || rows == 0 {
		return false
	}
	return true
}
