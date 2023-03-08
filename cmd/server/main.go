package main

import (
	"context"
	"log"

	_ "github.com/aaronland/go-uid-whosonfirst"
	_ "github.com/aaronland/go-uid-proxy"				
	"github.com/aaronland/go-uid-server/app/server"		
)

func main() {

	ctx := context.Background()
	logger := log.Default()

	err := server.Run(ctx, logger)

	if err != nil {
		logger.Fatalf("Failed to run server application, %w", err)
	}
}

