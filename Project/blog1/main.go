package main

import (
	routers "blog/routes"
)

func main() {
	r := routers.InitRouter()

	r.Run()
}
