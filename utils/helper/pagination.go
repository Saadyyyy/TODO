package helper

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func Pagination(ctx *gin.Context) (int, int) {
	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil || page < 1 {
		page = 1
	}

	perPage, err := strconv.Atoi(ctx.Query("perPage"))
	if err != nil || perPage < 1 {
		perPage = 1
	}

	return page, perPage
}
