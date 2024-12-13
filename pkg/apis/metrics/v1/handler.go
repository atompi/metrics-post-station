package v1

import (
	"io"
	"net/http"

	"github.com/atompi/metrics-post-station/pkg/handler"
	redisutil "github.com/atompi/metrics-post-station/pkg/util/redis"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetHandler(c *handler.Context) {
	opts := c.Options
	rdb := redisutil.New(opts.APIServer.Redis)
	defer rdb.Close()

	m := c.GinContext.Query("module")
	t := c.GinContext.Query("target")
	if m == "" || t == "" {
		c.GinContext.JSON(http.StatusOK, gin.H{"response": "bad request, param module or target not found"})
		return
	}
	key := m + "__" + t

	res, err := Get(rdb, key, opts.APIServer.Redis.DialTimeout)
	if err != nil {
		c.GinContext.JSON(http.StatusInternalServerError, gin.H{"response": "get key: " + key + " value failed: " + err.Error()})
		zap.S().Errorf("get key: %v failed: %v", key, err)
		return
	}
	c.GinContext.String(http.StatusOK, res)
}

func SetHandler(c *handler.Context) {
	opts := c.Options
	rdb := redisutil.New(opts.APIServer.Redis)
	defer rdb.Close()

	job := c.GinContext.Param("job")
	instance := c.GinContext.Param("instance")
	key := job + "__" + instance

	bodyBytes, err := io.ReadAll(c.GinContext.Request.Body)
	if err != nil {
		c.GinContext.JSON(http.StatusBadRequest, gin.H{"response": "bad request, failed to read request body"})
		return
	}
	value := string(bodyBytes)

	err = Set(rdb, key, value, opts.APIServer.Redis.DialTimeout)
	if err != nil {
		c.GinContext.JSON(http.StatusInternalServerError, gin.H{"response": "put key: " + key + " value failed: " + err.Error()})
		zap.S().Errorf("put key: %v failed: %v", key, err)
		return
	}
	c.GinContext.String(http.StatusOK, "")
}
