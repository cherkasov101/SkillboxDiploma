package stateCodes

import (
	"encoding/csv"
	"fmt"
	"os"
)

var (
	stateCodesMap map[string]string
	fileName      = "../pkg/stateCodes/states.scv"
)

func init() {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	states, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		return
	}

	stateCodesMap = make(map[string]string)

	for _, st := range states {
		stateCodesMap[st[0]] = st[1]
	}
}

func IsExist(code string) bool {
	if stateCodesMap[code] == "" {
		return false
	}
	return true
}
