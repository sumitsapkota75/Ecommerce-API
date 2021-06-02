package infrastructure

import (
	"context"
	"path/filepath"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

// NewFirebaseApp -> creates new firebase app instance
func NewFirebaseApp(logger Logger) *firebase.App {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	serviceAccountKeyFilePath, err := filepath.Abs("./serviceAccountKey.json")
	if err != nil {
		logger.Zap.Panic("Unable to load serviceAccountKey.json file")
	}
	opt := option.WithCredentialsFile(serviceAccountKeyFilePath)

	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		logger.Zap.Fatalf("Firebase app initialization failed: %v", err)
		return nil
	}
	return app

}

// NewFBAuth -> creates new firebase auth client
func NewFBAuth(logger Logger, app *firebase.App) *auth.Client {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	firebaseAuth, err := app.Auth(ctx)
	if err != nil {
		logger.Zap.Fatalf("Firebase Authentication: %v", err)
	}

	return firebaseAuth
}

// NewFCMClient  -> creates new firebase cloud messaging client
func NewFCMClient(logger Logger, app *firebase.App) *messaging.Client {
	messagingClient, err := app.Messaging(context.Background())
	if err != nil {
		logger.Zap.Fatalf("Firebase messaing: %v", err)
	}
	return messagingClient
}


