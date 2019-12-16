package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/ivonofian/shipper/conf"
)

//DeleteProduct DeleteProduct
func DeleteProduct(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	db := conf.GetDB()

	if ProductExists(db, params["id"]) == false {
		http.Error(w, "Product Not Found", http.StatusNotFound)
		return
	}

	stmt, err := db.Prepare("DELETE FROM product WHERE id = ?")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = stmt.Exec(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Post with ID = %s was deleted", params["id"])
}
