package model

import "gorm.io/gorm"

type Settings struct {
	ID                   uint   `json:"-" gorm:"primaryKey"`
	UserID               uint   `json:"-"`
	AnswerLength         string `json:"answerLength" enums:"very-short,short,default,long,very-long" default:"default"`
	AnswerInBulletPoints bool   `json:"answerInBulletPoints" default:"false"`
}

type UpdateUserConfigRequest struct {
	AnswerLength         *string `json:"answerLength" enums:"very-short,short,default,long,very-long" default:"default" validate:"optional"`
	AnswerInBulletPoints *bool   `json:"answerInBulletPoints" default:"false" validate:"optional"`
}

func (c *DBConnector) CreateSettings(userid uint) bool {
	settings := Settings{
		UserID:               userid,
		AnswerLength:         "default",
		AnswerInBulletPoints: false,
	}
	return c.DB.Save(&settings).Error != nil
}

func (c *DBConnector) GetSettings(userid uint) (Settings, error) {
	return gorm.G[Settings](c.DB).Where("user_id = ?", userid).First(ctx)
}

func (c *DBConnector) StoreSettings(userid uint, req UpdateUserConfigRequest) bool {
	settings, err := c.GetSettings(userid)
	if err != nil {
		return false
	}

	settings.AnswerInBulletPoints = dereferenceOrDefault(req.AnswerInBulletPoints, settings.AnswerInBulletPoints)
	settings.AnswerLength = dereferenceOrDefault(req.AnswerLength, settings.AnswerLength)

	return c.DB.Save(&settings).Error != nil
}
