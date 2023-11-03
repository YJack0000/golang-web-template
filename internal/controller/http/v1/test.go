package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"llm_plateform_resourcemanagement/internal/entity"
	"llm_plateform_resourcemanagement/internal/usecase"
	"llm_plateform_resourcemanagement/pkg/logger"
)

type testRoutes struct {
	t usecase.Test
	l logger.Interface
}

func newTestRoutes(handler *gin.RouterGroup, t usecase.Test, l logger.Interface) {
	r := &testRoutes{t, l}

	h := handler.Group("/test")
	{
		h.GET("/history", r.history)
		h.POST("/add", r.add)
	}
}

type historyResponse struct {
	History []entity.Test `json:"history"`
}

// @Summary     Show history
// @Description Show all test history
// @ID          history
// @Tags  	    test
// @Accept      json
// @Produce     json
// @Success     200 {object} historyResponse
// @Failure     500 {object} eResponse
// @Router      /test/history [get]
func (r *testRoutes) history(c *gin.Context) {
	Tests, err := r.t.History(c.Request.Context())
	if err != nil {
		r.l.Error(err, "http - v1 - history")
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	c.JSON(http.StatusOK, historyResponse{Tests})
}

type addRequest struct {
	Test string `json:"test"       example:"test"`
}

// @Summary     add history
// @Description add test history
// @ID          add
// @Tags  	    test
// @Accept      json
// @Produce     json
// @Success     200 {object} sResponse
// @Param    request body addRequest true "request body"
// @Router      /test/add [post]
func (r *testRoutes) add(c *gin.Context) {
	var request addRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		r.l.Error(err, "http - v1 - doTranslate")
		errorResponse(c, http.StatusBadRequest, "invalid request body")

		return
	}

	err := r.t.Add(c.Request.Context(), entity.Test{Test: request.Test})
	if err != nil {
		r.l.Error(err, "http - v1 - add")
		errorResponse(c, http.StatusInternalServerError, "database problems")
	}

	successResponse(c, http.StatusAccepted, "success")
}
