package controller

type AssociationRequest struct {
	Add bool   `json:"add" example:"true" default:"false"` // whether to add (set to true) or remove (set to false)
	Ids []uint `json:"ids" example:"[5, 10]" default:"[]"` // Ids of the objects to associate
}
