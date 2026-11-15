package model

import "github.com/TwiN/logr"

func (c *DBConnector) AssociateSubtopicToSubtopic(parent uint, child uint, uid uint) bool {
	subtopic, err := c.GetSubtopic(child, uid)
	if err != nil {
		return false
	}
	parentSubtopic, err := c.GetSubject(parent, uid)
	if err != nil {
		return false
	}

	parentSubtopic.Subtopics = append(parentSubtopic.Subtopics, subtopic)

	c.DB.Save(&parentSubtopic)

	return true
}

func (c *DBConnector) DeassociateSubtopicFromSubtopic(parent uint, child uint, uid uint) bool {
	parentSubtopic, err := c.GetSubtopic(parent, uid)
	if err != nil {
		return false
	}

	idx := -1
	for i, subtopic := range parentSubtopic.Children {
		if subtopic.ID == child {
			idx = i
			break
		}
	}

	if idx < 0 {
		return false
	}

	parentSubtopic.Children = remove(parentSubtopic.Children, idx)

	c.DB.Save(&parentSubtopic)

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

	subtopic.Assets = append(subtopic.Assets, asset)

	c.DB.Save(&subtopic)

	return true
}

func (c *DBConnector) DeassociateAssetFromSubtopic(subtopicid uint, assetid uint, uid uint) bool {
	subtopic, err := c.GetSubtopic(subtopicid, uid)
	if err != nil {
		return false
	}

	idx := -1
	for i, asset := range subtopic.Assets {
		if asset.ID == assetid {
			idx = i
			break
		}
	}

	if idx < 0 {
		return false
	}

	subtopic.Assets = remove(subtopic.Assets, idx)

	c.DB.Save(&subtopic)

	return true
}

// ───────────────────────────────────────────────────────────────────

func (c *DBConnector) AssociateChatToSubtopic(subtopicid uint, chatid uint, uid uint) bool {
	chat, err := c.GetChat(chatid, uid)
	if err != nil {
		logr.Debugf("Failed to get chat %e", err)
		return false
	}
	subtopic, err := c.GetSubtopic(subtopicid, uid)
	if err != nil {
		return false
	}

	subtopic.Chats = append(subtopic.Chats, chat)

	c.DB.Save(&subtopic)

	return true
}

func (c *DBConnector) DeassociateChatFromSubtopic(subtopicid uint, chatid uint, uid uint) bool {
	subtopic, err := c.GetSubtopic(subtopicid, uid)
	if err != nil {
		return false
	}

	idx := -1
	for i, chat := range subtopic.Chats {
		if chat.ID == chatid {
			idx = i
			break
		}
	}

	if idx < 0 {
		return false
	}

	subtopic.Chats = remove(subtopic.Chats, idx)

	c.DB.Save(&subtopic)

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

	subtopic.Flashcards = append(subtopic.Flashcards, flashcard)

	c.DB.Save(&subtopic)

	return true
}

func (c *DBConnector) DeassociateFlashcardFromSubtopic(subtopicid uint, flashcardid uint, uid uint) bool {
	subtopic, err := c.GetSubtopic(subtopicid, uid)
	if err != nil {
		return false
	}

	idx := -1
	for i, flashcard := range subtopic.Flashcards {
		if flashcard.ID == flashcardid {
			idx = i
			break
		}
	}

	if idx < 0 {
		return false
	}

	subtopic.Flashcards = remove(subtopic.Flashcards, idx)

	c.DB.Save(&subtopic)

	return true
}
