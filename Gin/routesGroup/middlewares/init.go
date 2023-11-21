package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

//var wg sync.WaitGroup

func InitMiddleware(context *gin.Context) {
	// 判断用户是否登录
	fmt.Println(time.Now())
	fmt.Println(context.Request.URL)

	context.Set("username", "张三")

	cCp := context.Copy()
	// 定义一个 goroutine 统计日志
	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("Done! in path" + cCp.Request.URL.Path)
	}()
}
