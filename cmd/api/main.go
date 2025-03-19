package main

import (
	"log"
	"net/http"

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

	repo := repository.NewProductRepository(dbConnection)
	service := service.NewProductService(repo)
	handler := handler.NewProductRepository(service)

	http.HandleFunc("/products", handler.GetAll)

	log.Println("Server started at :8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}