package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"go-contacts/app"
	"net/http"
	"os"
)

func main() {

	router := mux.NewRouter()
	router.Use(app.JwtAuthentication) // добавляем middleware проверки JWT-токена

	port := os.Getenv("PORT") //Получить порт из файла .env; мы не указали порт, поэтому при локальном тестировании должна возвращаться пустая строка
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router) //Запустите приложение, посетите localhost:8000/api

	if err != nil {
		fmt.Print(err)
	}

	router.HandleFunc("/api/user/new",
		controllers.CreateAccount).Methods("POST")

	router.HandleFunc("/api/user/login",
		controllers.Authenticate).Methods("POST")

	router.HandleFunc("/api/me/contacts",
		controllers.GetContactsFor).Methods("GET")
}
