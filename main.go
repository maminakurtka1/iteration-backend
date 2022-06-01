package main

import (
	"context"
	// "iteration-backend/config"
	"iteration-backend/database"
)

// cfg := config.LoadConfig()

func main() {
	ctx := context.Background()
	db := database.OpenConnect(ctx)
	setupRouter(db)
}
