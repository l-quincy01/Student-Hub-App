package main

import (
	"context"
	"os"
	"student-hub-app/config"
	"student-hub-app/internal/controller"
	"student-hub-app/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func init() {
	config.SetupEnv()
}

func main() {
	ctx := context.Background()

	client := config.SetupMongoDb(ctx)
	defer config.CloseMognoDb(ctx, client)

	conf := config.SetupGoogleOauth(ctx)

	app := fiber.New()

	googleAuthService := service.NewGoogleAuthService(conf)
	authController := controller.NewAuthController(googleAuthService)

	authController.SetupRoutes(app)

	port := os.Getenv("PORT")

	app.Use(logger.New())
	app.Listen(port)
}
