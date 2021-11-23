package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func BootApp() {
	mapUrls()

	fmt.Println("Booting up the API ...")

	router.Run(":8080")
}
