package controller

import (
	"github.com/gin-gonic/gin"
)

type CreateChatRequest struct {
	Name string `example:"Analysis" json:"name" maxLength:"250" minLength:"3"`
	Icon string `example:"fa-gear" json:"icon" maxLength:"250" minLength:"0"` // The Icon of the chat. Must be an icon that can be rendered in the browser
}

// Request handler to create a new chat
//
// @Summary 	Create a new chat [WIP]
// @Description	Creates a new chat linked to the current user
// @Tags		chats
// @Produce		json
// @Param 		payload body CreateChatRequest true "chat details to store"
// @Success		201	{object}		nil
// @Failure 	400 {object} 		nil
// @Failure		401	{object}		nil
// @Failure		500	{object}		httputil.HTTPError
// @Router		/chats/create [put]
func (m *Controller) CreateChat(c *gin.Context) {
	c.JSON(200, nil)
}

// Request handler to get a chat
//
// @Summary 	Get the specified chat [WIP]
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
	c.JSON(200, nil)
}

type UpdateChatRequest struct {
	Name *string `example:"Analysis" json:"name" maxLength:"250" minLength:"3"`
	Icon *string `example:"fa-gear" json:"icon" maxLength:"250" minLength:"0"` // The Icon of the chat. Must be an icon that can be rendered in the browser
}

// Request handler to update a chat
//
// @Summary 	Updates the specified chat [WIP]
// @Description	Updates the chat. Note that all attributes in the body are optional and if omitted, will not be updated
// @Tags		chats
// @Produce		json
// @Param 		id path string true "The id of the chat to update"
// @Param 		payload body UpdateChatRequest true "Chat details to update with. All attributes optional!"
// @Success		200	{object}		nil
// @Failure 	400 {object} 		nil
// @Failure		401	{object}		nil
// @Failure 	404 {object} 		nil "chat not found"
// @Failure		500	{object}		httputil.HTTPError
// @Router		/chats/update/{id} [patch]
func (m *Controller) UpdateChat(c *gin.Context) {
	c.JSON(200, nil)
}

// Request handler to delete a chat
//
// @Summary 	Deletes the given chat [WIP]
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
	c.JSON(200, nil)
}
