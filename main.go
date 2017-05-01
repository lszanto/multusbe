package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/goware/cors"
	"github.com/jinzhu/gorm"
	"github.com/lszanto/multusbe/handlers"
	"github.com/lszanto/multusbe/multus"
	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
)

func main() {
	config := multus.LoadConfig("config.json")
	db, err := gorm.Open(config.DBEngine, config.DBString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// db.AutoMigrate(&models.User{})
	// db.AutoMigrate(&models.Doc{})

	// // Create a new token object, specifying signing method and the claims
	// // you would like it to contain.
	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	// 	"foo": "bar",
	// 	"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	// })

	// // Sign and get the complete encoded token as a string using the secret
	// tokenString, err := token.SignedString(config.SecretKey)

	// // fmt.Println(tokenString, err)
	// fmt.Println("create model")
	// user := models.User{}
	// user.Username = "luke"
	// user.Email = "luke.found@gmail.com"
	// user.Password = "cools"
	// hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	// user.Password = string(hash)
	// fmt.Println("creating user")
	// db.Create(&user)
	// fmt.Println("Created")

	r := chi.NewRouter()

	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	cors := cors.New(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)

	// apply middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	docHandler := handlers.NewDocHandler(db)
	userHandler := handlers.NewUserHandler(db)

	r.Route("/api", func(r chi.Router) {
		r.Route("/doc", func(r chi.Router) {
			r.Post("/", docHandler.Create)
			r.Get("/:title", docHandler.Get)
		})

		r.Route("/user", func(r chi.Router) {
			r.Post("/login", userHandler.Login)
			r.Post("/create", userHandler.Create)
		})
	})

	fmt.Println("Server to listen on port :3000")
	http.ListenAndServe(":3000", r)
}
