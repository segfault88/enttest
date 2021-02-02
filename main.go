package main

import (
	"context"
	"fmt"
	"log"

	"github.com/davecgh/go-spew/spew"
	_ "github.com/mattn/go-sqlite3"
	"github.com/segfault88/enttest/ent"
)

func main() {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	u, err := createUser(context.Background(), client)
	if err != nil {
		log.Fatalf("failed creating user: %v", err)
	}

	spew.Dump(u)
}

func createUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.Create().Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %v", err)
	}
	log.Printf("user was created id: %#v", u.ID)
	return u, nil
}
