package main

import (
	"test-mekar-backend/config"

	"github.com/gorilla/mux"
)

func main() {
	conf := config.NewConfig()
	db := config.NewDatabase(conf)
	router := mux.NewRouter()

	config.NewRoutes(db, router)
	defer db.Close()
	config.StartRouter(router)
}
