package onboardingService

import (
	"fmt"
	"github.com/iraunit/open-hearing/respository/onboardingRepo"
)

type OnboardingService interface {
	Test()
	OnboardNewUser(user *OnboardingUser) error
}

type OnboardingServiceImpl struct {
	onboardingRepo onboardingRepo.OnboardingRepo
}

func NewOnboardingRestHandler(onboardingRepo onboardingRepo.OnboardingRepo) *OnboardingServiceImpl {
	return &OnboardingServiceImpl{
		onboardingRepo: onboardingRepo,
	}
}

func (impl *OnboardingServiceImpl) Test() {
	fmt.Printf("OnboardingServiceImpl Test")
}

func (impl *OnboardingServiceImpl) OnboardNewUser(user *OnboardingUser) error {
	userDto := onboardingRepo.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		PhoneNo:  user.PhoneNo,
	}
	err := impl.onboardingRepo.OnboardNewUser(userDto)
	if err != nil {
		return err
	}
	return nil
}
