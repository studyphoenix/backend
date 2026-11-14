package controller

import (
	"fmt"
	"strconv"

	"github.com/TwiN/logr"
	"github.com/gin-gonic/gin"
	"github.com/studyphoenix/backend/model"
)

// Request handler to associate subtopic(s)
//
// @Summary 	associate subtopic(s)
// @Description	associates the subtopic(s) to this subtopic
// @Tags		subtopics
// @Produce		json
// @Param 		id path string true "The id of the subtopic to update"
// @Param 		payload body model.AssociationRequest true "The associations to add"
// @Success		204	{object}		nil
// @Failure 	400 {object} 		nil
// @Failure		401	{object}		nil
// @Failure 	404 {object} 		nil "subtopic or subtopic(s) id not found"
// @Failure		500	{object}		httputil.HTTPError
// @Router		/subtopics/associate/subtopics/{id} [patch]
func (m *Controller) AssociateSubtopicToSubtopic(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	var req model.AssociationRequest
	err = c.BindJSON(&req)
	if err != nil {
		logr.Debug(fmt.Sprint("Failed to parse body to create association: ", err))
		c.AbortWithStatus(400)
		return
	}

	for _, assetId := range req {
		m.db.AssociateSubtopicToSubtopic(uint(id), assetId)
	}

	c.JSON(204, nil)
}

// Request handler to deassociate subtopic(s)
//
// @Summary 	deassociate subtopic(s)
// @Description	deassociates the subtopic(s) to this subtopic
// @Tags		subtopics
// @Produce		json
// @Param 		id path string true "The id of the subtopic to update"
// @Param 		payload body model.AssociationRequest true "The associations to remove"
// @Success		204	{object}		nil
// @Failure 	400 {object} 		nil
// @Failure		401	{object}		nil
// @Failure 	404 {object} 		nil "subtopic or subtopic(s) id not found"
// @Failure		500	{object}		httputil.HTTPError
// @Router		/subtopics/deassociate/subtopics/{id} [patch]
func (m *Controller) DeassociateSubtopicFromSubtopic(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	var req model.AssociationRequest
	err = c.BindJSON(&req)
	if err != nil {
		logr.Debug(fmt.Sprint("Failed to parse body to create association: ", err))
		c.AbortWithStatus(400)
		return
	}

	for _, assetId := range req {
		m.db.DeassociateSubtopicFromSubtopic(uint(id), assetId)
	}

	c.JSON(204, nil)
}

// Request handler to associate asset(s)
//
// @Summary 	associate asset(s)
// @Description	associates the asset(s) depending on the value of "add" to this subtopic
// @Tags		subtopics
// @Produce		json
// @Param 		id path string true "The id of the subtopic to update"
// @Param 		payload body model.AssociationRequest true "The associations to add"
// @Success		204	{object}		nil
// @Failure 	400 {object} 		nil
// @Failure		401	{object}		nil
// @Failure 	400 {object} 		nil "subtopic or asset id not found"
// @Failure		500	{object}		httputil.HTTPError
// @Router		/subtopics/associate/assets/{id} [patch]
func (m *Controller) AssociateAssetToSubtopic(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	var req model.AssociationRequest
	err = c.BindJSON(&req)
	if err != nil {
		logr.Debug(fmt.Sprint("Failed to parse body to create association: ", err))
		c.AbortWithStatus(400)
		return
	}

	for _, assetId := range req {
		m.db.AssociateAssetToSubtopic(uint(id), assetId)
	}

	c.JSON(204, nil)
}

// Request handler to deassociate asset(s)
//
// @Summary 	deassociate asset(s)
// @Description	deassociates the asset(s) from this subtopic
// @Tags		subtopics
// @Produce		json
// @Param 		id path string true "The id of the subtopic to update"
// @Param 		payload body model.AssociationRequest true "The associations to remove"
// @Success		204	{object}		nil
// @Failure 	400 {object} 		nil
// @Failure		401	{object}		nil
// @Failure 	400 {object} 		nil "subtopic or asset id not found"
// @Failure		500	{object}		httputil.HTTPError
// @Router		/subtopics/deassociate/assets/{id} [patch]
func (m *Controller) DeassociateAssetFromSubtopic(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	var req model.AssociationRequest
	err = c.BindJSON(&req)
	if err != nil {
		logr.Debug(fmt.Sprint("Failed to parse body to create association: ", err))
		c.AbortWithStatus(400)
		return
	}

	for _, assetId := range req {
		m.db.DeassociateAssetFromSubject(uint(id), assetId)
	}

	c.JSON(204, nil)
}

