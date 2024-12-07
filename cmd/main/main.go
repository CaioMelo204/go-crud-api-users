package main

import (
	"FirstCRUD/pkg/routes"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/user", routes.CreateUser).Methods("POST")
	router.HandleFunc("/user", routes.GetAllUsers).Methods("GET")
	router.HandleFunc("/user/{id}", routes.GetUser).Methods("GET")
	router.HandleFunc("/user/{id}", routes.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{id}", routes.DeleteUser).Methods("DELETE")
	fmt.Printf("Listening on port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", router))
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
