package sortfilter

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrInvalidSortByQueryParam     = errors.New("malformed sortBy query parameter, should be field.orderdirection")
	ErrInvalidSortByOrderdirection = errors.New("malformed orderdirection in sortBy query parameter, should be asc or desc")
	ErrInvalidSortField            = errors.New("unknown field in sortBy query parameter")
)

func ValidateAndReturnSortQuery(sortBy string, sortFields []string) (string, error) {
	splits := strings.Split(sortBy, ".")
	if len(splits) != 2 {
		return "", fmt.Errorf("validate and return sort query: %w", ErrInvalidSortByQueryParam)
	}

	field, order := splits[0], splits[1]
	if order != "desc" && order != "asc" {
		return "", fmt.Errorf("validate and return sort query: %w", ErrInvalidSortByOrderdirection)
	}

	if !stringInSlice(sortFields, field) {
		return "", fmt.Errorf("validate and return sort query: %w: should be one of this: %v", ErrInvalidSortField, sortFields)
	}

	return fmt.Sprintf("ORDER BY %s %s", field, strings.ToUpper(order)), nil
}
