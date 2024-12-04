package main

import (
	"github.com/Rohanrevanth/muzi-go/database"
	"github.com/Rohanrevanth/muzi-go/http"
)

func main() {
	database.ConnectDatabase()
	// database.InitializeRedis()
	http.StartServer()
}
