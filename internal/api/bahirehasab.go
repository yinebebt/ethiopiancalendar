package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yinebebt/ethiopiancalendar/internal/module/bahirehasab"
)

// BahireHasab
// @Summary      Bahirehasab (ባህረ-ሐሳብ)
// @Description  Get years fasting and festival date
// @Tags         Bahirehasab
// @Accept       json
// @Produce      json
// @Param        year	path	string	true	"year"  example("2016")
// @Success      200  {object}  bahirehasab.Festival
// @Failure      400  {object}  error
// @Failure      404  {object}  error
// @Failure      500  {object}  error
// @Router       /bahire-hasab/{year} [get]
func BahireHasab(ctx *gin.Context) {
	yearString, state := ctx.Params.Get("year")
	if !state {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Empty year value",
		})
		return
	}
	year, err := strconv.Atoi(yearString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "not a valid year",
		})
		return
	}

	festival, err := bahirehasab.BahireHasab(year)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": festival,
	})
}
