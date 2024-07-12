package operation

import (
	"fmt"
	"os"
	"strconv"
)

func ReadFiles(fileName string) float64 {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 0
	}
	textConvertion := string(data)
	float, err := strconv.ParseFloat(textConvertion, 64)
	if err != nil {
		return 0
	}
	return float

}

func WriteFiles(writeFiles float64, writeFloatAsString string) {
	data := fmt.Sprint(writeFiles)
	os.WriteFile(writeFloatAsString, []byte(data), 0644)

}
