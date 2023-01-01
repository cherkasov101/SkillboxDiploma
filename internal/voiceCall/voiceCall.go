package voiceCall

import (
	"SkillboxDiploma/internal/stateCodes"
	"log"
	"os"
	"strconv"
	"strings"
)

var fileName = "../../skillbox-diploma/voice.data"

type VoiceData struct {
	Сountry             string  `json:"сountry"`
	Load                string  `json:"load"`
	ResponseTime        string  `json:"responseTime"`
	Provider            string  `json:"provider"`
	connectionStability float32 `json:"connectionStability"`
	Frequency           string  `json:"frequency"`
	CallDuration        string  `json:"callDuration"`
	Unknown             string  `json:"unknown"`
}

func GetData() ([]VoiceData, error) {
	bytesData, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var data []VoiceData

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
				voiceCall := VoiceData{
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
