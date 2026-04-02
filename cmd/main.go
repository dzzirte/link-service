package main

import (
	"awesomeproject/configs"
	"awesomeproject/internal/auth"
	"awesomeproject/internal/link"
	"awesomeproject/internal/stat"
	"awesomeproject/internal/user"
	"awesomeproject/pkg/db"
	"awesomeproject/pkg/event"
	"awesomeproject/pkg/middleware"
	"fmt"
	"net/http"
)

func App(envFile string) http.Handler {
	conf := configs.LoadConfig(envFile)
	db := db.NewDb(conf)
	router := http.NewServeMux()
	eventBus := event.NewEventBus()

	//Repositories
	linkRepository := link.NewLinkRepository(db)
	userRepository := user.NewUserRepository(db)
	statRepository := stat.NewStatRepository(db)

	//Services
	authService := auth.NewAuthService(userRepository)
	statService := stat.NewStatService(&stat.StatServiceDeps{
		EventBus:       eventBus,
		StatRepository: statRepository,
	})

	//Handlers
	auth.SetupAuthHandler(router, auth.AuthHandlerDeps{
		Config:      conf,
		AuthService: authService,
	})
	link.SetupLinkHandler(router, link.LinkHandlderDeps{
		LinkRepository: linkRepository,
		EventBus:       eventBus,
		Config:         conf,
	})
	stat.SetupStatHandler(router, stat.StatHandlerDeps{
		StatRepository: statRepository,
		Config:         conf,
	})
	go statService.AddClick()
	//Middlewares
	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)
	return stack(router)
}

func main() {
	app := App(".env")
	server := http.Server{
		Addr:    ":8081",
		Handler: app,
	}
	fmt.Println("Server is listening on port 8081")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
