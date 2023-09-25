package handler

import (
	"net/http"

	"github.com/Ricardo-Cardozo/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func ShowOpeningHandler(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		sendError(
			c,
			http.StatusBadRequest,
			errParamIsRequired("id", "queryParameter").Error(),
		)

		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(
			c,
			http.StatusNotFound,
			id,
		)

		return
	}

	sendSuccess(c, "SHOW_OPENING", opening)
}
