package model

import (
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/TwiN/logr"
	"gorm.io/gorm"
)

func (c *DBConnector) GetAsset(id uint, uid uint) (Asset, error) {
	return gorm.G[Asset](c.DB).Where("id = ? AND user_id = ?", id, uid).First(ctx)
}

func (c *DBConnector) UpdateAsset(id uint, req UpdateAssetRequest, uid uint) bool {
	asset, err := c.GetAsset(id, uid)
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

func (c *DBConnector) CreateAsset(req CreateAssetRequest, userid uint) (uint, bool) {
	asset := Asset{
		Name:      req.Name,
		MediaType: req.MediaType,
		Icon:      req.Icon,
		UserID:    userid,
		UpdatedAt: time.Now(),
	}

	result := gorm.WithResult()
	err := gorm.G[Asset](c.DB, result).Create(ctx, &asset)
	if err != nil {
		return 0, false
	}
	return asset.ID, true
}

func (c *DBConnector) DeleteAsset(id uint, uid uint) bool {
	asset, err := c.GetAsset(id, uid)
	if err != nil {
		logr.Error(err.Error())
		return false
	}

	c.DB.Exec("DELETE FROM `subtopics_assets` WHERE asset_id = ?; DELETE FROM `subjects_assets` WHERE asset_id = ?", id, id)
	if rows, err := gorm.G[Asset](c.DB).Where("id = ?", id).Delete(ctx); err != nil || rows == 0 {
		return false
	}

	os.Remove(filepath.Join(c.Config.Datadir, strconv.FormatUint(uint64(id), 10), MEDIA_TYPE_MAP[asset.MediaType]))

	return true
}
