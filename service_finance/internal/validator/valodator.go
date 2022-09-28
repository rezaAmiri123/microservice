package validator

import (
	"errors"
	"fmt"

	"github.com/rezaAmiri123/microservice/service_finance/internal/utils"
)

var (
	ErrWrongCorrency   = errors.New("incorrect corrency")
	ErrBalancePositive = errors.New("balance must be grater than 0")
)

func ValidateString(value string, minLength int, maxLength int) error {
	n := len(value)
	if n < minLength || n > maxLength {
		return fmt.Errorf("must contain from %d-%d characters", minLength, maxLength)
	}
	return nil
}

func ValidateCurrency(value string) error {
	if !utils.IsSupportedCurrency(value) {
		return ErrWrongCorrency
	}
	return nil
}

func ValidateBalance(value int64) error {
	if value < 0 {
		return ErrBalancePositive
	}
	return nil
}
