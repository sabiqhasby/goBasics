package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 1000, errors.New("fail to find balance file")
	}

	balanceText := string(data)
	value, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000, errors.New("failed to parse stored balance value")
	}
	return value, nil
}

func WriteFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}
