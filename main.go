package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	_ "github.com/mattn/go-sqlite3"
	"github.com/segfault88/enttest/ent"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	gofakeit.Seed(time.Now().UnixNano())
	client, err := ent.Open("sqlite3", "file:ent.sqlite?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	count := rand.Int() % 100
	log.Printf("creating: %d users", count)
	for i := 0; i < count; i++ {
		_, err := createUser(context.Background(), client)
		if err != nil {
			log.Fatalf("failed creating user: %v", err)
		}
	}
}

func createUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	var age int
	for age <= 0 {
		age = rand.Int() % 100
	}
	u, err := client.User.Create().
		SetAge(age).
		SetName(gofakeit.Name()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %v", err)
	}
	log.Printf("user was created id: %s", u)
	return u, nil
}
