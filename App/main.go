package main

import (
	"database/sql"
	//"fmt"
	"log"
	"net/http"
	"encoding/json"
    _"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"os"
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

//sets up our router and paths to listen for with actions
func handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/api/products", allProducts).Methods("GET")
	// myRouter.HandleFunc("/api/product/cpu", getCpu).Methods("GET")
	// myRouter.HandleFunc("/api/product/gpu", getGpu).Methods("GET")
	// myRouter.HandleFunc("/api/product/ram", getRam).Methods("GET")
	// myRouter.HandleFunc("/api/product/ssd", getSsd).Methods("GET")
	// myRouter.HandleFunc("/api/product/motherboard", getMotherboard).Methods("GET")
	myRouter.HandleFunc("/api/product/{name}", searchProduct).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", myRouter))
}

//all the functions used by the router
func allProducts(w http.ResponseWriter, r *http.Request) {
	runSQL(w, "SELECT * FROM product")
}

// Enpoints ready for category if needed/ I were to make this public

// func getCpu(w http.ResponseWriter, r *http.Request) {
// 	runSQL(w, "SELECT * FROM product p WHERE p.product_category='CPU'")
// }

// func getGpu(w http.ResponseWriter, r *http.Request) {
// 	runSQL(w, "SELECT * FROM product p WHERE p.product_category='GPU'")
// }

// func getRam(w http.ResponseWriter, r *http.Request) {
// 	runSQL(w, "SELECT * FROM product p WHERE p.product_category='Ram'")
// }

// func getSsd(w http.ResponseWriter, r *http.Request) {
// 	runSQL(w, "SELECT * FROM product p WHERE p.product_category='SSD'")
// }

// func getMotherboard(w http.ResponseWriter, r *http.Request) {
// 	runSQL(w, "SELECT * FROM product p WHERE p.product_category='Motherboard'")
//}

func searchProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["name"]
	query := "SELECT * FROM product p WHERE p.product_name like '%" + key + "%' or p.product_description like '%" + key + "%' or p.product_category like '%" + key + "%' "
	runSQL(w, query)
}

// Cross-Oiring Resource Sharing, allows my FE to access the API
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

//function that send a query over to database and exports a JSON from it
func runSQL(w http.ResponseWriter, x string) {
	var allProducts []Product

	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")

	//gets the password from my .env file
	err := godotenv.Load(".env")
	if err != nil {log.Fatal(err)}
	dbpassword := os.Getenv("DB_PASSWORD")
	dbpath := "root:" + dbpassword + "@tcp(sqldb:3306)/products"

	//open connection to sql and run query
	db, err := sql.Open("mysql", dbpath)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	result, err := db.Query(x)
	if err != nil {
		panic(err.Error())
	}

	//converts the responce of the query to my struct and then encodes a JSON 
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
