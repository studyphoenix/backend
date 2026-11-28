package model

import "github.com/TwiN/logr"

func (c *DBConnector) AssociateSubtopicToSubtopic(parent uint, child uint, uid uint) bool {
	subtopic, err := c.GetSubtopic(child, uid)
	if err != nil {
		return false
	}
	parentSubtopic, err := c.GetSubtopic(parent, uid)
	if err != nil {
		return false
	}

	if err := c.DB.Model(&parentSubtopic).Association("Children").Append(&subtopic); err != nil {
		return false
	}

	return true
}

func (c *DBConnector) DeassociateSubtopicFromSubtopic(parent uint, child uint, uid uint) bool {
	parentSubtopic, err := c.GetSubtopic(parent, uid)
	if err != nil {
		return false
	}
	childSubtopic, err := c.GetSubtopic(child, uid)
	if err != nil {
		return false
	}

	if err := c.DB.Model(&parentSubtopic).Association("Children").Delete(&childSubtopic); err != nil {
		return false
	}

	return true
}

// ───────────────────────────────────────────────────────────────────

func (c *DBConnector) AssociateAssetToSubtopic(subtopicid uint, assetid uint, uid uint) bool {
	asset, err := c.GetAsset(assetid, uid)
	if err != nil {
		return false
	}
	subtopic, err := c.GetSubtopic(subtopicid, uid)
	if err != nil {
		return false
	}

	if err := c.DB.Model(&subtopic).Association("Assets").Append(&asset); err != nil {
		return false
	}

	return true
}

func (c *DBConnector) DeassociateAssetFromSubtopic(subtopicid uint, assetid uint, uid uint) bool {
	subtopic, err := c.GetSubtopic(subtopicid, uid)
	if err != nil {
		return false
	}
	asset, err := c.GetAsset(assetid, uid)
	if err != nil {
		return false
	}

	if err := c.DB.Model(&subtopic).Association("Assets").Delete(&asset); err != nil {
		return false
	}

	return true
}

// ───────────────────────────────────────────────────────────────────

func (c *DBConnector) AssociateChatToSubtopic(subtopicid uint, chatid uint, uid uint) bool {
	chat, err := c.GetChat(chatid, uid)
	if err != nil {
		logr.Debugf("Failed to get chat %w", err)
		return false
	}
	subtopic, err := c.GetSubtopic(subtopicid, uid)
	if err != nil {
		return false
	}

	if err := c.DB.Model(&subtopic).Association("Chats").Append(&chat); err != nil {
		return false
	}

	return true
}

func (c *DBConnector) DeassociateChatFromSubtopic(subtopicid uint, chatid uint, uid uint) bool {
	subtopic, err := c.GetSubtopic(subtopicid, uid)
	if err != nil {
		return false
	}
	chat, err := c.GetChat(chatid, uid)
	if err != nil {
		return false
	}

	if err := c.DB.Model(&subtopic).Association("Chats").Delete(&chat); err != nil {
		return false
	}

	return true
}

// ───────────────────────────────────────────────────────────────────

func (c *DBConnector) AssociateFlashcardToSubtopic(subtopicid uint, flashcardid uint, uid uint) bool {
	flashcard, err := c.GetFlashcard(flashcardid, uid)
	if err != nil {
		return false
	}
	subtopic, err := c.GetSubtopic(subtopicid, uid)
	if err != nil {
		return false
	}

	if err := c.DB.Model(&subtopic).Association("Flashcards").Append(&flashcard); err != nil {
		return false
	}

	return true
}

func (c *DBConnector) DeassociateFlashcardFromSubtopic(subtopicid uint, flashcardid uint, uid uint) bool {
	subtopic, err := c.GetSubtopic(subtopicid, uid)
	if err != nil {
		return false
	}
	flashcard, err := c.GetFlashcard(flashcardid, uid)
	if err != nil {
		return false
	}

	if err := c.DB.Model(&subtopic).Association("Flashcards").Delete(&flashcard); err != nil {
		return false
	}

	return true
}
