package onboardingRepo

import (
	"fmt"
	"github.com/go-pg/pg"
)

type User struct {
	tableName struct{} `sql:"users"`
	Name      string   `pg:"name"`
	Email     string   `pg:"email"`
	Password  string   `pg:"password"`
	PhoneNo   string   `pg:"phone_no"`
}

type OnboardingRepo interface {
	OnboardNewUser(user User) error
}

type OnboardingRepoImpl struct {
	db *pg.DB
}

func NewOnboardingRepoImpl(db *pg.DB) *OnboardingRepoImpl {
	return &OnboardingRepoImpl{
		db: db,
	}
}

func (impl *OnboardingRepoImpl) OnboardNewUser(user User) error {
	err := impl.db.Insert(&user)
	if err != nil {
		fmt.Printf(err.Error())
		return err
	}
	return nil
}
