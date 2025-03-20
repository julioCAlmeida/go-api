package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/julioCAlmeida/go-api/db"
	"github.com/julioCAlmeida/go-api/internal/handler"
	"github.com/julioCAlmeida/go-api/internal/repository"
	"github.com/julioCAlmeida/go-api/internal/service"
	_ "github.com/lib/pq"
)

func main() {	
	dbConnection, err := db.Connect()
	if err != nil {
		panic(err)
	}
	defer dbConnection.Close()

	r := chi.NewRouter()

	repo := repository.NewProductRepository(dbConnection)
	service := service.NewProductService(repo)
	handler := handler.NewProductRepository(service)

	r.Get("/products", handler.GetAll)
	r.Get("/product/{id}", handler.GetById)
	r.Post("/product", handler.Create)
	r.Put("/product/{id}", handler.Update)
	r.Delete("/product/{id}", handler.Delete)

	log.Println("Server started at :8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}