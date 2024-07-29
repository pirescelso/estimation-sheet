// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	domain "github.com/celsopires1999/estimation/internal/domain"
	"github.com/jackc/pgx/v5/pgtype"
)

type Baseline struct {
	BaselineID  string
	Code        string
	Review      int32
	Title       string
	Description pgtype.Text
	StartDate   pgtype.Date
	Duration    int32
	ManagerID   string
	EstimatorID string
	CreatedAt   pgtype.Timestamp
	UpdatedAt   pgtype.Timestamp
}

type Budget struct {
	BudgetID    string
	PortfolioID string
	CostID      string
	Amount      float64
	CreatedAt   pgtype.Timestamp
	UpdatedAt   pgtype.Timestamp
}

type BudgetAllocation struct {
	BudgetAllocationID string
	BudgetID           string
	AllocationDate     pgtype.Date
	Amount             float64
	CreatedAt          pgtype.Timestamp
	UpdatedAt          pgtype.Timestamp
}

type Cost struct {
	CostID      string
	BaselineID  string
	CostType    string
	Description string
	Comment     pgtype.Text
	Amount      float64
	Currency    string
	Tax         float64
	CreatedAt   pgtype.Timestamp
	UpdatedAt   pgtype.Timestamp
}

type CostAllocation struct {
	CostAllocationID string
	CostID           string
	AllocationDate   pgtype.Date
	Amount           float64
	CreatedAt        pgtype.Timestamp
	UpdatedAt        pgtype.Timestamp
}

type Plan struct {
	PlanID      string
	Code        string
	Name        string
	Assumptions domain.Assumptions
	CreatedAt   pgtype.Timestamp
	UpdatedAt   pgtype.Timestamp
}

type Portfolio struct {
	PortfolioID string
	BaselineID  string
	PlanID      string
	StartDate   pgtype.Date
	CreatedAt   pgtype.Timestamp
	UpdatedAt   pgtype.Timestamp
}

type User struct {
	UserID    string
	Email     string
	UserName  string
	Name      string
	UserType  string
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}
