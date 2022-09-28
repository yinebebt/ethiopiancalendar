package main

import (
	"ethiopianDateConverter/ethioGrego"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", homePage)
	router.GET("/ad-from-et/:date", getAdFromEt)
	router.GET("/et-from-ad/:date", getEtFromAd)

	router.Run(":8080")
}

func homePage(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"Message": "EthioGrego Server",
	})
}

func getAdFromEt(ctx *gin.Context) {
	dateString, state := ctx.Params.Get("date")
	if state {
		dateString = strings.TrimPrefix(dateString, "date=")
	}
	var splitedDate = strings.Split(dateString, "-")
	if len(splitedDate) < 3 || len(splitedDate) > 3 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": "not a valid date",
		})
	} else {
		day, _ := strconv.Atoi(splitedDate[2])
		month, _ := strconv.Atoi(splitedDate[1])
		year, _ := strconv.Atoi(splitedDate[0])

		date := ethioGrego.To_gregorian(year, month, day)

		ctx.JSON(http.StatusOK, gin.H{
			"Date in Gregorian": date.Format("2006-01-02"),
		})
	}
}

func getEtFromAd(ctx *gin.Context) {
	dateString, state := ctx.Params.Get("date")
	if state {
		dateString = strings.TrimPrefix(dateString, "date=")
	}
	var splitedDate = strings.Split(dateString, "-")
	if len(splitedDate) < 3 || len(splitedDate) > 3 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": "not a valid date",
		})
	} else {
		day, _ := strconv.Atoi(splitedDate[2])
		month, _ := strconv.Atoi(splitedDate[1])
		year, _ := strconv.Atoi(splitedDate[0])
		EtDate := ethioGrego.To_ethiopian(year, month, day)
		ctx.JSON(http.StatusOK, gin.H{
			"Date in Ethiopian": EtDate.Format("2006-01-02"),
		})
	}
}
