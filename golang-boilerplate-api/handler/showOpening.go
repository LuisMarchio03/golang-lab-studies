package handler

import (
	"net/http"

	"github.com/LuisMarchio03/golang-boilerplate-api/schemas"
	"github.com/gin-gonic/gin"
)

// @BaseParh /api/v1

// @Summary Show Opening
// @Description Show a job opening
// @Tags Opening
// @Accept json
// @Producer json
// @Param id query string true "Opening identification"
// @Success 200 {object} ShowOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening [get]
func ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamsIsRequired("id",
			"queryParameter").Error())
		return
	}
	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "opening not found")
		return
	}

	sendSuccess(ctx, "show-opening", opening)
}
