package weinvite

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func HandlerPersonPost(c *gin.Context) {
	var data Person
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(200, gin.H{
			"message": "Error Get Body",
		})

	}
	if err := data.Insert(context.Background()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
	}

	c.JSON(200, gin.H{
		"message": "data Sukses di tambahkan",
		"data":    data,
	})

}

func HandlerGetPerson(c *gin.Context) {
	path := strings.Split(string(c.Request.URL.Path), "/")
	fmt.Println(path)
}
