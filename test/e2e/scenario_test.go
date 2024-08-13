package e2e_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/celsopires1999/estimation/internal/domain"
	"github.com/celsopires1999/estimation/internal/testutils"
	"github.com/golang-migrate/migrate/v4"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/suite"
)

type E2EScenarioSuite struct {
	suite.Suite
	dbpool          *pgxpool.Pool
	m               *migrate.Migrate
	manager         userOutput
	estimator       userOutput
	baseline        baselineOutput
	costPO          costOutput
	costConsulting  costOutput
	planBP          planOutput
	planFC03        planOutput
	portfolioIDBP   portfolioIDOutput
	portfolioIDFC03 portfolioIDOutput
	mu              sync.Mutex
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
	s.postCostsPO()
	s.postCostsConsulting()
	s.postPlanBP()
	s.postPlanFC03()
	s.postPortfolioBP()
	s.postPortfolioFC03()
	s.getPortfolioBP()
	s.getPortfolioFC03()
}

func (s *E2EScenarioSuite) postUsersManager() {
	c := http.Client{}

	input := userInput{
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

	var output userOutput
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

	input := userInput{
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

	var output userOutput
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

	input := baselineInput{
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

	var output baselineOutput
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

func (s *E2EScenarioSuite) postCostsPO() {
	c := http.Client{}

	input := costInput{
		CostType:       "one_time",
		Description:    "Mão de obra do PO",
		Comment:        "estimativa do PO",
		Amount:         180_000,
		Currency:       "BRL",
		Tax:            0.00,
		ApplyInflation: true,
		CostAllocations: []costAllocationInput{
			{
				Year:   2024,
				Month:  1,
				Amount: 30_000,
			},
			{
				Year:   2024,
				Month:  2,
				Amount: 30_000,
			},
			{
				Year:   2024,
				Month:  3,
				Amount: 30_000,
			},
			{
				Year:   2024,
				Month:  4,
				Amount: 30_000,
			},
			{
				Year:   2024,
				Month:  5,
				Amount: 30_000,
			},
			{
				Year:   2024,
				Month:  6,
				Amount: 30_000,
			},
		},
	}

	message, err := json.Marshal(input)
	s.Nil(err)

	payload := bytes.NewReader(message)

	r, err := c.Post("http://localhost:9000/api/v1/baselines/"+s.baseline.BaselineID+"/costs", "application/json", payload)
	s.Nil(err)
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	s.Nil(err)

	s.Equal(http.StatusCreated, r.StatusCode)

	var output costOutput
	err = json.Unmarshal(body, &output)
	s.Nil(err)
	s.Equal(input.CostType, output.CostType)
	s.Equal(input.Description, output.Description)
	s.Equal(input.Comment, output.Comment)
	s.Equal(input.Amount, output.Amount)
	s.Equal(input.Currency, output.Currency)
	s.Equal(input.Tax, output.Tax)
	s.Equal(input.ApplyInflation, output.ApplyInflation)
	s.Len(output.CostAllocations, len(input.CostAllocations))
	for i, a := range output.CostAllocations {
		s.Equal(input.CostAllocations[i].Year, a.Year)
		s.Equal(input.CostAllocations[i].Month, a.Month)
		s.Equal(input.CostAllocations[i].Amount, a.Amount)
	}

	s.mu.Lock()
	s.costPO = output
	s.mu.Unlock()
}

func (s *E2EScenarioSuite) postCostsConsulting() {
	c := http.Client{}

	input := costInput{
		CostType:       "one_time",
		Description:    "External Consulting",
		Comment:        "estimativa de consultoria externa",
		Amount:         80_000,
		Currency:       "EUR",
		Tax:            23.10,
		ApplyInflation: false,
		CostAllocations: []costAllocationInput{
			{
				Year:   2024,
				Month:  4,
				Amount: 30_000,
			},
			{
				Year:   2024,
				Month:  6,
				Amount: 50_000,
			},
		},
	}

	message, err := json.Marshal(input)
	s.Nil(err)

	payload := bytes.NewReader(message)

	r, err := c.Post("http://localhost:9000/api/v1/baselines/"+s.baseline.BaselineID+"/costs", "application/json", payload)
	s.Nil(err)
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	s.Nil(err)

	s.Equal(http.StatusCreated, r.StatusCode)

	var output costOutput
	err = json.Unmarshal(body, &output)
	s.Nil(err)
	s.Equal(input.CostType, output.CostType)
	s.Equal(input.Description, output.Description)
	s.Equal(input.Comment, output.Comment)
	s.Equal(input.Amount, output.Amount)
	s.Equal(input.Currency, output.Currency)
	s.Equal(input.Tax, output.Tax)
	s.Equal(input.ApplyInflation, output.ApplyInflation)
	s.Len(output.CostAllocations, len(input.CostAllocations))
	for i, a := range output.CostAllocations {
		s.Equal(input.CostAllocations[i].Year, a.Year)
		s.Equal(input.CostAllocations[i].Month, a.Month)
		s.Equal(input.CostAllocations[i].Amount, a.Amount)
	}

	s.mu.Lock()
	s.costConsulting = output
	s.mu.Unlock()
}

func (s *E2EScenarioSuite) postPlanBP() {
	c := http.Client{}

	input := planInput{
		Code: "BP 2025",
		Name: "Business Plan",
		Assumptions: domain.Assumptions{
			domain.Assumption{
				Year:      2024,
				Inflation: 4.00,
				Currencies: []domain.CurrencyAssumption{{
					Currency: domain.USD,
					Exchange: 4.50,
				}, {
					Currency: domain.EUR,
					Exchange: 5.50,
				}},
			},
			domain.Assumption{
				Year:      2025,
				Inflation: 5.20,
				Currencies: []domain.CurrencyAssumption{{
					Currency: domain.USD,
					Exchange: 5.00,
				}, {
					Currency: domain.EUR,
					Exchange: 6.00,
				}},
			},
			domain.Assumption{
				Year:      2026,
				Inflation: 5.26,
				Currencies: []domain.CurrencyAssumption{{
					Currency: domain.USD,
					Exchange: 5.55,
				}, {
					Currency: domain.EUR,
					Exchange: 6.66,
				}},
			},
			domain.Assumption{
				Year:      2027,
				Inflation: 5.30,
				Currencies: []domain.CurrencyAssumption{{
					Currency: domain.USD,
					Exchange: 5.77,
				}, {
					Currency: domain.EUR,
					Exchange: 6.88,
				}},
			},
		},
	}

	message, err := json.Marshal(input)
	s.Nil(err)

	payload := bytes.NewReader(message)

	r, err := c.Post("http://localhost:9000/api/v1/plans", "application/json", payload)
	s.Nil(err)
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	s.Nil(err)

	s.Equal(http.StatusCreated, r.StatusCode)

	var output planOutput
	err = json.Unmarshal(body, &output)
	s.Nil(err)
	s.Equal(input.Code, output.Code)
	s.Equal(input.Name, output.Name)
	s.Equal(input.Assumptions, output.Assumptions)

	s.mu.Lock()
	s.planBP = output
	s.mu.Unlock()
}

func (s *E2EScenarioSuite) postPlanFC03() {
	c := http.Client{}

	input := planInput{
		Code: "FC 03 2025",
		Name: "Forecast 03 2025",
		Assumptions: domain.Assumptions{
			domain.Assumption{
				Year:      2024,
				Inflation: 4.00,
				Currencies: []domain.CurrencyAssumption{{
					Currency: domain.USD,
					Exchange: 4.50,
				}, {
					Currency: domain.EUR,
					Exchange: 5.50,
				}},
			},
			domain.Assumption{
				Year:      2025,
				Inflation: 6.20,
				Currencies: []domain.CurrencyAssumption{{
					Currency: domain.USD,
					Exchange: 6.00,
				}, {
					Currency: domain.EUR,
					Exchange: 7.00,
				}},
			},
			domain.Assumption{
				Year:      2026,
				Inflation: 6.26,
				Currencies: []domain.CurrencyAssumption{{
					Currency: domain.USD,
					Exchange: 6.55,
				}, {
					Currency: domain.EUR,
					Exchange: 7.66,
				}},
			},
			domain.Assumption{
				Year:      2027,
				Inflation: 6.30,
				Currencies: []domain.CurrencyAssumption{{
					Currency: domain.USD,
					Exchange: 6.77,
				}, {
					Currency: domain.EUR,
					Exchange: 7.88,
				}},
			},
		},
	}

	message, err := json.Marshal(input)
	s.Nil(err)

	payload := bytes.NewReader(message)

	r, err := c.Post("http://localhost:9000/api/v1/plans", "application/json", payload)
	s.Nil(err)
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	s.Nil(err)

	s.Equal(http.StatusCreated, r.StatusCode)

	var output planOutput
	err = json.Unmarshal(body, &output)
	s.Nil(err)
	s.Equal(input.Code, output.Code)
	s.Equal(input.Name, output.Name)
	s.Equal(input.Assumptions, output.Assumptions)

	s.mu.Lock()
	s.planFC03 = output
	s.mu.Unlock()
}

func (s *E2EScenarioSuite) postPortfolioBP() {
	c := http.Client{}

	input := portfolioInput{
		BaselineID:  s.baseline.BaselineID,
		PlanID:      s.planBP.PlanID,
		ShiftMonths: 11,
	}

	message, err := json.Marshal(input)
	s.Nil(err)

	payload := bytes.NewReader(message)

	r, err := c.Post("http://localhost:9000/api/v1/portfolios", "application/json", payload)
	s.Nil(err)
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	s.Nil(err)

	s.Equal(http.StatusCreated, r.StatusCode)

	var output portfolioIDOutput
	err = json.Unmarshal(body, &output)
	s.Nil(err)

	_, err = uuid.Parse(output.PortfolioID)
	s.Nil(err, "PortfolioID should be a valid uuidv4")

	s.mu.Lock()
	s.portfolioIDBP = output
	s.mu.Unlock()
}

func (s *E2EScenarioSuite) postPortfolioFC03() {
	c := http.Client{}

	input := portfolioInput{
		BaselineID:  s.baseline.BaselineID,
		PlanID:      s.planFC03.PlanID,
		ShiftMonths: 18,
	}

	message, err := json.Marshal(input)
	s.Nil(err)

	payload := bytes.NewReader(message)

	r, err := c.Post("http://localhost:9000/api/v1/portfolios", "application/json", payload)
	s.Nil(err)
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	s.Nil(err)

	s.Equal(http.StatusCreated, r.StatusCode)

	var output portfolioIDOutput
	err = json.Unmarshal(body, &output)
	s.Nil(err)

	_, err = uuid.Parse(output.PortfolioID)
	s.Nil(err, "PortfolioID should be a valid uuidv4")

	s.mu.Lock()
	s.portfolioIDFC03 = output
	s.mu.Unlock()
}

func (s *E2EScenarioSuite) getPortfolioBP() {
	c := http.Client{}

	r, err := c.Get("http://localhost:9000/api/v1/portfolios/" + s.portfolioIDBP.PortfolioID)
	s.Nil(err)
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	s.Nil(err)

	s.Equal(http.StatusOK, r.StatusCode)

	var output portfolioOutput
	err = json.Unmarshal(body, &output)
	s.Nil(err)

	_, err = uuid.Parse(output.PortfolioID)
	s.Nil(err, "PortfolioID should be a valid uuidv4")

	expected := expectedPortfolio("testdata/expected_portfolio_BP.json")

	if !cmp.Equal(expected, output,
		cmpopts.IgnoreFields(portfolioOutput{}, "PortfolioID", "CreatedAt", "UpdatedAt"),
		cmpopts.IgnoreFields(budgetOutput{}, "BudgetID", "PortfolioID", "CreatedAt", "UpdatedAt"),
	) {
		s.T().Logf("%s", cmp.Diff(expected, output,
			cmpopts.IgnoreFields(portfolioOutput{}, "PortfolioID", "CreatedAt", "UpdatedAt"),
			cmpopts.IgnoreFields(budgetOutput{}, "BudgetID", "PortfolioID", "CreatedAt", "UpdatedAt"),
		))
		s.Fail("Portfolio is not equal to expected")
	}
}

func expectedPortfolio(filename string) portfolioOutput {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var buffer bytes.Buffer
	if _, err := io.Copy(&buffer, f); err != nil {
		panic(err)
	}

	var output portfolioOutput

	err = json.Unmarshal(buffer.Bytes(), &output)
	if err != nil {
		panic(err)
	}

	return output
}

func (s *E2EScenarioSuite) getPortfolioFC03() {
	c := http.Client{}

	r, err := c.Get("http://localhost:9000/api/v1/portfolios/" + s.portfolioIDFC03.PortfolioID)
	s.Nil(err)
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	s.Nil(err)

	s.Equal(http.StatusOK, r.StatusCode)

	var output portfolioOutput
	err = json.Unmarshal(body, &output)
	s.Nil(err)

	_, err = uuid.Parse(output.PortfolioID)
	s.Nil(err, "PortfolioID should be a valid uuidv4")

	expected := expectedPortfolio("testdata/expected_portfolio_FC03.json")

	if !cmp.Equal(expected, output,
		cmpopts.IgnoreFields(portfolioOutput{}, "PortfolioID", "CreatedAt", "UpdatedAt"),
		cmpopts.IgnoreFields(budgetOutput{}, "BudgetID", "PortfolioID", "CreatedAt", "UpdatedAt"),
	) {
		s.T().Logf("%s", cmp.Diff(expected, output,
			cmpopts.IgnoreFields(portfolioOutput{}, "PortfolioID", "CreatedAt", "UpdatedAt"),
			cmpopts.IgnoreFields(budgetOutput{}, "BudgetID", "PortfolioID", "CreatedAt", "UpdatedAt"),
		))
	}
}
