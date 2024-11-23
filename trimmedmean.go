package trimmedmean

import (
	"errors"
	"sort"
)

func TrimmedMeanInt(data []int, trimLower float64, trimUpper ...float64) (float64, error) {
	if len(data) == 0 {
		return 0, errors.New("data slice is empty")
	}
	var trimUpperVal float64
	if len(trimUpper) == 0 {
		trimUpperVal = trimLower
	} else {
		trimUpperVal = trimUpper[0]
	}

	if trimLower < 0 || trimUpperVal < 0 || trimLower >= 0.5 || trimUpperVal >= 0.5 {
		return 0, errors.New("trim proportions should be between 0 and 0.5")
	}

	sortedData := make([]int, len(data))
	copy(sortedData, data)
	sort.Ints(sortedData)

	n := len(sortedData)
	lowerIndex := int(float64(n) * trimLower)
	upperIndex := n - int(float64(n)*trimUpperVal)

	trimmedData := sortedData[lowerIndex:upperIndex]

	sum := 0
	for _, v := range trimmedData {
		sum += v
	}
	mean := float64(sum) / float64(len(trimmedData))

	return mean, nil
}

// TrimmedMeanFloat computes the trimmed mean of a slice of floats.
func TrimmedMeanFloat(data []float64, trimLower float64, trimUpper ...float64) (float64, error) {
	if len(data) == 0 {
		return 0, errors.New("data slice is empty")
	}
	var trimUpperVal float64
	if len(trimUpper) == 0 {
		trimUpperVal = trimLower
	} else {
		trimUpperVal = trimUpper[0]
	}

	if trimLower < 0 || trimUpperVal < 0 || trimLower >= 0.5 || trimUpperVal >= 0.5 {
		return 0, errors.New("trim proportions should be between 0 and 0.5")
	}

	sortedData := make([]float64, len(data))
	copy(sortedData, data)
	sort.Float64s(sortedData)

	n := len(sortedData)
	lowerIndex := int(float64(n) * trimLower)
	upperIndex := n - int(float64(n)*trimUpperVal)

	trimmedData := sortedData[lowerIndex:upperIndex]

	sum := 0.0
	for _, v := range trimmedData {
		sum += v
	}
	mean := sum / float64(len(trimmedData))

	return mean, nil
}
