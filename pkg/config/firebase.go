package config

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
	"path/filepath"
)

var FirebaseApp *firebase.App

func InitFirebase() {
	ctx := context.Background()

	serviceAccountKeyFilePath, err := filepath.Abs("./serviceAccountKey.json")

	if err != nil {
		panic("Unable to load serviceAccountKeys.json file")
	}

	opt := option.WithCredentialsFile(serviceAccountKeyFilePath)

	app, err := firebase.NewApp(ctx, nil, opt)

	if err != nil {
		panic("Firebase load error")
	}

	FirebaseApp = app
}
