package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"

	handlers "example/handlers"
	database "example/database"
)

func main(){
	r := mux.NewRouter()
	fmt.Println("New router")
	db, errDb:= database.SetupDb()
	if(errDb != nil) {
		fmt.Println("Can not connect to db")
	}
	
	userHandler := handlers.NewUserHandler(db)
	fmt.Println("Create handdler")
	r.HandleFunc("/", func (writer http.ResponseWriter, request *http.Request){
		fmt.Fprintf(writer, "Hello World! A hi hi update")
	}).Methods(http.MethodGet)
	r.HandleFunc("/api/people/{id}", userHandler.GetById).Methods(http.MethodGet)
	r.HandleFunc("/api/people/create", userHandler.Create).Methods(http.MethodPost)
	r.HandleFunc("/api/people", userHandler.GetAll).Methods(http.MethodGet)
	r.HandleFunc("/api/people", userHandler.Update).Methods(http.MethodPut)
	fmt.Println("Create endpoints")
	err := http.ListenAndServe(":8088", r)
	fmt.Println("Serve")
	if(err != nil) {
		fmt.Println("Error ", err)
		fmt.Println(err)
		log.Fatal(err)
	}
	fmt.Println("Run")
}

