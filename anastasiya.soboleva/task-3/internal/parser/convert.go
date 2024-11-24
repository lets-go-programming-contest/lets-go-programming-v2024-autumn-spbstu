package parser

import (
	"fmt"
	"strconv"
	"strings"

	"anastasiya.soboleva/task-3/internal/models"
)

func ConvertToModels(valCurs models.ValCurs) ([]models.Currency, error) {
	var currencies []models.Currency
	for _, v := range valCurs.Valutes {
		numCode, err := strconv.Atoi(v.NumCode)
		if err != nil {
			return nil, fmt.Errorf("error converting NumCode %s to integer: %w", v.NumCode, err)
		}
		value, err := strconv.ParseFloat(strings.ReplaceAll(v.Value, ",", "."), 32)
		if err != nil {
			return nil, fmt.Errorf("error converting Value %s to float: %w", v.Value, err)
		}
		currencies = append(currencies, models.Currency{
			NumCode:  numCode,
			CharCode: v.CharCode,
			Value:    float32(value),
		})
	}
	return currencies, nil
}
