package model

type User struct {
	ID         uint        `json:"id" gorm:"primaryKey"`
	Uid        string      `json:"-"` // UserID as specified by OIDC OP
	Email      string      `json:"email"`
	Username   string      `json:"username"`
	Subjects   []Subject   `json:"-" gorm:"constraint:OnDelete:CASCADE"`
	Subtopics  []Subtopic  `json:"-" gorm:"constraint:OnDelete:CASCADE"`
	Assets     []Asset     `json:"-" gorm:"constraint:OnDelete:CASCADE"`
	Flashcards []Flashcard `json:"-" gorm:"constraint:OnDelete:CASCADE"`
	Chats      []Chat      `json:"-" gorm:"constraint:OnDelete:CASCADE"`
}
