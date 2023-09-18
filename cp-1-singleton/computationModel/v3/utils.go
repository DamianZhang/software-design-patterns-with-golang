package computationModel

import (
	"fmt"
	"strconv"
)

func ParseFloat64Slice(stringSlice []string) ([]float64, error) {
	float64Slice := make([]float64, 0)

	for _, s := range stringSlice {
		num, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return nil, fmt.Errorf("parse float64 slice failed to parse float: %s", err)
		}

		float64Slice = append(float64Slice, num)
	}

	return float64Slice, nil
}
