package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/jinzhu/gorm"
)

//var employees []employee

type employee struct {
	ID      string `json:"id"`
	NAME    string `json:"name"`
	BALANCE int32  `json:"balance"`
}

var db *gorm.DB

func initDB() {
	var err error
	dataSourceNAME := "root:Ayogamaya@28@tcp(localhost:3306)/"
	db, err = gorm.Open("mysql", dataSourceNAME)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	// db.Exec("CREATE DATABASE employees_db")
	db.Exec("USE employees_db")

	db.AutoMigrate(&employee{})
}

func postEmployee(w http.ResponseWriter, r *http.Request) {
	var newEmployee employee
	json.NewDecoder(r.Body).Decode(&newEmployee)
	db.Create(&newEmployee)
	fmt.Println(newEmployee)
	w.Header().Set("Content-Type", "application/json")
	//employees = append(employees, newEmployee)
	json.NewEncoder(w).Encode(newEmployee)
}

func getEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	Employee_ID := params["id"]
	var employees []employee
	db.Find(&employees,Employee_ID).Where("id =", Employee_ID)
	json.NewEncoder(w).Encode(employees)

}
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/employees", postEmployee).Methods("POST")

	router.HandleFunc("/employees_get", getEmployee).Methods("GET")

	initDB()
	log.Fatal(http.ListenAndServe(":7000", router))
}
