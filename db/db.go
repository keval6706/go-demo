package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"go.demo/ent"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitializeFirebase initializes the Firebase app and returns the Firestore client.
func InitializeDB() (db *gorm.DB) {
	dsn := "host=localhost user=postgres password=root dbname=key_community port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	db.Logger.LogMode(logger.Info)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Postgres Connected.")

	return db
}

func Ent() (*ent.Client, context.Context) {
	ctx := context.Background()
	url := "postgresql://postgres:root@localhost/key_community"
	db, err := sql.Open("pgx", url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Ent Postgres Connected.")

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv)), ctx
}
