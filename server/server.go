package server

import (
	"os"
    "context"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/vclemenzi/safeshort/routes"
)

func Server() {
    ctx := context.Background()
    r := gin.Default()
    rdb := redis.NewClient(&redis.Options{
        Addr:     os.Getenv("REDIS_URL"),
        Password: os.Getenv("REDIS_PASSWORD"),
        DB:       0,
    })

    r.StaticFile("/", "./public/index.html")

    r.GET("/:id", func (c *gin.Context) { routes.Redirect(c, rdb, ctx) })

    r.POST("/api/short", func (c *gin.Context) { routes.Short(c, rdb, ctx) })

    r.Run()
}
