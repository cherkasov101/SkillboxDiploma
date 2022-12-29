package voiceCall

import (
	"SkillboxDiploma/pkg/stateCodes"
	"log"
	"os"
	"strconv"
	"strings"
)

var fileName = "../../skillbox-diploma/voice.data"

type voiceData struct {
	Сountry             string
	Load                string
	ResponseTime        string
	Provider            string
	connectionStability float32
	Frequency           string
	CallDuration        string
	Unknown             string
}

func GetData() ([]voiceData, error) {
	bytesData, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var data []voiceData

	content := strings.Split(string(bytesData), "\n")
	for _, call := range content {
		c := strings.Split(call, ";")
		if len(c) == 8 && stateCodes.IsExist(c[0]) {
			checkProvider := c[3] == "TransparentCalls" || c[3] == "E-Voice" || c[3] == "JustPhone"
			if checkProvider {
				stab, err := strconv.ParseFloat(c[4], 32)
				if err != nil {
					log.Fatal(err)
					return nil, err
				}
				stability := float32(stab)
				voiceCall := voiceData{
					c[0],
					c[1],
					c[2],
					c[3],
					stability,
					c[5],
					c[6],
					c[7],
				}
				data = append(data, voiceCall)
			}
		}
	}

	return data, nil
}
