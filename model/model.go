package model

import (
	"context"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBConnector struct {
	DB *gorm.DB
}

var ctx = context.Background()

func Init(prod bool) *DBConnector {
	var conf gorm.Config
	if prod {
		conf = gorm.Config{}
	} else {
		conf = gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		}
	}

	db, err := gorm.Open(sqlite.Open("janishutz-store.db"), &conf)
	if err != nil {
		log.Fatal("[DB] Failed to connect to DB with error", err)
	}

	// FIXME: Only used if using sqlite db
	db.Exec("PRAGMA foreign_keys = ON")

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Subtopic{})
	db.AutoMigrate(&Subject{})
	db.AutoMigrate(&Asset{})
	db.AutoMigrate(&Chat{})
	db.AutoMigrate(&ChatMessage{})
	db.AutoMigrate(&Flashcard{})

	return &DBConnector{
		DB: db,
	}
}

func dereferenceOrDefault[T any](val *T, defaultValue T) T {
	if val != nil {
		return *val
	} else {
		return defaultValue
	}
}
