package main

import (
	"github.com/claudiootte/mailingListWeb-go-mysql/controller"
	"github.com/claudiootte/mailingListWeb-go-mysql/model"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	model.ConnectDB()
	controller.InitializeRouter()

}
