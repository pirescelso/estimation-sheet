package testutils

import (
	"time"

	"github.com/Pallinder/go-randomdata"
	"github.com/celsopires1999/estimation/internal/domain"
	"github.com/google/uuid"
)

type UserFakeBuilder struct {
	UserID    string
	Email     string
	UserName  string
	Name      string
	UserType  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUserFakeBuilder() *UserFakeBuilder {
	return &UserFakeBuilder{
		UserID:    uuid.NewString(),
		Email:     randomdata.Email(),
		UserName:  randomdata.Letters(8),
		Name:      randomdata.FullName(randomdata.RandomGender),
		UserType:  randomdata.StringSample("manager", "estimator"),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (b *UserFakeBuilder) WithName(name string) *UserFakeBuilder {
	b.Name = name
	return b
}

func (b *UserFakeBuilder) WithManager() *UserFakeBuilder {
	b.UserType = "manager"
	return b
}

func (b *UserFakeBuilder) WithEstimator() *UserFakeBuilder {
	b.UserType = "estimator"
	return b
}

func (b *UserFakeBuilder) Build() *domain.User {
	return &domain.User{
		UserID:    b.UserID,
		Email:     b.Email,
		UserName:  b.UserName,
		Name:      b.Name,
		UserType:  domain.UserType(b.UserType),
		CreatedAt: b.CreatedAt,
		UpdatedAt: b.UpdatedAt,
	}
}
