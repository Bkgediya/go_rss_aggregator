package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Bkgediya/go_rss_aggregator/internal/db"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.Connect()

	// db_url := os.Getenv("DATABASE_URL")

	// conn, err := pgx.Connect(context.Background(), db_url)
	// if err != nil {
	// 	log.Fatal("Unable to connect to database:", err)
	// }
	// defer conn.Close(context.Background())

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("PORT is not found in the environment")
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handleErr)

	router.Mount("/v1", v1Router)

	// router.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello World"))
	// })

	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Printf("Server started on port %s", port)

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
