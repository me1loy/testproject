package main

import (
	"context"
	"log"
	"net/http"
	"os"
	cmd1 "server/database/init"
	rou "server/routers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	ctx := context.Background()

	cfg := cmd1.GetConfig()

	db, err := cmd1.NewClient(ctx, cfg.Storage)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	repo1 := cmd1.NewRepository(db)

	handler := &rou.Handler{
		Repo3: repo1,
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://localhost:8080"},
		AllowedMethods: []string{"GET", "POST"},
		AllowedHeaders: []string{"Content-Type"},
	}))

	r.Post("/support", handler.CreateUser)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		rou.RenderSupportPage(w, rou.SupportPageData{})
	})

	fs := http.FileServer(http.Dir("static"))
	r.Handle("/*", fs)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	addr := ":" + port
	log.Printf("Starting server on %s", addr)
	err = http.ListenAndServe(addr, r)

	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
