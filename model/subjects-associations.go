package model

func remove[T any](slice []T, s int) []T {
	return append(slice[:s], slice[s+1:]...)
}

func (c *DBConnector) AsociateAssetToSubject(subjectid uint, assetid uint) bool {
	asset, err := c.GetAsset(assetid)
	if err != nil {
		return false
	}
	subject, err := c.GetSubject(subjectid)
	if err != nil {
		return false
	}

	subject.Assets = append(subject.Assets, asset)

	c.DB.Save(&subject)

	return true
}

func (c *DBConnector) DeasociateAssetFromSubject(subjectid uint, assetid uint) bool {
	subject, err := c.GetSubject(subjectid)
	if err != nil {
		return false
	}

	idx := -1
	for i, asset := range subject.Assets {
		if asset.ID == assetid {
			idx = i
			break
		}
	}

	if idx < 0 {
		return false
	}

	subject.Assets = remove(subject.Assets, idx)

	return true
}

func (c *DBConnector) AsociateSubtopicToSubject(subjectid uint, subtopicid uint) bool {
	subtopic, err := c.GetSubtopic(subtopicid)
	if err != nil {
		return false
	}
	subject, err := c.GetSubject(subjectid)
	if err != nil {
		return false
	}

	subject.Subtopics = append(subject.Subtopics, subtopic)

	c.DB.Save(&subject)

	return true
}

func (c *DBConnector) DeasociateSubtopicFromSubject(subjectid uint, subtopicid uint) bool {
	subject, err := c.GetSubject(subjectid)
	if err != nil {
		return false
	}

	idx := -1
	for i, subtopic := range subject.Subtopics {
		if subtopic.ID == subtopicid {
			idx = i
			break
		}
	}

	if idx < 0 {
		return false
	}

	subject.Subtopics = remove(subject.Subtopics, idx)

	return true
}
