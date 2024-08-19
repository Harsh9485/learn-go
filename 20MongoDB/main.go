package main

import (
	"GOMongoDB/controllers"
	"GOMongoDB/util"

	"github.com/gorilla/mux"
)

func main() {
    

    Client := util.ConnectDB()
	
    r := mux.NewRouter()
    r.HandleFunc("/watched", func ()  {
        controllers.AddMovie()
    }).Methods("GET")
}