package main // import "github.com/ivonofian/shipper/restapi"

import (
	"net/http"

	ghand "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/ivonofian/shipper/conf"
	"github.com/ivonofian/shipper/handlers"
)

func main() {
	conf.ConfigureDB()

	router := mux.NewRouter()
	router.HandleFunc("/products", handlers.GetProducts).Methods("GET")
	router.HandleFunc("/products", handlers.CreateProduct).Methods("POST")
	router.HandleFunc("/products/{id}", handlers.GetProduct).Methods("GET")
	router.HandleFunc("/products/{id}", handlers.UpdateProduct).Methods("PUT")
	router.HandleFunc("/products/{id}", handlers.DeleteProduct).Methods("DELETE")

	// Where ORIGIN_ALLOWED is like `scheme://dns[:port]`, or `*` (insecure)
	headersOk := ghand.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOk := ghand.AllowedOrigins([]string{"*"})
	methodsOk := ghand.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})

	http.ListenAndServe(":8001", ghand.CORS(originsOk, headersOk, methodsOk)(router))
}
