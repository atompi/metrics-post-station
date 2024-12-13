package router

import (
	selfmetrics "github.com/atompi/go-kits/metrics/handler"
	metricsapi "github.com/atompi/metrics-post-station/pkg/apis/metrics/v1"
	"github.com/atompi/metrics-post-station/pkg/options"
	"github.com/gin-gonic/gin"
)

type RouterGroupFunc func(*gin.RouterGroup, options.Options)

func MetricsRouter(routerGroup *gin.RouterGroup, opts options.Options) {
	routerGroup.GET(opts.APIServer.Metrics.Path, selfmetrics.NewPromHandler())
}

func ApisRouter(routerGroup *gin.RouterGroup, opts options.Options) {
	apisGroup := routerGroup.Group(opts.APIServer.Prefix)

	metricsapi.Router(apisGroup, opts)
}

func Register(e *gin.Engine, opts options.Options) {
	routerGroupFuncs := []RouterGroupFunc{}

	if opts.APIServer.Metrics.Enable {
		e.Use(selfmetrics.Handler(""))
		routerGroupFuncs = append(routerGroupFuncs, MetricsRouter)
	}

	routerGroupFuncs = append(
		routerGroupFuncs,
		ApisRouter,
	)

	rootRouterGroup := e.Group("/")

	for _, routerGroup := range routerGroupFuncs {
		routerGroup(rootRouterGroup, opts)
	}
}
