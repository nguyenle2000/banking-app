package operation

import (
	"fmt"
	"os"
	"strconv"
)

func WriteBalanceToFiles(fileName float64, value string) {
	converText := fmt.Sprint(fileName)
	os.WriteFile(value, []byte(converText), 0644)
}

func ReadBalanceFromFiles(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 0, fmt.Errorf("can't read file %v", data)
	}
	fileText := string(data)

	fileRead, err := strconv.ParseFloat(fileText, 64)
	if err != nil {
		return 0, err
	}
	return fileRead, nil

}
