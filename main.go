/*
Ethiopian Calendar tool for Golang
Copyright (c) 2022 Yinebeb Tariku <yintar5@gmail.com>
*/

package main

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"gitlab.com/Yinebeb-01/ethiopiandateconverter/ethioGrego"

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

// Ethiopian to Gregorian to handler
func getAdFromEt(ctx *gin.Context) {
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

// Gregorian to Ethiopian handler
func getEtFromAd(ctx *gin.Context) {
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
