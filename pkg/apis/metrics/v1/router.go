package v1

import (
	"github.com/atompi/metrics-post-station/pkg/handler"
	"github.com/atompi/metrics-post-station/pkg/options"

	"github.com/gin-gonic/gin"
)

func Router(routerGroup *gin.RouterGroup, opts options.Options) {
	MetricsGroup := routerGroup.Group("/metrics")
	{
		MetricsGroup.GET("", handler.NewHandler(GetHandler, opts))
		MetricsGroup.POST("/job/:job/instance/:instance", handler.NewHandler(SetHandler, opts))
	}
}
