package main

import (
	"database/sql"
	//"fmt"
	"log"
	"net/http"
	"encoding/json"
    _"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Product struct {
	ID 			int `json:"ID"`
	Name        string `json:"name"`
	Category    string `json:"category"`
	Image       string `json:"image"`
	Description string `json:"description"`
	Price       float32 `json:"price"`
}


func main() {
	handleRequests()
}

func handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	//myRouter.HandleFunc("/", home)
	myRouter.HandleFunc("/api/products", allProducts).Methods("GET")
	myRouter.HandleFunc("/api/product/cpu", getCpu).Methods("GET")
	myRouter.HandleFunc("/api/product/gpu", getGpu).Methods("GET")
	myRouter.HandleFunc("/api/product/ram", getRam).Methods("GET")
	myRouter.HandleFunc("/api/product/ssd", getSsd).Methods("GET")
	myRouter.HandleFunc("/api/product/motherboard", getMotherboard).Methods("GET")
	log.Fatal(http.ListenAndServe(":8088", myRouter))
}

func allProducts(w http.ResponseWriter, r *http.Request) {
	runSQL(w, "SELECT * FROM product")
}

func getCpu(w http.ResponseWriter, r *http.Request) {
	runSQL(w, "SELECT * FROM product p WHERE p.product_category='CPU'")
}

func getGpu(w http.ResponseWriter, r *http.Request) {
	runSQL(w, "SELECT * FROM product p WHERE p.product_category='GPU'")
}

func getRam(w http.ResponseWriter, r *http.Request) {
	runSQL(w, "SELECT * FROM product p WHERE p.product_category='Ram'")
}

func getSsd(w http.ResponseWriter, r *http.Request) {
	runSQL(w, "SELECT * FROM product p WHERE p.product_category='SSD'")
}

func getMotherboard(w http.ResponseWriter, r *http.Request) {
	runSQL(w, "SELECT * FROM product p WHERE p.product_category='Motherboard'")
}


func runSQL(w http.ResponseWriter, x string) {
	var allProducts []Product

	db, err := sql.Open("mysql", "root:econ21@tcp(127.0.0.1:3306)/products")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	result, err := db.Query(x)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var product Product
		err =result.Scan(&product.ID, &product.Name, &product.Category, &product.Image, &product.Description, &product.Price )
		if err != nil {
			panic(err.Error())
		}
		allProducts = append(allProducts, product)	
	}

	json.NewEncoder(w).Encode(allProducts)
}
