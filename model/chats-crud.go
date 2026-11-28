package model

import (
	"time"

	"github.com/TwiN/logr"
	"gorm.io/gorm"
)

func (c *DBConnector) GetChat(id uint, uid uint) (Chat, error) {
	return gorm.G[Chat](c.DB).Where("id = ? AND user_id = ?", id, uid).Preload("ChatMessages", nil).First(ctx)
}

func (c *DBConnector) UpdateChat(id uint, req UpdateChatRequest, uid uint) bool {
	chat, err := c.GetChat(id, uid)
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

func (c *DBConnector) DeleteChat(id uint, uid uint) bool {
	_, err := c.GetChat(id, uid)
	if err != nil {
		return false
	}

	err = c.DB.Transaction(func(tx *gorm.DB) error {
		c := DBConnector{
			Config: c.Config,
			DB:     tx,
		}

		return c.deleteChatWithoutTransaction(id, uid)
	})

	if err != nil {
		logr.Errorf("[DELETE CHATS] Delete transaction failed: %w", err)
		return false
	} else {
		return true
	}
}

func (c *DBConnector) deleteChatWithoutTransaction(id uint, uid uint) error {
	c.DB.Exec("DELETE FROM `subtopics_chats` WHERE chat_id = ?", id)

	if rows, err := gorm.G[Chat](c.DB).Where("id = ?", id).Delete(ctx); err != nil || rows == 0 {
		return err
	}

	return nil
}
