package main

import (
	"awesomeproject/configs"
	"awesomeproject/internal/auth"
	"awesomeproject/internal/link"
	"awesomeproject/pkg/db"
	"fmt"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()

	//Repositories
	linkRepository := link.NewLinkRepository(db)

	//Handler
	auth.SetupHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})
	link.SetupLinkHandler(router, link.LinkHandlderDeps{LinkRepository: linkRepository})

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
