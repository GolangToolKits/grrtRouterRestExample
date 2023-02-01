package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	ph "github.com/GolangToolKits/grrtRouterRestExample/handlers"
	// "github.com/gorilla/mux"
	mux "github.com/GolangToolKits/grrt"
)

func main() {

	var sh ph.StoreHandler

	h := sh.New()

	router := mux.NewRouter()
	router.EnableCORS()
	router.CORSAllowCredentials()
	router.SetCorsAllowedHeaders("X-Requested-With, Content-Type, api-key, customer-key, Origin")
	router.SetCorsAllowedOrigins("*")
	router.SetCorsAllowedMethods("GET, DELETE, POST, PUT")

	port := "3000"
	envPort := os.Getenv("PORT")
	if envPort != "" {
		portInt, _ := strconv.Atoi(envPort)
		if portInt != 0 {
			port = envPort
		}
	}

	router.HandleFunc("/rs/product/get/{id}", h.GetProduct).Methods("GET")
	router.HandleFunc("/rs/product/get/{id}/{sku}", h.GetProductWithIDAndSku).Methods("GET")
	router.HandleFunc("/rs/products", h.GetProducts).Methods("GET")
	router.HandleFunc("/rs/product/add", h.AddProduct).Methods("POST")
	router.HandleFunc("/rs/product/update", h.UpdateProduct).Methods("PUT")
	fmt.Println("running on Port:", port)
	http.ListenAndServe(":"+port, (router))

}

//go mod init github.com/GolangToolKits/grrtRouterRestExample
