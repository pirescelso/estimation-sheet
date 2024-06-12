package domain

import (
	"time"

	"github.com/google/uuid"
)

type Project struct {
	ProjectID string
	StarDate  time.Time
}

func NewProject(startDate time.Time) *Project {
	return &Project{
		ProjectID: uuid.New().String(),
		StarDate:  startDate,
	}
}

func (p *Project) AddMonths(months int) error {
	p.StarDate = p.StarDate.AddDate(0, months, 0)
	return nil
}
