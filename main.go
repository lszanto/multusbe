package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/lszanto/multusbe/handlers"
	"github.com/lszanto/multusbe/multus"
	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
)

func main() {
	config := multus.LoadConfig("config.json")

	db, err := gorm.Open("mysql", config.DBString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := chi.NewRouter()

	// apply middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	docHandler := handlers.NewDocHandler(db)

	r.Route("/api", func(r chi.Router) {
		r.Route("/doc", func(r chi.Router) {
			r.Post("/", docHandler.Create)
			r.Get("/:id", docHandler.GetByID)
			r.Get("/exists/:title", docHandler.Exists)
		})
	})

	http.ListenAndServe(":3000", r)
}
