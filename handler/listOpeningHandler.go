package handler

import (
	"net/http"

	"github.com/Ricardo-Cardozo/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func ListOpeningsHandler(c *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(
			c,
			http.StatusInternalServerError,
			"error listing openings",
		)

		return
	}

	sendSuccess(c, "LIST_OPENINGS", openings)
}
