package SMSData

import (
	"SkillboxDiploma/pkg/stateCodes"
	"log"
	"os"
	"strings"
)

var fileName = "../../skillbox-diploma/sms.data"

type SMSData struct {
	Сountry      string `json:"сountry"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"responseTime"`
	Provider     string `json:"provider"`
}

func GetData() ([]SMSData, error) {
	bytesData, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var data []SMSData

	content := strings.Split(string(bytesData), "\n")
	for _, sms := range content {
		s := strings.Split(sms, ";")
		if len(s) == 4 && stateCodes.IsExist(s[0]) {
			checkProvider := s[3] == "Topolo" || s[3] == "Rond" || s[3] == "Kildy"
			if checkProvider {
				newSMS := SMSData{
					s[0],
					s[1],
					s[2],
					s[3],
				}
				data = append(data, newSMS)
			}
		}
	}

	return data, nil
}
