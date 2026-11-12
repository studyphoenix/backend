package controller

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/TwiN/logr"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/hci-group-15/backend/httputil"
	"github.com/hci-group-15/backend/model"
)

// Request handler to create a new chat
//
// @Summary 	Create a new chat
// @Description	Creates a new chat linked to the current user
// @Tags		chats
// @Produce		json
// @Param 		payload body model.CreateChatRequest true "chat details to store"
// @Success		201	{object}		nil
// @Failure 	400 {object} 		nil
// @Failure		401	{object}		nil
// @Failure		500	{object}		httputil.HTTPError
// @Router		/chats/create [put]
func (m *Controller) CreateChat(c *gin.Context) {
	var req model.CreateChatRequest
	err := c.BindJSON(&req)
	if err != nil {
		logr.Debug(fmt.Sprint("Failed to parse body to create chat: ", err))
		c.AbortWithStatus(400)
		return
	}

	uid, ok := sessions.Default(c).Get("userid").(uint)
	if !ok {
		c.AbortWithStatus(401)
		return
	}
	if m.db.CreateChat(req, uid) {
		c.JSON(200, nil)
	} else {
		httputil.NewError(c, 500, errors.New("ERR_CREATE_CHAT"))
	}
}

// Request handler to get a chat
//
// @Summary 	Get the specified chat
// @Description	Returns the specified chat details, including extra metadata
// @Tags		chats
// @Produce		json
// @Param 		id path string true "The id of the chat to retrieve"
// @Success		200	{object}		model.Chat
// @Failure		401	{object}		nil
// @Failure 	404 {object} 		nil
// @Failure		500	{object}		httputil.HTTPError
// @Router		/chats/get/{id} [get]
func (m *Controller) GetChat(c *gin.Context) {
	id, err := strconv.ParseUint(c.Query("id"), 10, 32)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	chat, err := m.db.GetChat(uint(id))
	if err != nil {
		c.AbortWithStatus(404)
		return
	}
	c.JSON(200, chat)
}

// Request handler to update a chat
//
// @Summary 	Updates the specified chat
// @Description	Updates the chat. Note that all attributes in the body are optional and if omitted, will not be updated
// @Tags		chats
// @Produce		json
// @Param 		id path string true "The id of the chat to update"
// @Param 		payload body model.UpdateChatRequest true "Chat details to update with. All attributes optional!"
// @Success		200	{object}		nil
// @Failure 	400 {object} 		nil
// @Failure		401	{object}		nil
// @Failure 	404 {object} 		nil "chat not found"
// @Failure		500	{object}		httputil.HTTPError
// @Router		/chats/update/{id} [patch]
func (m *Controller) UpdateChat(c *gin.Context) {
	id, err := strconv.ParseUint(c.Query("id"), 10, 32)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	var req model.UpdateChatRequest
	err = c.BindJSON(&req)
	if err != nil {
		logr.Debug(fmt.Sprint("Failed to parse body to update chat, error: ", err))
		c.Abort()
		return
	}

	if m.db.UpdateChat(uint(id), req) {
		c.JSON(200, nil)
	} else {
		c.JSON(404, nil)
	}
}

// Request handler to delete a chat
//
// @Summary 	Deletes the given chat
// @Description	Deletes the given chat including all subtopics, assets and flashcards.
// @Tags		chats
// @Produce		json
// @Param 		id path string true "The id of the chat to delete"
// @Success		200	{object}		nil
// @Failure		401	{object}		nil
// @Failure 	404 {object} 		nil "chat not found"
// @Failure		500	{object}		httputil.HTTPError
// @Router		/chats/delete/{id} [delete]
func (m *Controller) DeleteChat(c *gin.Context) {
	id, err := strconv.ParseUint(c.Query("id"), 10, 32)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	if m.db.DeleteChat(uint(id)) {
		c.JSON(200, nil)
	} else {
		c.JSON(404, nil)
	}
}
