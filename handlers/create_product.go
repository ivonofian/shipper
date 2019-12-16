package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/ivonofian/shipper/conf"
	"github.com/ivonofian/shipper/models"
)

//CreateProduct CreateProduct
func CreateProduct(w http.ResponseWriter, r *http.Request) {

	p := models.Product{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// panic(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.Unmarshal(body, &p)

	db := conf.GetDB()

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into product (name,sku,qty,created,updated) values (?,?,?,?,?)")
	_, err = stmt.Exec(p.Name, p.Sku, p.Qty, time.Now(), time.Now())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	tx.Commit()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("200 - Product Inserted!"))
}
