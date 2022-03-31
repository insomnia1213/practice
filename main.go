package main

import (
	"github.com/kataras/iris"
)

func main() {
	ap := iris.New()
	ap.Run(iris.Addr(":8080"))
}
