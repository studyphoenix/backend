package model

import (
	"time"

	"gorm.io/gorm"
)

func (c *DBConnector) GetChat(id uint) (Asset, error) {
	return gorm.G[Asset](c.DB).Where("id = ?", id).First(ctx)
}

func (c *DBConnector) UpdateChat(id uint, req UpdateAssetRequest) bool {
	asset, err := c.GetAsset(id)
	if err != nil {
		return false
	}

	asset.Icon = dereferenceOrDefault(req.Icon, asset.Icon)
	asset.Name = dereferenceOrDefault(req.Name, asset.Name)
	asset.UpdatedAt = time.Now()

	if c.DB.Save(&asset).Error != nil {
		return false
	}
	return true
}

func (c *DBConnector) CreateChat(req CreateAssetRequest) bool {
	asset := Asset{
		Name:      req.Name,
		MediaType: req.MediaType,
		Icon:      req.Icon,
		UpdatedAt: time.Now(),
	}

	if c.DB.Save(&asset).Error != nil {
		return false
	}
	return true
}

func (c *DBConnector) DeleteChat(id uint) bool {
	if rows, err := gorm.G[Chat](c.DB).Where("id = ?", id).Delete(ctx); err != nil || rows == 0 {
		return false
	}
	return true
}
