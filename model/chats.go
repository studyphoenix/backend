package model

import "time"

// TODO: swaggo enum definition for the mode
type Chat struct {
	ID           uint          `json:"id" gorm:"primaryKey" example:"1"`
	Mode         uint          `json:"mode" example:"0"` // the mode to be used for the UI
	Name         string        `json:"name" example:"My nice chat name"`
	UpdatedAt    time.Time     `json:"updatedAt" example:"2026-12-31T08:32:08Z"`
	SubtopicID   uint          `json:"subtopicId" example:"0"` // The id of the subtopic it belongs to
	ChatMessages []ChatMessage `json:"chatMessages" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type ChatMessage struct {
	ID           uint       `json:"id" gorm:"primaryKey" example:"10"`
	ChatID       uint       `json:"chatId" example:"1"`
	CreatedAt    time.Time  `json:"createdAt" example:"2026-12-31T08:32:08Z"`
	EditedAt     *time.Time `json:"editedAt,omitempty" example:"2026-12-31T08:34:12Z"`
	UserMessage  string     `json:"userMessage" example:"Explain why ..."`
	ModelMessage string     `json:"modelMessage" example:"Certianly! This..."`
}
