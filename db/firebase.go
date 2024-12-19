package database

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

// InitializeFirebase initializes the Firebase app and returns the Firestore client.
func Firebase() (*firestore.Client, context.Context) {
	ctx := context.Background()
	sa := option.WithCredentialsFile("google-cloud.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Connected To Firestore")

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	return client, ctx
}
