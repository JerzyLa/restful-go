package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	models "jsonstore/helper"
	"log"
	"net/http"
	"time"
)

type DBClient struct {
	db *gorm.DB
}

// PostPackage saves the package information
func (driver *DBClient) PostPackage(w http.ResponseWriter, r *http.Request) {
	var Package = models.Package{}
	postBody, _ := ioutil.ReadAll(r.Body)
	Package.Data = string(postBody)
	driver.db.Save(&Package)
	responseMap := map[string]interface{}{"id": Package.ID}
	w.Header().Set("Content-Type", "application/json")
	response, _ := json.Marshal(responseMap)
	w.Write(response)
}

type PackageResponse struct {
	Package models.Package `json:"Package"`
}

// GetPackage fetches the original URL for the given
// encoded(short) string
func (driver *DBClient) GetPackage(w http.ResponseWriter, r *http.Request) {
	var Package = models.Package{}
	vars := mux.Vars(r)

	driver.db.First(&Package, vars["id"])
	var PackageData interface{}

	json.Unmarshal([]byte(Package.Data), &PackageData)
	var response = PackageResponse{Package: Package}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	respJSON, _ := json.Marshal(response)
	w.Write(respJSON)
}

// GetPackagesbyWeight fetches all packages with given weight
func (driver *DBClient) GetPackagesbyWeight(w http.ResponseWriter, r *http.Request) {
	var packages []models.Package
	weight := r.FormValue("weight")
	// Handle response details
	var query = "select * from \"Package\" where data->>'weight'=?"
	driver.db.Raw(query, weight).Scan(&packages)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	respJSON, _ := json.Marshal(packages)
	w.Write(respJSON)
}

func main() {
	db, err := models.InitDB()
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	dbclient := &DBClient{db: db}

	r := mux.NewRouter()
	r.HandleFunc("/v1/package/{id:[a-zA-Z0-9]*}",
		dbclient.GetPackage).Methods("GET")
	r.HandleFunc("/v1/package",
		dbclient.PostPackage).Methods("POST")
	r.HandleFunc("/v1/package",
		dbclient.GetPackagesbyWeight).Methods("GET")
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}