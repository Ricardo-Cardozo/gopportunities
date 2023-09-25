package handler

import (
	"fmt"
	"net/http"

	"github.com/Ricardo-Cardozo/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteOpeningHandler(c *gin.Context) {
	id := c.Query("id")

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
			fmt.Sprintf("opening with id: %s not found", id),
		)

		return
	}

	if err := db.Delete(&opening, id).Error; err != nil {
		sendError(
			c,
			http.StatusInternalServerError,
			fmt.Sprintf("error delete opening with id: %s", id),
		)

		return
	}

	sendSuccess(c, "DELETE_OPENING", opening)
}
