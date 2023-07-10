package routes

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func Redirect(c *gin.Context, rdb *redis.Client, ctx context.Context) {
    c.Redirect(308, "https://www.google.com")
}
