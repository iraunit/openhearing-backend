//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/iraunit/open-hearing/api/restHandler"
	"github.com/iraunit/open-hearing/api/router"
	"github.com/iraunit/open-hearing/respository/db"
	"github.com/iraunit/open-hearing/respository/onboardingRepo"
	"github.com/iraunit/open-hearing/service/onboardingService"
)

func InitializeApp() *App {
	wire.Build(
		NewApp,
		db.NewPgDb,
		router.NewMuxRouter,
		restHandler.NewOnboardingRestHandler, wire.Bind(new(restHandler.OnboardingRestHandler), new(*restHandler.OnboardingRestHandlerImpl)),
		onboardingService.NewOnboardingRestHandler, wire.Bind(new(onboardingService.OnboardingService), new(*onboardingService.OnboardingServiceImpl)),
		onboardingRepo.NewOnboardingRepoImpl, wire.Bind(new(onboardingRepo.OnboardingRepo), new(*onboardingRepo.OnboardingRepoImpl)),
	)
	return &App{}
}
