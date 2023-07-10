package routes

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func Redirect(c *gin.Context, rdb *redis.Client, ctx context.Context) {
    id := c.Param("id")
    url, err := rdb.Get(ctx, fmt.Sprintf("short:%v", id)).Result()

    if err != nil {
        c.JSON(500, gin.H{
            "status":  500,
            "message": "Internal Server Error",
        })

        return
    }

    c.Redirect(307, url)
}
