package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/driver/postgres"
  	"gorm.io/gorm"
	"net/http"
)

var db *gorm.DB

// gorm uses ID as pk, and pluralize struct name to snake_cases as table name.
type User struct{
	gorm.Model
	id uint 
	name string
	age uint8
}

// var users = []user{}

func main() {

	dsn := "host=localhost user=postgres password=eieiza dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	fmt.Println(db)

	if err != nil {
		panic("failed to connect database")
	}
	
	db.Debug().AutoMigrate(&User{})

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	//nested Route
	r.Route("/users", func(r chi.Router){
		r.Post("/", register)
		r.Route("/{id}", func(r chi.Router){
			r.Get("/searchResult", getById)
			r.Put("/editDetails", editById)
		})
		r.Get("/{searchOptions}-{searchParams}", searchByOptions)
	})

	http.ListenAndServe(":3000", r)
	fmt.Println("#############################")
}


func register(w http.ResponseWriter, r *http.Request) {
	user := User{
		name: "test",
		age: 20,
	}

	result := db.Create(&user)
	_ = result

}	

func getById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	_ = id
	//selectedUser = gorm stuffs

	// if selectedUser == nil {
    // w.WriteHeader(404)
    // w.Write([]byte("article not found"))
    // return
}

func editById(w http.ResponseWriter, r *http.Request) {

}

func searchByOptions(w http.ResponseWriter, r *http.Request) {

}
