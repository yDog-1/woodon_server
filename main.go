package main

import (
	"context"

	"github.com/yDog-1/woodon_server/ent"
)

type DBcontext struct {
	client *ent.Client
	ctx    context.Context
}

func NewDBcontext(client *ent.Client, ctx context.Context) *DBcontext {
	return &DBcontext{client: client, ctx: ctx}
}

func (db *DBcontext) CreateUser(u *ent.User) (*ent.User, error) {
	user, err := db.client.User.
		Create().
		SetName(u.Name).
		SetEmail(u.Email).
		Save(db.ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}
