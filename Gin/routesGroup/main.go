package main

import (
	"fmt"
	"gin/routesGroup/routes"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
	"os"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("temp/*")
	routes.AdminRoutesInit(r)
	routes.ApiRoutesInit(r)
	routes.DefaultRoutesInit(r)
	routes.ArticleRoutesInit(r)

	cfg, err := ini.Load("../conf/app.ini")

	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	// 典型读取操作，默认分区可以使用空字符串表示
	fmt.Println("App Mode:", cfg.Section("").Key("app_mode").String())
	fmt.Println("Data Path:", cfg.Section("paths").Key("data").String())

	r.Run()

}
