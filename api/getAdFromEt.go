package api

import (
	"errors"
	"github.com/Yinebeb-01/ethiopiandateconverter/ethioGrego"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

// Show AdToEt godoc
// @Summary      Convert date
// @Description  get string by date
// @Tags         AdToEt
// @Accept       json
// @Produce      json
// @Param        date   path     string  true  "date"
// @Success      200  {object}  time.Time
// @Failure      400  {object}  error
// @Failure      404  {object}  error
// @Failure      500  {object}  error
// @Router       /ad-from-et/{date} [get]
// Ethiopian to Gregorian to handler
func GetAdFromEt(ctx *gin.Context) {
	dateString, state := ctx.Params.Get("date")
	if state {
		dateString = strings.TrimPrefix(dateString, "date=")
	}
	var splitDate = strings.Split(dateString, "-")
	if len(splitDate) > 3 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"response": "not a valid date",
		})
	} else {
		day, _ := strconv.Atoi(splitDate[2])
		month, _ := strconv.Atoi(splitDate[1])
		year, _ := strconv.Atoi(splitDate[0])

		date, err := ethioGrego.ToGregorian(year, month, day)
		if err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"response": date.Format("2006-01-02"),
			})
		}
		if err.Error() == "not a valid date" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"response": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"response": errors.New("internal server error"),
			})
		}
	}
}
