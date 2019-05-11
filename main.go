package main

import (
	"fmt"
	"net/http"
	"os"
	"umesh/app"
	"umesh/controllers"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	// POST
	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/item", controllers.CreateItem).Methods("POST")

	router.HandleFunc("/api/bill/new", controllers.CreateCustomerBill).Methods("POST")
	router.HandleFunc("/api/bill/item/new", controllers.RegisterItemToBill).Methods("POST")

	// GET
	router.HandleFunc("/api/item", controllers.FetchItem).Queries("src", "", "item", "").Methods("GET")
	router.HandleFunc("/api/bill/item", controllers.GetBillItems).Queries("bid", "").Methods("GET")
	router.HandleFunc("/api/bill/all", controllers.GetAllBillsByCustomer).Queries("cid", "").Methods("GET")
	router.HandleFunc("/api/bill", controllers.GetBillByTitle).Queries("cid", "", "title", "").Methods("GET")

	router.Use(app.JwtAuthentication) //attach JWT auth middleware

	port := os.Getenv("PORT") //Get port from .env file, we did not specify any port so this should return an empty string when tested locally
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
