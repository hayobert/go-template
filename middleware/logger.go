package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		logger := logrus.WithFields(logrus.Fields{
			"Method": ctx.Request.Method,
			"URI":    ctx.Request.RequestURI,
			"Host":   ctx.Request.Host,
			"trace":  ctx.Request.Header.Get("X-B3-TraceId"),
		})

		ctx.Set("logger", logger)
		ctx.Next()
	}
}
