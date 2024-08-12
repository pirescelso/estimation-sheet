package domain

import (
	"fmt"
	"time"

	"github.com/celsopires1999/estimation/internal/common"
	"github.com/google/uuid"
)

type UserType string

func (ut UserType) String() string {
	return string(ut)
}

const (
	Manager   UserType = "manager"
	Estimator UserType = "estimator"
)

type User struct {
	UserID    string    `validate:"required,uuid4" `
	Email     string    `validate:"required,email"`
	UserName  string    `validate:"required"`
	Name      string    `validate:"required"`
	UserType  UserType  `validate:"required,oneof=manager estimator"`
	CreatedAt time.Time `validate:"-"`
	UpdatedAt time.Time `validate:"-"`
}

type RestoreUserProps User

func NewUser(email string, username string, name string, userType UserType) *User {
	return &User{
		UserID:   uuid.NewString(),
		Email:    email,
		UserName: username,
		Name:     name,
		UserType: userType,
	}
}

func RestoreUser(props RestoreUserProps) *User {
	return &User{
		UserID:    props.UserID,
		Email:     props.Email,
		UserName:  props.UserName,
		Name:      props.Name,
		UserType:  props.UserType,
		CreatedAt: props.CreatedAt,
		UpdatedAt: props.UpdatedAt,
	}
}

func (u *User) ChangeEmail(email *string) {
	if email == nil {
		return
	}
	u.Email = *email
}

func (u *User) ChangeUserName(userName *string) {
	if userName == nil {
		return
	}
	u.UserName = *userName
}

func (u *User) ChangeName(name *string) {
	if name == nil {
		return
	}
	u.Name = *name
}

func (u *User) ChangeUserType(userTypeStr *string) {
	if userTypeStr == nil {
		return
	}
	u.UserType = UserType(*userTypeStr)
}

func (u *User) Validate() error {
	err := common.Validate.Struct(u)
	if err != nil {
		return common.NewDomainValidationError(fmt.Errorf("user domain validation failed: %w", err))
	}
	return nil
}
