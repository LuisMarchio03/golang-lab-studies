package handler

import (
	"fmt"
	"net/http"

	"github.com/LuisMarchio03/golang-boilerplate-api/schemas"
	"github.com/gin-gonic/gin"
)

// @BaseParh /api/v1

// @Summary Delete Opening
// @Description Delete a job opening
// @Tags Opening
// @Accept json
// @Producer json
// @Param id query string true "Opening identification"
// @Success 200 {object} DeleteOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening [delete]
func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamsIsRequired("id",
			"queryParameter").Error())
		return
	}
	opening := schemas.Opening{}

	// find opening
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf(
			"opening with id: %s not found",
			id,
		))
		return
	}

	// delete opening
	if err := db.Delete(&opening).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting opening with id: %s", id))
		return
	}

	sendSuccess(ctx, "delete-opening", opening)
}
