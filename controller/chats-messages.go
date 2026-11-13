package controller

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/TwiN/logr"
	"github.com/gin-gonic/gin"
	"github.com/hci-group-15/backend/httputil"
	"github.com/hci-group-15/backend/model"
)

// Request handler to create a new subject
//
// @Summary 	Add a new chat message to the chat with id `{id}`
// @Description	Creates a new subject linked to the current user
// @Tags		chats
// @Produce		json
// @Param 		id path string true "The id of the subject to retrieve"
// @Param 		payload body model.AddChatMessageRequest true "Subject details to store"
// @Success		201	{object}		nil
// @Failure 	400 {object} 		nil
// @Failure		401	{object}		nil
// @Failure		500	{object}		httputil.HTTPError
// @Router		/chats/messages/add/{id} [put]
func (m *Controller) AddChatMessageToChat(c *gin.Context) {
	chatId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	var req model.AddChatMessageRequest
	err = c.BindJSON(&req)
	if err != nil {
		logr.Debug(fmt.Sprint("Failed to parse body to create chat message: ", err))
		c.AbortWithStatus(400)
		return
	}

	if m.db.CreateChatMessage(req, uint(chatId)) {
		c.JSON(201, nil)
	} else {
		httputil.NewError(c, 500, errors.New("ERR_ADD_CHAT_MESSAGE"))
		c.Abort()
	}
}

// Request handler to update a chat message
//
// @Summary 	Updates the specified chat message
// @Description	Updates the chat message. Note that all attributes in the body are optional and if omitted, will not be updated
// @Tags		chats
// @Produce		json
// @Param 		id path string true "The id of the chat message to update"
// @Param 		payload body model.UpdateChatMessageRequest true "Message content to update. All attributes optional!"
// @Success		204	{object}		nil
// @Failure 	400 {object} 		nil
// @Failure		401	{object}		nil
// @Failure 	404 {object} 		nil "message not found"
// @Failure		500	{object}		httputil.HTTPError
// @Router		/chats/messages/update/{id} [patch]
func (m *Controller) UpdateChatMessage(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	var req model.UpdateChatMessageRequest
	err = c.BindJSON(&req)
	if err != nil {
		logr.Debug(fmt.Sprint("Failed to parse body to update chat message, error: ", err))
		c.Abort()
		return
	}

	if m.db.UpdateChatMessage(uint(id), req) {
		c.JSON(204, nil)
	} else {
		c.JSON(404, nil)
	}
}

// Request handler to delete a delete a chat message
//
// @Summary 	Deletes the given subject
// @Description	Deletes the given subject including all subtopics, assets and flashcards.
// @Tags		chats
// @Produce		json
// @Param 		id path string true "The id of the chat message to delete"
// @Success		204	{object}		nil
// @Failure		401	{object}		nil
// @Failure 	404 {object} 		nil "chat or message not found"
// @Failure		500	{object}		httputil.HTTPError
// @Router		/chats/messages/delete/{id} [delete]
func (m *Controller) DeleteChatMessage(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	if m.db.DeleteChatMessage(uint(id)) {
		c.JSON(204, nil)
	} else {
		c.JSON(404, nil)
	}
}
