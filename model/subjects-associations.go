package model

func (c *DBConnector) AssociateAssetToSubject(subjectid uint, assetid uint, uid uint) bool {
	asset, err := c.GetAsset(assetid, uid)
	if err != nil {
		return false
	}
	subject, err := c.GetSubject(subjectid, uid)
	if err != nil {
		return false
	}

	if err := c.DB.Model(&subject).Association("Assets").Append(&asset); err != nil {
		return false
	}

	return true
}

func (c *DBConnector) DeassociateAssetFromSubject(subjectid uint, assetid uint, uid uint) bool {
	subject, err := c.GetSubject(subjectid, uid)
	if err != nil {
		return false
	}
	asset, err := c.GetAsset(assetid, uid)
	if err != nil {
		return false
	}

	if err := c.DB.Model(&subject).Association("Assets").Delete(&asset); err != nil {
		return false
	}

	return true
}
