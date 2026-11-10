package model

type Settings struct {
	ID                   uint   `json:"-" gorm:"primaryKey"`
	UserID               uint   `json:"-"`
	AnswerLength         string `json:"answerLength" enums:"very-short,short,default,long,very-long"`
	AnswerInBulletPoints bool   `json:"answerInBulletPoints"`
}
