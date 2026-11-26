package controller

import (
	"errors"
	"fmt"

	"github.com/TwiN/logr"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/studyphoenix/backend/httputil"
	"github.com/studyphoenix/backend/model"
)

// Request handler to get a user's subjects
//
// @Summary 	Search for subjects or get all
// @Description	Returns a list of user subjects, without any subtopics. Those need to be requested separately
// @Tags		subjects
// @Produce		json
// @Param 		q query string false "The search query (matches name)"
// @Param 		exactMatch query bool false "When set to true (default false), name needs to match exactly for result to be returned"
// @Param 		count query uint false "The maximum number of results to return. Defaults to 25"
// @Param 		offset query uint false "The offset in the results list. Defaults to 0"
// @Param 		minSubtopicCount query uint false "Minimum number of subtopics the subject needs to have. Defaults to 0"
// @Param 		maxSubtopicCount query int false "Maximum number of subtopics the subject needs to have. Defaults to -1 (meaning not limited)"
// @Param 		minAssetCount query uint false "The minimum number of assets the subject and its subtopics need to have. Defaults to 0"
// @Param 		maxAssetCount query uint false "The maximum number of assets the subject and its subtopics need to have. Defaults to -1 (meaning not limited)"
// @Success		200	{object}		[]model.Subject
// @Failure		400	{object}		nil
// @Failure		401	{object}		nil
// @Failure		500	{object}		httputil.HTTPError
// @Router		/subjects/search [get]
func (m *Controller) QuerySubjects(c *gin.Context) {
	var req model.SubjectQuery
	err := c.BindQuery(&req)
	if err != nil {
		logr.Debug(fmt.Sprint("Failed to parse query: ", err))
		c.AbortWithStatus(400)
		return
	}

	uid, ok := sessions.Default(c).Get("userid").(uint)
	if !ok {
		c.AbortWithStatus(401)
		return
	}
	flashcards, err := m.db.QuerySubject(req, uid)
	logr.Debugf("Count: %d", len(flashcards))
	if err != nil {
		httputil.NewError(c, 500, errors.New("Failed to process query request"))
		c.Abort()
	} else {
		c.JSON(200, flashcards)
	}
}
