package model

import (
	"os"
	"path/filepath"
	"strconv"
	"time"

	"gorm.io/gorm"
)

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
	asset.UpdatedAt = time.Now()

	if c.DB.Save(&asset).Error != nil {
		return false
	}
	return true
}

func (c *DBConnector) CreateAsset(req CreateAssetRequest, userid uint) bool {
	asset := Asset{
		Name:      req.Name,
		MediaType: req.MediaType,
		Icon:      req.Icon,
		UserID:    userid,
		UpdatedAt: time.Now(),
	}

	if c.DB.Save(&asset).Error != nil {
		return false
	}
	return true
}

func (c *DBConnector) DeleteAsset(id uint) bool {
	asset, err := c.GetAsset(id)
	if err != nil {
		return false
	}
	if rows, err := gorm.G[Asset](c.DB).Where("id = ?", id).Delete(ctx); err != nil || rows == 0 {
		return false
	}

	os.Remove(filepath.Join(c.Config.Datadir, strconv.FormatUint(uint64(id), 10), MEDIA_TYPE_MAP[asset.MediaType]))

	return true
}
