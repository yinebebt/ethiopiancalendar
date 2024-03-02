package api

import (
	"errors"
	"gitlab.com/Yinebeb-01/ethiopiandateconverter/internal/module/ethioGrego"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// Ethiopian : Gregorian to Ethiopian date converter
// @Summary      Gregorian to Ethiopian
// @Description  Get ethiopian date string from gregorian date
// @Tags         Date-Conversion
// @Accept       json
// @Produce      json
// @Param        date	path	string	true	"date"
// @Success      200  {object}  time.Time
// @Failure      400  {object}  error
// @Failure      404  {object}  error
// @Failure      500  {object}  error
// @Router       /ad-to-et/{date} [get]
func Ethiopian(ctx *gin.Context) {
	dateString, state := ctx.Params.Get("date")
	if !state {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "empty date",
		})
		return
	}
	var date = strings.Split(dateString, "-")
	if len(date) > 3 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "not a valid date",
		})
		return
	}

	day, _ := strconv.Atoi(date[2])
	month, _ := strconv.Atoi(date[1])
	year, _ := strconv.Atoi(date[0])
	EtDate, err := ethioGrego.Ethiopian(year, month, day)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"ethiopian_date": EtDate.Format("2006-01-02"),
	})
}

// Gregorian : Ethiopian to Gregorian date converter
// @Summary      Ethiopian to Gregorian
// @Description  Get gregorian date string from ethiopian date
// @Tags         Date-Conversion
// @Accept       json
// @Produce      json
// @Param        date	path	string  true  "date"
// @Success      200  {object}  time.Time
// @Failure      400  {object}  error
// @Failure      404  {object}  error
// @Failure      500  {object}  error
// @Router       /et-to-ad/{date} [get]
func Gregorian(ctx *gin.Context) {
	dateString, state := ctx.Params.Get("date")
	if !state {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "empty date",
		})
		return
	}
	var date = strings.Split(dateString, "-")
	if len(date) > 3 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "not a valid date",
		})
		return
	}

	day, _ := strconv.Atoi(date[2])
	month, _ := strconv.Atoi(date[1])
	year, _ := strconv.Atoi(date[0])

	resDate, err := ethioGrego.Gregorian(year, month, day)
	if err != nil {
		if err.Error() == "not a valid date" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": errors.New("internal server error"),
			})
		}
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"gregorian_date": resDate.Format("2006-01-02"),
	})
}
