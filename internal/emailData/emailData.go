package emailData

import (
	"SkillboxDiploma/internal/stateCodes"
	"log"
	"os"
	"strconv"
	"strings"
)

var fileName = "../../skillbox-diploma/email.data"

type EmailData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	DeliveryTime int    `json:"deliveryTime"`
}

var correctProviders = []string{
	"Gmail",
	"Yahoo",
	"Hotmail",
	"MSN",
	"Orange",
	"Comcast",
	"AOL",
	"Live",
	"RediffMail",
	"GMX",
	"Proton Mail",
	"Yandex",
	"Mail.ru",
}

func GetData() ([]EmailData, error) {
	bytesData, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var data []EmailData

	content := strings.Split(string(bytesData), "\n")
	for _, email := range content {
		e := strings.Split(email, ";")
		if len(e) == 3 && stateCodes.IsExist(e[0]) {
			if checkProvider(e[1]) {
				time, err := strconv.Atoi(e[2])
				if err != nil {
					log.Fatal(err)
					return nil, err
				}
				eD := EmailData{
					e[0],
					e[1],
					time,
				}
				data = append(data, eD)
			}
		}
	}

	return data, nil
}

func checkProvider(provider string) bool {
	for _, p := range correctProviders {
		if p == provider {
			return true
		}
	}
	return false
}
