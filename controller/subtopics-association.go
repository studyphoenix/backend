package controller

import "github.com/gin-gonic/gin"

// Request handler to associate subtopic(s)
//
// @Summary 	associate subtopic(s) [WIP]
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
	c.JSON(200, nil)
}

// Request handler to deassociate subtopic(s)
//
// @Summary 	deassociate subtopic(s) [WIP]
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
func (m *Controller) DeassociateSubtopicToSubtopic(c *gin.Context) {
	c.JSON(200, nil)
}

// Request handler to associate asset(s)
//
// @Summary 	associate asset(s) [WIP]
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
	c.JSON(200, nil)
}

// Request handler to deassociate asset(s)
//
// @Summary 	deassociate asset(s) [WIP]
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
func (m *Controller) DeassociateAssetToSubtopic(c *gin.Context) {
	c.JSON(200, nil)
}

// Request handler to associate chat(s)
//
// @Summary 	associate chat(s) [WIP]
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
	c.JSON(200, nil)
}

// Request handler to deassociate chat(s)
//
// @Summary 	deassociate chat(s) [WIP]
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
func (m *Controller) DeassociateChatToSubtopic(c *gin.Context) {
	c.JSON(200, nil)
}

// Request handler to associate flashcard(s)
//
// @Summary 	associate flashcard(s) [WIP]
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
	c.JSON(200, nil)
}

// Request handler to deassociate flashcards(s)
//
// @Summary 	deassociate chat(s) [WIP]
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
func (m *Controller) DeassociateFlashcardToSubtopic(c *gin.Context) {
	c.JSON(200, nil)
}
