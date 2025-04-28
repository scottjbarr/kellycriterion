package kellycriteron

import (
	"errors"
	"fmt"
)

// Calculate return the result of the Kelly Critereon formula.
//
// An error will be returned if invalid input is received.
func Calculate(winRate, reward float64) (float64, error) {
	if err := validate(winRate, reward); err != nil {
		return 0, err
	}

	// rate = (win rate * reward) - (loss rate) / reward
	rate := ((winRate * reward) - (1.0 - winRate)) / reward

	return rate, nil
}

func validate(winRate, reward float64) error {
	errs := []error{}

	if winRate <= 0 || winRate > 1.0 {
		errs = append(errs, fmt.Errorf("winRate must be >= 0 and <= 1.0"))
	}

	if reward <= 0 {
		errs = append(errs, fmt.Errorf("reward must be > 0"))
	}

	return errors.Join(errs...)
}
