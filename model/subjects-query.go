package model

import (
	"errors"

	"gorm.io/gorm"
)

type SubjectQuery struct {
	Name             *string `form:"q" binding:"omitempty,printascii"`
	ExactMatch       *bool   `form:"exactMatch" binding:"omitempty,boolean"`
	Offset           *uint   `form:"offset" binding:"omitempty,min=5,max=50"` // OFFSET
	Limit            *uint   `form:"count" binding:"omitempty,min=5,max=50"`  // LIMIT
	MinSubtopicCount *uint   `form:"minSubtopicCount" binding:"omitempty,numeric,min=0"`
	MaxSubtopicCount *uint   `form:"maxSubtopicCount" binding:"omitempty,numeric,min=0"`
	MinAssetCount    *uint   `form:"minAssetCount" binding:"omitempty,numeric,min=0"`
	MaxAssetCount    *uint   `form:"maxAssetCount" binding:"omitempty,numeric,min=0"`
}

func (c *DBConnector) QuerySubject(query SubjectQuery, uid uint) ([]Subject, error) {
	sql := gorm.G[Subject](c.DB).Where("user_id = ?", uid)

	// Name
	if query.Name != nil {
		if query.ExactMatch != nil && !*query.ExactMatch {
			sql = sql.Where("name = ?", *query.Name)
		} else {
			sql = sql.Where("name LIKE ?", "%"+*query.Name+"%")
		}
	}

	// Subtopic count
	if query.MaxSubtopicCount != nil || query.MinSubtopicCount != nil {
		sq := c.DB.Select("COUNT(id) as cnt, subject_id").Group("subject_id").Table("subtopics")

		sq2 := c.DB.Select("subject_id")
		if query.MaxSubtopicCount != nil {
			sq2.Where("cnt < ?", *query.MaxSubtopicCount+1)
		}
		if query.MinSubtopicCount != nil {
			sq2.Where("cnt > ?", *query.MinSubtopicCount-1)
		}
		sq2 = sq2.Table("(?)", sq)

		sql = sql.Where("id = (?)", sq2)
	}

	// Asset count
	if query.MaxAssetCount != nil || query.MinAssetCount != nil {
		sq := c.DB.Select("COUNT(asset_id) as cnt, subject_id").Group("subject_id").Table("subjects_assets")

		sq2 := c.DB.Select("subject_id")
		if query.MaxAssetCount != nil {
			sq2.Where("cnt < ?", *query.MaxAssetCount+1)
		}
		if query.MinAssetCount != nil {
			sq2.Where("cnt > ?", *query.MinAssetCount-1)
		}
		sq2 = sq2.Table("(?)", sq)

		sql = sql.Where("id = (?)", sq2)
	}

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
