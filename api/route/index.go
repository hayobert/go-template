package route

import (
	"go-template/api/controller"

	"github.com/gin-gonic/gin"
)

func InitIndex(g *gin.Engine) {
	index := controller.NewIndex()
	group := g.Group("index")
	{
		group.GET("index", index.Test)
	}
}
