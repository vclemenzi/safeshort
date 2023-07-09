package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func Redirect(c *gin.Context, rdb *redis.Client) {
    c.Redirect(308, "https://www.google.com")
}
