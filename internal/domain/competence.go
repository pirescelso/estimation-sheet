package domain

import (
	"fmt"
	"time"

	"github.com/celsopires1999/estimation/internal/common"
	"github.com/google/uuid"
)

type Competence struct {
	CompetenceID string    `validate:"required,uuid4"`
	Code         string    `validate:"required,max=20"`
	Name         string    `validate:"required,max=50"`
	CreatedAt    time.Time `validate:"-"`
	UpdatedAt    time.Time `validate:"-"`
}

type RestoreCompetenceProps Competence

func NewCompetence(code string, name string) *Competence {
	return &Competence{
		CompetenceID: uuid.NewString(),
		Code:         code,
		Name:         name,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

func RestoreCompetence(props RestoreCompetenceProps) *Competence {
	return &Competence{
		CompetenceID: props.CompetenceID,
		Code:         props.Code,
		Name:         props.Name,
		CreatedAt:    props.CreatedAt,
		UpdatedAt:    props.UpdatedAt,
	}
}

func (c *Competence) ChangeCode(code *string) {
	if code == nil {
		return
	}
	c.Code = *code
}

func (c *Competence) ChangeName(name *string) {
	if name == nil {
		return
	}
	c.Name = *name
}

func (c *Competence) Validate() error {
	err := common.Validate.Struct(c)
	if err != nil {
		return common.NewDomainValidationError(fmt.Errorf("competence domain validation failed: %w", err))
	}
	return nil
}
