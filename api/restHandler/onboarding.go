package restHandler

import (
	"encoding/json"
	"fmt"
	"github.com/iraunit/open-hearing/service/onboardingService"
	"net/http"
)

type OnboardingRestHandler interface {
	Test(http.ResponseWriter, *http.Request)
	OnboardNewUser(w http.ResponseWriter, r *http.Request)
}

type OnboardingRestHandlerImpl struct {
	onboardingSvc onboardingService.OnboardingService
}

func NewOnboardingRestHandler(onboardingService onboardingService.OnboardingService) *OnboardingRestHandlerImpl {
	return &OnboardingRestHandlerImpl{
		onboardingSvc: onboardingService,
	}
}

func (handler *OnboardingRestHandlerImpl) Test(w http.ResponseWriter, r *http.Request) {
	handler.onboardingSvc.Test()
	_ = json.NewEncoder(w).Encode(map[string]string{"hello": "world"})

	return
}

func (handler *OnboardingRestHandlerImpl) OnboardNewUser(w http.ResponseWriter, r *http.Request) {

	user := onboardingService.OnboardingUser{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]string{"err": "could not decode request"})
		return
	}
	err = handler.onboardingSvc.OnboardNewUser(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error in onboarding new user: ", err)
		_ = json.NewEncoder(w).Encode(map[string]string{"err": "Internal Server Error"})
		return
	}
	w.WriteHeader(http.StatusAccepted)
	_ = json.NewEncoder(w).Encode(user)
}
