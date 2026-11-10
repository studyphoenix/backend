package model

type Settings struct {
	ID     uint `json:"-" gorm:"primaryKey"`
	UserID uint `json:"-"`
}
