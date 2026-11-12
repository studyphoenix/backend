package model

func (c *DBConnector) AsociateSubtopicToSubtopic(parent uint, child uint) bool {
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

func (c *DBConnector) DeasociateSubtopicFromSubtopic(parent uint, child uint) bool {
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

	return true
}
