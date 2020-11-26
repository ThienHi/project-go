package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Customer struct {
	No          int       `json:"no" gorm:"primary key; auto_increment;"`
	Name        string    `json:"name" gorm:"size 50;"`
	License     string    `json:"license" gorm:"size 255;"`
	StartDate   time.Time `json:"start_date" gorm:"CURRENT_TIMESTAMP;"`
	ExpriteDate time.Time `json:"exprite_date" gorm:"CURRENT_TIMESTAMP;"`
	Status      bool      `json:"status"`
}

type CreateCustomer struct {
	Name    string `json:"name" gorm:"size 50"`
	License string `json:"license" gorm:"size 50"`
	Status  bool   `json:"status"`
}

var Customers []Customer
var prevCusID = 0

func createOrder(w http.ResponseWriter, r *http.Request) {
	var cus Customer
	json.NewDecoder(r.Body).Decode(&cus)
	prevCusID++
	// cus.No = strconv.Itoa(No)
	Customers = append(Customers, cus)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cus)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/get", getCustomers)
	log.Fatal(http.ListenAndServe(":2603", nil))
}

func returnAllCustomers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Customers")
	json.NewEncoder(w).Encode(Customers)
}

func getCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Customers)
}

func main() {
	Customers = []Customer{
		Customer{No: 1, Name: "ThienHi", License: "Article Content", StartDate: time.Now().Local().UTC(), ExpriteDate: time.Now().Local().UTC(), Status: true},
		Customer{No: 1, Name: "Kim", License: "Article Content", Status: false},
	}
	http.HandleFunc("/customer", returnAllCustomers)

	handleRequests()
}
