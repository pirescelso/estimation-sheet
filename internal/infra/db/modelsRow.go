package db

import "github.com/jackc/pgx/v5/pgtype"

type BaselineRow struct {
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
	Manager     string
	Estimator   string
}

type BudgetRow struct {
	BudgetID     string
	PortfolioID  string
	CostType     string
	Description  string
	Comment      pgtype.Text
	CostAmount   float64
	CostCurrency string
	CostTax      float64
	Amount       float64
	CreatedAt    pgtype.Timestamp
	UpdatedAt    pgtype.Timestamp
}

type WorkloadRow struct {
	WorkloadID     string
	PortfolioID    string
	CompetenceCode string
	CompetenceName string
	Comment        pgtype.Text
	Hours          int32
	CreatedAt      pgtype.Timestamp
	UpdatedAt      pgtype.Timestamp
}

type PortfolioRow struct {
	PortfolioID string
	PlanCode    string
	Code        string
	Review      int32
	Title       string
	Description pgtype.Text
	StartDate   pgtype.Date
	Duration    int32
	Manager     string
	Estimator   string
	CreatedAt   pgtype.Timestamp
	UpdatedAt   pgtype.Timestamp
}
