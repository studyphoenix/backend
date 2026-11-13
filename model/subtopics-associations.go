package model

func (c *DBConnector) AssociateSubtopicToSubtopic(parent uint, child uint) bool {
	subtopic, err := c.GetSubtopic(child)
	if err != nil {
		return false
	}
	parentSubtopic, err := c.GetSubject(parent)
	if err != nil {
		return false
	}

	parentSubtopic.Subtopics = append(parentSubtopic.Subtopics, subtopic)

	c.DB.Save(&parentSubtopic)

	return true
}

func (c *DBConnector) DeassociateSubtopicFromSubtopic(parent uint, child uint) bool {
	parentSubtopic, err := c.GetSubtopic(parent)
	if err != nil {
		return false
	}

	idx := -1
	for i, subtopic := range parentSubtopic.Subtopics {
		if subtopic.ID == child {
			idx = i
			break
		}
	}

	if idx < 0 {
		return false
	}

	parentSubtopic.Subtopics = remove(parentSubtopic.Subtopics, idx)

	c.DB.Save(&parentSubtopic)

	return true
}

// ───────────────────────────────────────────────────────────────────

func (c *DBConnector) AssociateAssetToSubtopic(subtopicid uint, assetid uint) bool {
	asset, err := c.GetAsset(assetid)
	if err != nil {
		return false
	}
	subtopic, err := c.GetSubtopic(subtopicid)
	if err != nil {
		return false
	}

	subtopic.Assets = append(subtopic.Assets, asset)

	c.DB.Save(&subtopic)

	return true
}

func (c *DBConnector) DeassociateAssetFromSubtopic(subtopicid uint, assetid uint) bool {
	subtopic, err := c.GetSubtopic(subtopicid)
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

func (c *DBConnector) AssociateChatToSubtopic(subtopicid uint, chatid uint) bool {
	chat, err := c.GetChat(chatid)
	if err != nil {
		return false
	}
	subtopic, err := c.GetSubtopic(subtopicid)
	if err != nil {
		return false
	}

	subtopic.Chats = append(subtopic.Chats, chat)

	c.DB.Save(&subtopic)

	return true
}

func (c *DBConnector) DeassociateChatFromSubtopic(subtopicid uint, chatid uint) bool {
	subtopic, err := c.GetSubtopic(subtopicid)
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

func (c *DBConnector) AssociateFlashcardToSubtopic(subtopicid uint, flashcardid uint) bool {
	flashcard, err := c.GetFlashcard(flashcardid)
	if err != nil {
		return false
	}
	subtopic, err := c.GetSubtopic(subtopicid)
	if err != nil {
		return false
	}

	subtopic.Flashcards = append(subtopic.Flashcards, flashcard)

	c.DB.Save(&subtopic)

	return true
}

func (c *DBConnector) DeassociateFlashcardFromSubtopic(subtopicid uint, flashcardid uint) bool {
	subtopic, err := c.GetSubtopic(subtopicid)
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
