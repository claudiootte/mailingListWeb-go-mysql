package model

import "fmt"

// As credenciais do banco de dados

var User = "root"
var Password = "admin"
var TCP = "@tcp(localhost:3306)/"
var DBName = "godb01"

var DNS = fmt.Sprintln(User + ":" + Password + TCP + DBName + "?")
var Driver = "mysql"
