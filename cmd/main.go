package main

import (
	"awesomeproject/configs"
	"awesomeproject/internal/auth"
	"awesomeproject/pkg/db"
	"fmt"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	_ = db.NewDb(conf)
	router := http.NewServeMux()
	auth.SetupHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})
	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}
	//  добавить логер
	fmt.Println("Server is listening on port 8081")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
