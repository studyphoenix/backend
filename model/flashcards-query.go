package model

import "gorm.io/gorm"

func (c *DBConnector) FlashcardQuery() ([]Flashcard, error) {
	filter := ""
	return gorm.G[Flashcard](c.DB).Where(filter).Find(ctx)
}
