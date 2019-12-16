package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ivonofian/shipper/conf"
	"github.com/ivonofian/shipper/models"
)

//GetProducts GetProducts
func GetProducts(w http.ResponseWriter, r *http.Request) {

	var products []models.Product

	db := conf.GetDB()

	result, err := db.Query("select * from product")
	defer result.Close()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for result.Next() {
		p := models.Product{}
		err := result.Scan(&p.ID, &p.Name, &p.Sku, &p.Qty, &p.Created, &p.Updated)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		products = append(products, p)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)

}
