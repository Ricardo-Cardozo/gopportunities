package handler

import (
	"net/http"

	"github.com/Ricardo-Cardozo/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func UpdateOpeningHandler(c *gin.Context) {
	request := UpdateOpeningRequest{}

	c.BindJSON(&request)

	if err := request.Validate(); err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

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
			"OPENING_NOT_FOUND",
		)

		return
	}

	if request.Role != "" {
		opening.Role = request.Role
	}

	if request.Company != "" {
		opening.Company = request.Company
	}

	if request.Location != "" {
		opening.Location = request.Location
	}

	if request.Remote != nil {
		opening.Remote = *request.Remote
	}

	if request.Link != "" {
		opening.Link = request.Link
	}

	if request.Salary > 0 {
		opening.Salary = request.Salary
	}

	if err := db.Save(&opening).Error; err != nil {
		sendError(
			c,
			http.StatusInternalServerError,
			"ERROR_UPDATING_OPENING",
		)
	}

	sendSuccess(c, "UPDATE_OPENING", opening)
}
