package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/vclemenzi/safeshort/utils"
)

type Data struct {
    Url string `json:"url"`
}

func Short(c *gin.Context, rdb *redis.Client, ctx context.Context) {
    jsonBody, err := ioutil.ReadAll(c.Request.Body)
    
    if err != nil {
		c.JSON(500, gin.H{
			"status":  500,
			"message": "Internal Server Error: Unable to read the body",
		})

		return
	}

    var data Data
	json.Unmarshal(jsonBody, &data)

    isMalicius := utils.CheckURL(data.Url)

    if isMalicius {
        c.JSON(200, gin.H{
		    "status": 452,
            "message": "Blocked for security reasons: This url has been blocked as dangerous",
	    })

        return
    }

    id := utils.GenerateID()

    rdb.Set(ctx, fmt.Sprintf("short:%v", id), data.Url, 0)

    c.JSON(200, gin.H{
		"status":  200,
		"message": "Redirect successfully created",
        "url":     c.Request.Host + "/" + id,
	})
}
