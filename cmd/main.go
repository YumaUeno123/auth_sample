package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func signup(w http.ResponseWriter, req *http.Request) {
	fmt.Println("signup 関数実行")
}

func login(w http.ResponseWriter, req *http.Request) {
	fmt.Println("login 関数実行")
}

func main() {
	fmt.Println("Hello 世界")

	router := mux.NewRouter()
	router.HandleFunc("/signup", signup).Methods("POST")
	router.HandleFunc("/login", login).Methods("POST")

	log.Println("サーバー起動 : 8081 port で受信")

	log.Fatal(http.ListenAndServe(":8081", router))
}
