package http

import (
	"errors"
	"net/url"
	"strconv"

	"github.com/celsopires1999/estimation/internal/infra/db"
)

func permittedValue[T comparable](value T, permittedValues ...T) bool {
	for i := range permittedValues {
		if value == permittedValues[i] {
			return true
		}
	}
	return false
}

func readString(qs url.Values, key string, defaultValue string) string {
	s := qs.Get(key)
	if s == "" {
		return defaultValue
	}

	return s
}

// func readCSV(qs url.Values, key string, defaultValue []string) []string {
// 	csv := qs.Get(key)

// 	if csv == "" {
// 		return defaultValue
// 	}

// 	return strings.Split(csv, ",")
// }

func readInt(qs url.Values, key string, defaultValue int) int {
	s := qs.Get(key)

	if s == "" {
		return defaultValue
	}

	i, err := strconv.Atoi(s)
	if err != nil {
		return defaultValue
	}

	return i
}

func setFilter(qs url.Values, filters *db.Filters) error {
	filters.Page = readInt(qs, "page", 1)
	filters.PageSize = readInt(qs, "page_size", 20)
	filters.Sort = readString(qs, "sort", "-created_at")

	if !permittedValue(filters.Sort, filters.SortSafelist...) {
		return errors.New("invalid sort value")
	}

	return nil
}