// Request handler to associate chat(s)
//
// @Summary 	associate chat(s)
// @Description	associates the chat(s) to this subtopic
// @Tags		subtopics
// @Produce		json
// @Param 		id path string true "The id of the subtopic to update"
// @Param 		payload body model.AssociationRequest true "The associations to add"
// @Success		204	{object}		nil
// @Failure 	400 {object} 		nil
// @Failure		401	{object}		nil
// @Failure 	400 {object} 		nil "subtopic or chat id not found"
// @Failure		500	{object}		httputil.HTTPError
// @Router		/subtopics/associate/chats/{id} [patch]
func (m *Controller) AssociateChatToSubtopic(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	var req model.AssociationRequest
	err = c.BindJSON(&req)
	if err != nil {
		logr.Debug(fmt.Sprint("Failed to parse body to create association: ", err))
		c.AbortWithStatus(400)
		return
	}

	for _, assetId := range req {
		m.db.AssociateChatToSubtopic(uint(id), assetId)
	}

	c.JSON(204, nil)
}

// Request handler to deassociate chat(s)
//
// @Summary 	deassociate chat(s)
// @Description	deassociates the chat(s) from this subtopic
// @Tags		subtopics
// @Produce		json
// @Param 		id path string true "The id of the subtopic to update"
// @Param 		payload body model.AssociationRequest true "The associations to remove"
// @Success		204	{object}		nil
// @Failure 	400 {object} 		nil
// @Failure		401	{object}		nil
// @Failure 	400 {object} 		nil "subtopic or flashcard id not found"
// @Failure		500	{object}		httputil.HTTPError
// @Router		/subtopics/deassociate/chats/{id} [patch]
func (m *Controller) DeassociateChatFromSubtopic(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	var req model.AssociationRequest
	err = c.BindJSON(&req)
	if err != nil {
		logr.Debug(fmt.Sprint("Failed to parse body to create association: ", err))
		c.AbortWithStatus(400)
		return
	}

	for _, assetId := range req {
		m.db.DeassociateChatFromSubtopic(uint(id), assetId)
	}

	c.JSON(204, nil)
}

// Request handler to associate flashcard(s)
//
// @Summary 	associate flashcard(s)
// @Description	associates the flashcard(s) to this subtopic
// @Tags		subtopics
// @Produce		json
// @Param 		id path string true "The id of the subtopic to update"
// @Param 		payload body model.AssociationRequest true "The associations to add"
// @Success		204	{object}		nil
// @Failure 	400 {object} 		nil
// @Failure		401	{object}		nil
// @Failure 	400 {object} 		nil "subtopic or flashcard id not found"
// @Failure		500	{object}		httputil.HTTPError
// @Router		/subtopics/associate/flashcards/{id} [patch]
func (m *Controller) AssociateFlashcardToSubtopic(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	var req model.AssociationRequest
	err = c.BindJSON(&req)
	if err != nil {
		logr.Debug(fmt.Sprint("Failed to parse body to create association: ", err))
		c.AbortWithStatus(400)
		return
	}

	for _, assetId := range req {
		m.db.AssociateFlashcardToSubtopic(uint(id), assetId)
	}

	c.JSON(204, nil)
}

// Request handler to deassociate flashcards(s)
//
// @Summary 	deassociate chat(s)
// @Description	deassociates the chat(s) from this subtopic
// @Tags		subtopics
// @Produce		json
// @Param 		id path string true "The id of the subtopic to update"
// @Param 		payload body model.AssociationRequest true "The associations to remove"
// @Success		204	{object}		nil
// @Failure 	400 {object} 		nil
// @Failure		401	{object}		nil
// @Failure 	400 {object} 		nil "subtopic or flashcard id not found"
// @Failure		500	{object}		httputil.HTTPError
// @Router		/subtopics/deassociate/flashcards/{id} [patch]
func (m *Controller) DeassociateFlashcardFromSubtopic(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	var req model.AssociationRequest
	err = c.BindJSON(&req)
	if err != nil {
		logr.Debug(fmt.Sprint("Failed to parse body to create association: ", err))
		c.AbortWithStatus(400)
		return
	}

	for _, assetId := range req {
		m.db.DeassociateFlashcardFromSubtopic(uint(id), assetId)
	}

	c.JSON(204, nil)
}
