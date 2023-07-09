package server

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/vclemenzi/safeshort/routes"
)

func Server() {
    r := gin.Default()
    rdb := redis.NewClient(&redis.Options{
        Addr:     os.Getenv("REDIS_URL"),
        Password: os.Getenv("REDIS_PASSWORD"),
        DB:       0,
    })

    r.GET("/:id", func (c *gin.Context) { routes.Redirect(c, rdb) })
    
    r.POST("/api/short", func (c *gin.Context) { routes.Redirect(c, rdb) })

    r.Run()
}
