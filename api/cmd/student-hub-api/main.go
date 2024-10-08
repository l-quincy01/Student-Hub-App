package main

import (
	"context"
	"os"
	"student-hub-app/config"
	"student-hub-app/db"
	"student-hub-app/internal/controller"
	"student-hub-app/internal/repository"
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

	mongoDB := db.NewMongoDB()
	if err := mongoDB.Connect(ctx); err != nil {
		panic(err)
	}

	defer mongoDB.Disconnect(ctx)

	userManager := repository.NewUserManager(mongoDB)
	userCreator := service.NewUserCreator(userManager)

	cfg := config.SetupGoogleOauth(ctx)
	oauthService := service.NewOAuthService(cfg)

	authController := controller.NewAuthController(oauthService, userCreator)
	authController.Router(app)

	app.Use(logger.New())

	port := os.Getenv("PORT")
	app.Listen(port)
}
