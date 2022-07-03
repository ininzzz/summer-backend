package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	InitConfigure()
	r := gin.Default()
	register(r)
	r.Run()
}
