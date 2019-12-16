package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ivonofian/shipper/conf"
	"github.com/ivonofian/shipper/models"
)

//GetProduct GetProduct
func GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	db := conf.GetDB()

	if ProductExists(db, params["id"]) == false {
		http.Error(w, "Product Not Found", http.StatusNotFound)
		return
	}

	result, err := db.Query("SELECT * FROM product WHERE id = ?", params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer result.Close()
	var p models.Product
	for result.Next() {
		err := result.Scan(&p.ID, &p.Name, &p.Sku, &p.Qty, &p.Created, &p.Updated)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
	json.NewEncoder(w).Encode(p)
}

//ProductExists check Product
func ProductExists(db *sql.DB, id string) bool {
	sqlStmt := `SELECT id FROM product WHERE id = ?`
	err := db.QueryRow(sqlStmt, id).Scan(&id)
	if err != nil {
		if err != sql.ErrNoRows {
			// a real error happened! you should change your function return
			return false
		}

		return false
	}

	return true
}
