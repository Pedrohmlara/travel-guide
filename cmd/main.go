package main

import (
	"demo-travel-guide/internal/router"
	"demo-travel-guide/internal/utils"
)

func main() {
	utils.CheckEnv()
	router.Serve()
}
