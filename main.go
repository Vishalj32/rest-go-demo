package main

import (
	"log"
	"net/http"
	"rest-go-demo/controller"
	"rest-go-demo/database/mysql"
	"rest-go-demo/entity"

	"github.com/gorilla/mux"

	_ "github.com/jinzhu/gorm/dialects/mysql" //Required for MySQL dialect
)

func main() {
	dbInit()
	log.Println("Starting development server at http://127.0.0.1:3000/")
	log.Println("Quit the server with CONTROL-C.")

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/create", controller.CreatePerson).Methods("POST")
	myRouter.HandleFunc("/get/{id}", controller.GetPerson).Methods("GET")
	myRouter.HandleFunc("/get", controller.GetAllPersons).Methods("GET")
	myRouter.HandleFunc("/delete/{id}", controller.DeletePerson).Methods("DELETE")
	myRouter.HandleFunc("/update", controller.UpdatePerson).Methods("PUT")

	log.Fatal(http.ListenAndServe(":3000", myRouter))

	defer mysql.Connector.Close()
}

func dbInit() {
	dbConf := mysql.Config{
		ServerName: "localhost:3306",
		User:       "root",
		Password:   "root",
		DB:         "learning",
	}

	connectionString := mysql.GetConnectionString(dbConf)
	connector, err := mysql.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}

	mysql.MigrateDB(connector, &entity.Person{})
}
