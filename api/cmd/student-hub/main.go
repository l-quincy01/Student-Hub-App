package main

import (
	"context"
	"student-hub-app/config"
)

func init() {
	config.SetupEnv()
}

func main() {
	ctx := context.Background()

	client := config.SetupMongoDb(ctx)
	defer config.CloseMognoDb(ctx, client)
}
