package api

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gitlab.com/Yinebeb-01/ethiopiandateconverter/internal/ethioGrego"
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
	var splitedDate = strings.Split(dateString, "-")
	if len(splitedDate) > 3 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"response": "not a valid date",
		})
	} else {
		day, _ := strconv.Atoi(splitedDate[2])
		month, _ := strconv.Atoi(splitedDate[1])
		year, _ := strconv.Atoi(splitedDate[0])

		date, err := ethioGrego.To_gregorian(year, month, day)
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
