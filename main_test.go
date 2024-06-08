package main

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yDog-1/woodon_server/ent"
	"github.com/yDog-1/woodon_server/ent/enttest"

	_ "github.com/mattn/go-sqlite3"
)

func Test_CreateUser(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()
	db := NewDBcontext(client, context.Background())
	user, err := db.CreateUser(&ent.User{
		Name:  "test",
		Email: "test@example.com",
	})
	if err != nil {
		t.Fatalf("failed to create user: %v", err)
	}
	assert.Equal(t, "test", user.Name)
	assert.Equal(t, "test@example.com", user.Email)
}

func Test_CreateUser名前が空(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()
	db := NewDBcontext(client, context.Background())
	_, err := db.CreateUser(&ent.User{
		Name:  "",
		Email: "test@example.com",
	})
	assert.NotNil(t, err)
}
func Test_CreateUserメール被り(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()
	db := NewDBcontext(client, context.Background())
	_, err := db.CreateUser(&ent.User{
		Name:  "test",
		Email: "test@example.com",
	})
	if err != nil {
		t.Fatalf("failed to create user: %v", err)
	}
	_, err = db.CreateUser(&ent.User{
		Name:  "test",
		Email: "test@example.com",
	})
	assert.EqualError(t, err, "ent: constraint failed: UNIQUE constraint failed: users.email")
}
