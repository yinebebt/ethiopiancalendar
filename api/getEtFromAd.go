package api

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/Yinebeb-01/ethiopiandateconverter/ethioGrego"
	"net/http"
	"strconv"
	"strings"
)

// Show EtToAd godoc
// @Summary      Convert date
// @Description  get string by date
// @Tags         EtToAd
// @Accept       json
// @Produce      json
// @Param        date   path      string  true  "date"
// @Success      200  {object}  time.Time
// @Failure      400  {object}  error
// @Failure      404  {object}  error
// @Failure      500  {object}  error
// @Router       /et-from-ad/{date} [get]
// Gregorian to Ethiopian handler
func GetEtFromAd(ctx *gin.Context) {
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
		EtDate, err := ethioGrego.To_ethiopian(year, month, day)
		if err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"response": EtDate.Format("2006-01-02"),
			})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"response": err.Error(),
			})
		}
	}
}
