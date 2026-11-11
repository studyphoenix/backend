package model

import (
	"time"

	"gorm.io/gorm"
)

type Asset struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	UserID    uint       `json:"-"`
	Name      string     `json:"name" example:"my asset name"`
	UpdatedAt time.Time  `json:"updatedAt" example:"2027-01-01T00:00:00Z"`
	MediaType string     `json:"mediaType" example:"application/pdf"`
	Icon      string     `json:"icon" example:"fa-gear"`
	Subtopics []Subtopic `json:"-" gorm:"many2many:subtopics_assets;"`
	Subjects  []Subject  `json:"-" gorm:"many2many:subjects_assets;"`
}

type CreateAssetRequest struct {
	Name      string `json:"name" example:"Analysis Script" maxLength:"250" minLength:"3" validate:"required,min=3,max=250,printascii"`
	MediaType string `json:"mediaType" example:"application/pdf" minLength:"3" maxLength:"40" validate:"required,contains=/,min=3,max=40,printascii"` // The mediatype (formerly MIME type) of the asset. Will block replacement with asset of another type
	Icon      string `json:"icon" example:"fa-gear" maxLength:"250" minLength:"0" validate:"required,min=3,max=250,printascii"`                       // Must be a valid identifier for vuetify
}

type UpdateAssetRequest struct {
	Name *string `json:"name" example:"Analysis Script" maxLength:"250" minLength:"3" validate:"optional,min=3,max=250,printascii"`
	Icon *string `json:"icon" example:"fa-gear" maxLength:"250" minLength:"0" validate:"optional,min=3,max=250,printascii"` // Must be a valid identifier for vuetify
}

func (c *DBConnector) GetAsset(id uint) (Asset, error) {
	return gorm.G[Asset](c.DB).Where("id = ?", id).First(ctx)
}

func (c *DBConnector) UpdateAsset(id uint, req UpdateAssetRequest) bool {
	asset, err := c.GetAsset(id)
	if err != nil {
		return false
	}

	asset.Icon = dereferenceOrDefault(req.Icon, asset.Icon)
	asset.Name = dereferenceOrDefault(req.Name, asset.Name)

	if c.DB.Save(&asset).Error != nil {
		return false
	}
	return true
}

func (c *DBConnector) CreateAsset(req CreateAssetRequest) bool {
	asset := Asset{
		Name:      req.Name,
		MediaType: req.MediaType,
		Icon:      req.Icon,
	}

	if c.DB.Save(&asset).Error != nil {
		return false
	}
	return true
}

func (c *DBConnector) DeleteAsset(id uint) bool {
	if rows, err := gorm.G[Asset](c.DB).Where("id = ?", id).Delete(ctx); err != nil || rows == 0 {
		return false
	}
	return true
}
