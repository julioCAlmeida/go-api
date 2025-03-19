package utils

import (
	"errors"
	"strings"
)

func ValidateProductName(name string) error {
	if strings.TrimSpace(name) == "" {
		return errors.New("product name cannot be empty")
	}

	if len(name) > 100 {
		return errors.New("product name cannot be more than 100 characters")
	}
	return nil
}