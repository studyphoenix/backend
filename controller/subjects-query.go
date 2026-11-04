package controller

import "github.com/gin-gonic/gin"

// Request handler to get a user's subjects
//
// @Summary 	Search for subjects or get all [WIP]
// @Description	Returns a list of user subjects, without any subtopics. Those need to be requested separately
// @Tags		subjects
// @Produce		json
// @Param 		q query string false "The search query (matches name)"
// @Param 		count query uint false "The maximum number of results to return. Defaults to 25"
// @Param 		offset query uint false "The offset in the results list. Defaults to 0"
// @Param 		subtopic query bool false "Whether or not the subject must have a subtopic. Defaults to false"
// @Param 		assetCount query uint false "The minimum number of assets the subject and its subtopics need to have. Defaults to 0"
// @Success		200	{object}		[]model.Subject
// @Failure		401	{object}		nil
// @Failure		500	{object}		httputil.HTTPError
// @Router		/subjects/search [get]
func (m *Controller) QuerySubjects(c *gin.Context) {
	c.JSON(200, nil)
}
