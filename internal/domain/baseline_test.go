package domain_test

import (
	"errors"
	"testing"

	"github.com/celsopires1999/estimation/internal/domain"
	"github.com/celsopires1999/estimation/internal/testutils"
	"github.com/stretchr/testify/assert"
)

func TestUnitBaseline(t *testing.T) {
	t.Run("should create a baseline with valid values", func(t *testing.T) {
		faker := testutils.NewBaselineFakeBuilder()

		baseline := domain.NewBaseline(
			faker.Code,
			faker.Title,
			faker.Description,
			faker.StartDate,
			faker.Duration,
			faker.ManagerID,
			faker.EstimatorID,
		)

		assert.Equal(t, faker.Title, baseline.Title)
		assert.Equal(t, faker.Description, baseline.Description)
		assert.Equal(t, faker.StartDate, baseline.StartDate)
		assert.Equal(t, faker.ManagerID, baseline.ManagerID)
		assert.Equal(t, faker.EstimatorID, baseline.EstimatorID)
	})

	t.Run("should fail to create a baseline with invalid values", func(t *testing.T) {
		faker := testutils.NewBaselineFakeBuilder()

		code := ""
		title := ""
		description := ""
		startDate := faker.StartDate
		duration := int32(0)
		managerID := ""
		estimatorID := ""

		baseline := domain.NewBaseline(
			code,
			title,
			description,
			startDate,
			duration,
			managerID,
			estimatorID,
		)

		err := baseline.Validate()
		assert.Error(t, err)
		assert.True(t, errors.Is(err, domain.ErrBaselineDomainValidation))
	})
}
