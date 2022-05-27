package main

import (
	"iteration-backend/config"
)

func main() {
	config.LoadConfig()
	setupRouter()
}
