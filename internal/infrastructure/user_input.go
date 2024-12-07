package infrastructure

import (
	"fmt"
	"strconv"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	"github.com/manifoldco/promptui"
)

func GetValue(message string, minValue, maxValue float64) (float64, error) {
	validate := func(input string) error {
		value, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return &domain.InvalidInput{Message: "Invalid input."}
		}

		if value < minValue || value > maxValue {
			return &domain.InvalidInput{Message: "Invalid input."}
		}

		return nil
	}

	// Get value.
	promptValue := promptui.Prompt{
		Label:    message,
		Validate: validate,
	}

	ValueStr, err := promptValue.Run()
	if err != nil {
		return 0, fmt.Errorf("prompting value: %w", err)
	}

	value, err := strconv.ParseFloat(ValueStr, 64)
	if err != nil {
		return 0, fmt.Errorf("converting value to float64: %w", err)
	}

	return value, nil
}

func GetTrueOrFalse(message string) (bool, error) {
	promt := promptui.Select{
		Label: message,
		Items: []string{"Yes", "No"},
	}

	_, result, err := promt.Run()
	if err != nil {
		return false, fmt.Errorf("prompting value: %w", err)
	}

	if result == "Yes" {
		return true, nil
	}

	return false, nil
}
