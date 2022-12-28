package SMSData

import (
	"SkillboxDiploma/pkg/stateCodes"
	"encoding/csv"
	"fmt"
	"os"
)

var fileName = "../../skillbox-diploma/sms.data"

type SMSData struct {
	Сountry      string
	Bandwidth    string
	ResponseTime string
	Provider     string
}

func GetData() ([]SMSData, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ';'
	sms, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}

	var data []SMSData

	for _, stateData := range sms {
		checkProvider := stateData[3] == "Topolo" || stateData[3] == "Rond" || stateData[3] == "Kildy"
		if len(stateData) == 4 && stateCodes.IsExist(stateData[0]) && checkProvider {
			s := SMSData{
				stateData[0],
				stateData[1],
				stateData[2],
				stateData[3],
			}
			data = append(data, s)
		}
	}

	return data, nil
}
