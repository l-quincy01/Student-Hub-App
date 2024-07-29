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
	config.LoadEnv()
}

func main() {
	ctx := context.Background()
	app := fiber.New()

	client := config.ConnectMongodb(ctx)
	defer config.DisconnectMongodb(ctx, client)

	cfg := config.SetupGoogleOauth(ctx)
	oauthService := service.NewOAuthService(cfg)
	authController := controller.NewOAuthController(oauthService)
	authController.Router(app)

	app.Use(logger.New())

	port := os.Getenv("PORT")
	app.Listen(port)
}
