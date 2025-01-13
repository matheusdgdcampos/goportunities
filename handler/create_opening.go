package handler

import (
	"goportunities/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}
	ctx.BindJSON(&request)
	err := request.Validate()
	if err != nil {
		logger.Errorf("validation error: %v", err)
		sendError(ctx, http.StatusBadRequest, err.Error())
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
		logger.Errorf("error on create new opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error on create new opening and register on database")
		return
	}

	sendSuccess(ctx, "create_opening", opening)
}
