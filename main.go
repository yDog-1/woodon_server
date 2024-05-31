package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/yDog-1/woodon_server/ent"
)

type api struct {
	client *ent.Client
	ctx    context.Context
}

func (a *api) initEntClient() error {
	a.ctx = context.Background()
	// ent.Clientを作成
	client, err := ent.Open("mysql", "user:password@tcp(db:3306)/woodon?parseTime=True")
	if err != nil {
		return fmt.Errorf("failed opening connection to mysql: %v", err)
	}
	// スキーマの作成
	if err := client.Schema.Create(a.ctx); err != nil {
		return fmt.Errorf("failed creating schema resources: %v", err)
	}
	a.client = client
	return nil
}
func main() {
	var api api
	err := api.initEntClient()
	if err != nil {
		log.Fatalf("failed initEntClient(): %v", err)
	}
	defer api.client.Close()

	http.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	http.HandleFunc("POST /user", api.createUser)
	// HTTPサーバを8080ポートで起動
	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

func (a *api) createUser(w http.ResponseWriter, r *http.Request) {
	user, err := a.client.User.
		Create().
		SetName("test").
		SetEmail("test@example.com").
		Save(a.ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonUser, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonUser)
}
