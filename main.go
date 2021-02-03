package main

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	_ "github.com/mattn/go-sqlite3"
	"github.com/segfault88/enttest/ent"
	"github.com/segfault88/enttest/ent/user"
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

	// createUsers(context.Background(), client)
	queryUsers(context.Background(), client)
}

func createUsers(ctx context.Context, client *ent.Client) {
	count := rand.Int() % 100
	log.Printf("creating: %d users", count)
	for i := 0; i < count; i++ {
		createUser(context.Background(), client)
	}
}

func createUser(ctx context.Context, client *ent.Client) *ent.User {
	var age int
	for age <= 0 {
		age = rand.Int() % 100
	}
	u, err := client.User.Create().
		SetAge(age).
		SetName(gofakeit.Name()).
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating user: %v", err)
	}
	log.Printf("user was created id: %s", u)
	return u
}

func queryUsers(ctx context.Context, client *ent.Client) {
	users, err := client.User.Query().
		Where(user.And(user.AgeGT(18), user.AgeLT(60))).
		Order(ent.Asc(user.FieldAge)).
		All(context.Background())
	if err != nil {
		log.Fatalf("failed getting users: %v", err)
	}
	log.Printf("Found: %d users", len(users))
	for _, u := range users {
		log.Printf("found: %s", u)
	}
}
