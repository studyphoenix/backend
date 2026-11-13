package model

import "time"

type Chat struct {
	ID           uint          `json:"id" gorm:"primaryKey" example:"1"`
	UserID       uint          `json:"-"`
	Mode         uint          `json:"mode" example:"1" enums:"1,2,3"` // the mode to be used for the UI
	Name         string        `json:"name" example:"My nice chat name"`
	UpdatedAt    time.Time     `json:"updatedAt" example:"2026-12-31T08:32:08Z"`
	ChatMessages []ChatMessage `json:"chatMessages" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Subtopics    []Subtopic    `json:"-" gorm:"many2many:subtopics_chats;"`
}

// FIXME: May have to do many2many manually to be able to use constraints properly for consistency

type ChatMessage struct {
	ID           uint      `json:"id" gorm:"primaryKey" example:"10"`
	ChatID       uint      `json:"chatId" example:"1"`
	UpdatedAt    time.Time `json:"updatedAt" example:"2026-12-31T08:34:12Z"`
	UserMessage  string    `json:"userMessage" example:"Explain why ..."`
	ModelMessage string    `json:"modelMessage" example:"Certianly! This..."`
	Summary      string    `json:"summary" example:"summary for this message pair"`
	// TODO: What other fields?
}

type CreateChatRequest struct {
	Name string `example:"Analysis Explanation" json:"name" maxLength:"100" minLength:"3" validate:"required" binding:"required,min=3,max=100,printascii"`
	Mode uint   `example:"1" json:"mode" enums:"1,2,3" validate:"required" binding:"required,min=1,max=3"`
}

type UpdateChatRequest struct {
	Name *string `example:"Analysis Explanation" json:"name" maxLength:"100" minLength:"3" validate:"optional" binding:"omitempty,min=3,max=100,printascii"`
	Mode *uint   `example:"2" json:"mode" enums:"1,2,3" validate:"optional" binding:"omitempty,min=1,max=3"`
}

type AddChatMessageRequest struct {
	UserMessage  string `example:"User's message" json:"userMessage" maxLength:"20000" minLength:"3" validate:"required" binding:"required,min=3,max=20000"`
	ModelMessage string `example:"Model's message" json:"modelMessage" validate:"required" binding:"required"`
	Summary      string `example:"test messages" json:"summary" validate:"required" binding:"required,min=0,max=250" maxLength:"250" minLength:"0"`
}

type UpdateChatMessageRequest struct {
	UserMessage  *string `example:"User's message" json:"userMessage" maxLength:"20000" minLength:"3" validate:"optional" binding:"omitempty,min=3,max=20000"`
	ModelMessage *string `example:"Model's message" json:"modelMessage" validate:"optional"`
	Summary      *string `example:"test messages" json:"summary" validate:"optional" binding:"omitempty,min=0,max=250" maxLength:"250" minLength:"0"`
}
