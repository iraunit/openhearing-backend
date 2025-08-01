package onboardingService

type OnboardingUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Gender   string `json:"gender"`
	PhoneNo  string `json:"phone_no"`
}
