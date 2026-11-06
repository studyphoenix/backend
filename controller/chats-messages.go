package controller

import (
	"github.com/gin-gonic/gin"
)

type AddChatMessageRequest struct {
	UserMessage  string `example:"User's message" json:"name" maxLength:"20000" minLength:"3" validate:"required,min=3,max=20000"`
	ModelMessage string `example:"Model's message" json:"icon" validate:"required"`
	Summary      string `example:"test messages" json:"summary" validate:"required,min=0,max=250" maxLength:"250" minLength:"0"`
}

// Request handler to create a new subject
//
// @Summary 	Add a new chat message to the chat with id `{id}` [WIP]
// @Description	Creates a new subject linked to the current user
// @Tags		chats
// @Produce		json
// @Param 		id path string true "The id of the subject to retrieve"
// @Param 		payload body CreateSubjectRequest true "Subject details to store"
// @Success		201	{object}		nil
// @Failure 	400 {object} 		nil
// @Failure		401	{object}		nil
// @Failure		500	{object}		httputil.HTTPError
// @Router		/chats/messages/add/{id} [put]
func (m *Controller) AddChatMessageToChat(c *gin.Context) {
	c.JSON(200, nil)
}

type UpdateChatMessageRequest struct {
	UserMessage  string `example:"User's message" json:"name" maxLength:"20000" minLength:"3" validate:"optional,min=3,max=20000"`
	ModelMessage string `example:"Model's message" json:"icon" validate:"optional"`
	Summary      string `example:"test messages" json:"summary" validate:"optional,min=0,max=250" maxLength:"250" minLength:"0"`
}

// Request handler to update a chat message
//
// @Summary 	Updates the specified chat message [WIP]
// @Description	Updates the chat message. Note that all attributes in the body are optional and if omitted, will not be updated
// @Tags		chats
// @Produce		json
// @Param 		id path string true "The id of the chat message to update"
// @Param 		payload body UpdateSubjectRequest true "Message content to update. All attributes optional!"
// @Success		200	{object}		nil
// @Failure 	400 {object} 		nil
// @Failure		401	{object}		nil
// @Failure 	404 {object} 		nil "message not found"
// @Failure		500	{object}		httputil.HTTPError
// @Router		/chats/messages/update/{id} [patch]
func (m *Controller) UpdateChatMessage(c *gin.Context) {
	c.JSON(200, nil)
}

// Request handler to delete a delete a chat message
//
// @Summary 	Deletes the given subject [WIP]
// @Description	Deletes the given subject including all subtopics, assets and flashcards.
// @Tags		chats
// @Produce		json
// @Param 		id path string true "The id of the chat message to delete"
// @Success		200	{object}		nil
// @Failure		401	{object}		nil
// @Failure 	404 {object} 		nil "chat or message not found"
// @Failure		500	{object}		httputil.HTTPError
// @Router		/chats/messages/delete/{id} [delete]
func (m *Controller) DeleteChatMessage(c *gin.Context) {
	c.JSON(200, nil)
}
