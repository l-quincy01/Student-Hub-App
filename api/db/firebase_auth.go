package db

import (
	"context"
	"log"
	"os"
	"path/filepath"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

type FirebaseAuth struct {
	client *auth.Client
}

func (f *FirebaseAuth) Connect(ctx context.Context) {
	projectId := os.Getenv("PROJECT_ID")
	config := &firebase.Config{
		ProjectID: projectId,
	}

	credentialsJSON, err := os.ReadFile(filepath.Join("..", "..", "services-account.json"))
	if err != nil {
		log.Fatalf("error reading service account file: %v", err)
	}

	opt := option.WithCredentialsJSON(credentialsJSON)
	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	f.client = client
}
