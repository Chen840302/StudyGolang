package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"joey/src/api"
)


func Router() {
	r := gin.Default()
	// group第二個參數可以放api, 在這group底下的url都需要經過這個middleware, 用來做id認證
	r.GET("/", func(c *gin.Context){
		fmt.Println("OK")
	})

	v1 := r.Group("/v1")
	v1.GET("/person", api.GetPerson)
	v1.POST("/person", api.AddPerson)



	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func Test()  {
	fmt.Println("ddddddd")
}