package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main()  {
	r := gin.Default()

	r.GET("/someJSON", func(c *gin.Context){
		data := map[string]interface{}{
			"lang":"GO语言",
			"tag":"<br>",
		}

		// 输出 : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
		c.AsciiJSON(http.StatusOK,data)
	})

	r.Run(":8080")
}
