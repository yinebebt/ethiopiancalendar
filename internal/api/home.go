package api

import (
	"github.com/gin-gonic/gin"
)

// HomePage : root path
// @Summary		Home page
// @Description	Home Page
// @Tags	Date-Converter
// @Produce      html
// @Success      200  {object}  map[string]any
// @Failure      400  {object}  error
// @Failure      404  {object}  error
// @Failure      500  {object}  error
// @Router       / [get]
func HomePage(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/html")
	ctx.File("internal/assets/index.html")
}
