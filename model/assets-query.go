package model

import (
	"errors"

	"gorm.io/gorm"
)

type AssetQuery struct {
	Offset             *uint   `form:"offset" binding:"omitempty,min=5,max=50"` // OFFSET
	Limit              *uint   `form:"count" binding:"omitempty,min=5,max=50"`  // LIMIT
}

func (c *DBConnector) QueryAssets(query AssetQuery, uid uint) ([]Asset, error) {
	sql := gorm.G[Asset](c.DB).Where("user_id = ?", uid)

	// Offset
	if query.Offset != nil {
		offset := int(*query.Offset)
		if offset < 0 {
			return nil, errors.New("invalid offset specified")
		}
		sql = sql.Offset(offset)
	}

	// Limit
	limit := 25
	if query.Limit != nil {
		limit = int(*query.Limit)
	}

	if limit < 5 || limit > 50 {
		limit = 25
	}

	return sql.Limit(limit).Find(ctx)
}
