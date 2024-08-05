package main_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/celsopires1999/estimation/internal/testutils"
	"github.com/golang-migrate/migrate/v4"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/suite"
)

type E2EScenarioSuite struct {
	suite.Suite
	dbpool    *pgxpool.Pool
	m         *migrate.Migrate
	manager   UserOutput
	estimator UserOutput
	baseline  BaselineOutput
	mu        sync.Mutex
}

func (s *E2EScenarioSuite) SetupSuite() {
	s.dbpool, s.m = testutils.DBSetup()
}

func (s *E2EScenarioSuite) TearDownSuite() {
	defer s.dbpool.Close()
	err := s.m.Down()
	s.Nil(err)
}

func TestE2EScenario(t *testing.T) {
	suite.Run(t, new(E2EScenarioSuite))
}

func (s *E2EScenarioSuite) TestE2ECreate() {
	s.postUsersManager()
	s.postUsersEstimator()
	s.postBaselines()
}

func (s *E2EScenarioSuite) postUsersManager() {
	c := http.Client{}

	input := UserInput{
		Email:    "john.doe@userland.com",
		UserName: "john1234",
		Name:     "John Doe",
		UserType: "manager",
	}

	message, err := json.Marshal(input)
	s.Nil(err)

	payload := bytes.NewReader(message)

	r, err := c.Post("http://localhost:9000/api/v1/users", "application/json", payload)
	s.Nil(err)
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	s.Nil(err)

	s.Equal(http.StatusCreated, r.StatusCode)

	var output UserOutput
	err = json.Unmarshal(body, &output)
	s.Nil(err)

	_, err = uuid.Parse(output.UserID)
	s.Nil(err, "UserID should be a valid uuidv4")
	s.Equal(input.Email, output.Email)
	s.Equal(input.UserName, output.UserName)
	s.Equal(input.Name, output.Name)
	s.Equal(input.UserType, output.UserType)
	parsedTime, err := time.Parse(time.RFC3339, output.CreatedAt.Format(time.RFC3339))
	s.Nil(err)
	s.True(parsedTime.After(time.Date(1900, time.January, 1, 0, 0, 0, 0, time.UTC)))
	s.True(parsedTime.Before(time.Now()))
	s.True(output.UpdatedAt.IsZero())

	s.mu.Lock()
	s.manager = output
	s.mu.Unlock()
}

func (s *E2EScenarioSuite) postUsersEstimator() {
	c := http.Client{}

	input := UserInput{
		Email:    "marie.doe1110@userland.com",
		UserName: "marie123",
		Name:     "Marie Doe",
		UserType: "estimator",
	}

	message, err := json.Marshal(input)
	s.Nil(err)

	payload := bytes.NewReader(message)

	r, err := c.Post("http://localhost:9000/api/v1/users", "application/json", payload)
	s.Nil(err)
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	s.Nil(err)

	s.Equal(http.StatusCreated, r.StatusCode)

	var output UserOutput
	err = json.Unmarshal(body, &output)
	s.Nil(err)

	_, err = uuid.Parse(output.UserID)
	s.Nil(err, "UserID should be a valid uuidv4")
	s.Equal(input.Email, output.Email)
	s.Equal(input.UserName, output.UserName)
	s.Equal(input.Name, output.Name)
	s.Equal(input.UserType, output.UserType)
	parsedTime, err := time.Parse(time.RFC3339, output.CreatedAt.Format(time.RFC3339))
	s.Nil(err)
	s.True(parsedTime.After(time.Date(1900, time.January, 1, 0, 0, 0, 0, time.UTC)))
	s.True(parsedTime.Before(time.Now()))
	s.True(output.UpdatedAt.IsZero())

	s.mu.Lock()
	s.estimator = output
	s.mu.Unlock()
}

func (s *E2EScenarioSuite) postBaselines() {
	c := http.Client{}

	input := BaselineInput{
		Code:        "RIT123456789",
		Review:      1,
		Title:       "Logistics Cost & Time Management",
		Description: "This project will streamline our internal processes and increase overall efficiency",
		StartMonth:  1,
		StartYear:   2024,
		Duration:    12,
		ManagerID:   s.manager.UserID,
		EstimatorID: s.estimator.UserID,
	}

	message, err := json.Marshal(input)
	s.Nil(err)

	payload := bytes.NewReader(message)

	r, err := c.Post("http://localhost:9000/api/v1/baselines", "application/json", payload)
	s.Nil(err)
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	s.Nil(err)

	s.Equal(http.StatusCreated, r.StatusCode)

	var output BaselineOutput
	err = json.Unmarshal(body, &output)
	s.Nil(err)

	_, err = uuid.Parse(output.BaselineID)
	s.Nil(err, "BaselineID should be a valid uuidv4")
	s.Equal(input.Code, output.Code)
	s.Equal(input.Review, output.Review)
	s.Equal(input.Title, output.Title)
	s.Equal(input.Description, output.Description)
	s.Equal(time.Date(input.StartYear, time.Month(input.StartMonth), 1, 0, 0, 0, 0, time.UTC).Format("2006-01-02"), output.StartDate)
	s.Equal(input.Duration, output.Duration)
	s.Equal(input.ManagerID, output.ManagerID)
	s.Equal(input.EstimatorID, output.EstimatorID)
	parsedTime, err := time.Parse(time.RFC3339, output.CreatedAt.Format(time.RFC3339))
	s.Nil(err)
	s.True(parsedTime.After(time.Date(1900, time.January, 1, 0, 0, 0, 0, time.UTC)))
	s.True(parsedTime.Before(time.Now()))
	s.True(output.UpdatedAt.IsZero())

	s.mu.Lock()
	s.baseline = output
	s.mu.Unlock()
}

type UserInput struct {
	Email    string `json:"email"`
	UserName string `json:"user_name"`
	Name     string `json:"name"`
	UserType string `json:"user_type"`
}

type UserOutput struct {
	UserID    string    `json:"user_id"`
	Email     string    `json:"email"`
	UserName  string    `json:"user_name"`
	Name      string    `json:"name"`
	UserType  string    `json:"user_type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type BaselineInput struct {
	Code        string `json:"code"`
	Review      int    `json:"review"`
	Title       string `json:"title"`
	Description string `json:"description"`
	StartMonth  int    `json:"start_month"`
	StartYear   int    `json:"start_year"`
	Duration    int    `json:"duration"`
	ManagerID   string `json:"manager_id"`
	EstimatorID string `json:"estimator_id"`
}

type BaselineOutput struct {
	BaselineID  string    `json:"baseline_id"`
	Code        string    `json:"code"`
	Review      int       `json:"review"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartDate   string    `json:"start_date" layout:"2006-01-02"`
	Duration    int       `json:"duration"`
	ManagerID   string    `json:"manager_id"`
	Mananger    string    `json:"manager"`
	EstimatorID string    `json:"estimator_id"`
	Estimator   string    `json:"estimator"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
