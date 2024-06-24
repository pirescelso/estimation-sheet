package testutils

import (
	"github.com/Pallinder/go-randomdata"
	"github.com/celsopires1999/estimation/internal/service"
)

type UserFakeBuilder struct {
	Email    string
	UserName string
	Name     string
	UserType string
}

func NewUserFakeBuilder() *UserFakeBuilder {
	return &UserFakeBuilder{
		Email:    randomdata.Email(),
		UserName: randomdata.Letters(8),
		Name:     randomdata.FullName(randomdata.RandomGender),
		UserType: randomdata.StringSample("manager", "estimator"),
	}
}

func (b *UserFakeBuilder) WithManager() *UserFakeBuilder {
	b.UserType = "manager"
	return b
}

func (b *UserFakeBuilder) WithEstimator() *UserFakeBuilder {
	b.UserType = "estimator"
	return b
}

func (b *UserFakeBuilder) Build() service.CreateUserInputDTO {
	return service.CreateUserInputDTO{
		Email:    b.Email,
		UserName: b.UserName,
		Name:     b.Name,
		UserType: b.UserType,
	}
}
