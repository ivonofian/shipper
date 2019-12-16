package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/ivonofian/shipper/conf"
	"github.com/ivonofian/shipper/models"
)

//UpdateProduct UpdateProduct
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	db := conf.GetDB()
	if ProductExists(db, params["id"]) == false {
		http.Error(w, "Product Not Found", http.StatusNotFound)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// panic(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	p := models.Product{}
	json.Unmarshal(body, &p)

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("update product set name=?,sku=?,qty=?,updated=? where id=?")
	_, err = stmt.Exec(p.Name, p.Sku, p.Qty, time.Now(), params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	tx.Commit()

	//get
	result, err := db.Query("SELECT * FROM product WHERE id = ?", params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer result.Close()
	pr := models.Product{}
	for result.Next() {
		err := result.Scan(&pr.ID, &pr.Name, &pr.Sku, &pr.Qty, &pr.Created, &pr.Updated)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
	json.NewEncoder(w).Encode(pr)
}
