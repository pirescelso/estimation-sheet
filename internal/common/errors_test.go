package common_test

import (
	"fmt"
	"testing"

	"github.com/celsopires1999/estimation/internal/common"
)

func TestUnitNotFoundError(t *testing.T) {
	t.Run("should return error message as string", func(t *testing.T) {
		err := common.NewNotFoundError(fmt.Errorf("test"))
		if err.Error() != "test" {
			t.Errorf("expected error message to be %s, but got %s", "test", err.Error())
		}
	})
}

func TestUnitConflictError(t *testing.T) {
	t.Run("should return error message as string", func(t *testing.T) {
		err := common.NewConflictError(fmt.Errorf("there is an error"))
		expected := "there is an error"
		if err.Error() != expected {
			t.Errorf("expected error message to be %s, but got %s", expected, err.Error())
		}
	})
}
