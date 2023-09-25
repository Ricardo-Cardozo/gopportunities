package handler

import (
	"net/http"

	"github.com/Ricardo-Cardozo/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(c *gin.Context) {
	request := CreateOpeningRequest{}

	c.BindJSON(&request)

	logger.Infof("request received: %+v", request)

	if err := request.Validate(); err != nil {
		logger.Errorf("VALIDATION_ERROR: %v", err)

		sendError(c, http.StatusBadRequest, err.Error())

		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("ERROR_CREATING_OPENING: %v", err)

		sendError(c, http.StatusInternalServerError, err.Error())

		return
	}

	sendSuccess(c, "CREATE_OPENING", opening)
}
