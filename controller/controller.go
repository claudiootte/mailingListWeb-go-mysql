package controller

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/claudiootte/mailingListWeb-go-mysql/model"
)

// Abaixo configura-se o roteamendo dos endpoints disponíveis, sendo eles os principais (CRUD), e posteriormente inicializa-se na porta 8080

func InitializeRouter() {
	MyRouter := mux.NewRouter()
	MyRouter.HandleFunc("/", model.Home)
	MyRouter.HandleFunc("/create", model.Create)
	MyRouter.HandleFunc("/insert", model.Insert)
	MyRouter.HandleFunc("/edit", model.Edit)
	MyRouter.HandleFunc("/update", model.Update)
	MyRouter.HandleFunc("/delete", model.Delete)

	log.Println("Connecting to server")
	log.Fatal(http.ListenAndServe(":8080", MyRouter))

}
