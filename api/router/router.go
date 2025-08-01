package router

import (
	"github.com/gorilla/mux"
	"github.com/iraunit/open-hearing/api/restHandler"
)

type MuxRouter struct {
	Router       *mux.Router
	onboardingRH restHandler.OnboardingRestHandler
}

func NewMuxRouter(onboardingRestHandler restHandler.OnboardingRestHandler) *MuxRouter {
	return &MuxRouter{
		Router:       mux.NewRouter(),
		onboardingRH: onboardingRestHandler,
	}
}

func (r *MuxRouter) GetRouter() *mux.Router {
	r.Router.HandleFunc("/", r.onboardingRH.Test).Methods("GET")
	r.Router.HandleFunc("/", r.onboardingRH.OnboardNewUser).Methods("POST")
	return r.Router
}
