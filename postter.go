package main

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	"log"
//	"fmt"
	"io/ioutil"
	"math/rand"
)

var db *sql.DB
var err error

type Cars struct {
	ID int
	Brand string 
	Model string 
	Horse int
}


func createCar(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare("INSERT INTO cars(id, brand, model, horse_power) VALUES(?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	var newcar Cars
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	r.Body.Close()

	json.Unmarshal(body, &newcar)

	newcar.ID = rand.Intn(1000000) //de vegades fa que hi hagi socker hang up en el postman

	_, err = stmt.Exec(newcar.ID, newcar.Brand, newcar.Model, newcar.Horse)
	if err != nil {
		panic(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newcar)
/*
//	fmt.Println(w, "New post was created")

	//GET ID Y IMPRIMIR
	w.Header().Set("Content-Type", "application/json")
	result, err := db.Query("SELECT * FROM cars")
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	var car Cars

	for result.Next() {
		err := result.Scan(&car.ID, &car.Brand, &car.Model, &car.Horse)
		if err != nil {
			panic(err.Error())
		}
	}

	json.NewEncoder(w).Encode(car)*/
}


func main() {
	db, err = sql.Open("mysql", "postter:12345@tcp(my-mysql:3306)/test")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	router := mux.NewRouter().StrictSlash(true)

//	router.HandleFunc("/cars", getCars).Methods("GET")
//	router.HandleFunc("/cars/{id}", getCar).Methods("GET")
	router.HandleFunc("/cars", createCar).Methods("POST")

	log.Fatal(http.ListenAndServe(":8082", router))

}