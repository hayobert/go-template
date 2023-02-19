package controller

import (
	"go-template/api/controller/forms"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Index struct {
}

func NewIndex() *Index {
	return &Index{}
}

func (c *Index) Test(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, &forms.TestResp{OK: true})
}
