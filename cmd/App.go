package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/iraunit/open-hearing/api/router"
	"net/http"
)

type App struct {
	MuxRouter *router.MuxRouter
}

func NewApp(mux *router.MuxRouter) *App {
	return &App{
		MuxRouter: mux,
	}
}

func (app *App) Start() {

	app.MuxRouter.GetRouter()
	corsOrigins := handlers.AllowedOrigins([]string{
		"*",
	})

	fmt.Printf("Listening on port %s", "8000")
	if err := http.ListenAndServe(fmt.Sprintf(":%s", "8000"), handlers.CORS(corsOrigins)(app.MuxRouter.Router)); err != nil {
		panic(err)
	}
}
