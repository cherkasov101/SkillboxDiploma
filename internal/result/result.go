package result

import (
	"SkillboxDiploma/internal/billing"
	"SkillboxDiploma/internal/codes"
	"SkillboxDiploma/internal/email"
	"SkillboxDiploma/internal/incident"
	"SkillboxDiploma/internal/mms"
	"SkillboxDiploma/internal/sms"
	"SkillboxDiploma/internal/support"
	"SkillboxDiploma/internal/voice"
)

type ResultT struct {
	Status bool        `json:"status"` // True, если все этапы сбора данных прошли успешно, False во всех остальных случаях
	Data   *ResultSetT `json:"data"`   // Заполнен, если все этапы сбора  данных прошли успешно, nil во всех остальных случаях
	Error  string      `json:"error"`  // Пустая строка, если все этапы сбора данных прошли успешно, в случае ошибки заполнено текстом ошибки
}

type ResultSetT struct {
	SMS       [][]sms.SMSData                `json:"sms"`
	MMS       [][]mms.MMSData                `json:"mms"`
	VoiceCall []voice.VoiceData              `json:"voice_call"`
	Email     map[string][][]email.EmailData `json:"email"`
	Billing   billing.BillingData            `json:"billing"`
	Support   []int                          `json:"support"`
	Incidents []incident.IncidentData        `json:"incident"`
}

func GetResultData() ResultT {
	resultT := ResultT{
		false,
		nil,
		"",
	}

	voiceCall, err := voice.GetData()
	if err != nil {
		resultT.Error = err.Error()
		return resultT
	}

	resultSetT := ResultSetT{
		getSMSData(&resultT),
		getMMSData(&resultT),
		voiceCall,
		getEmailData(&resultT),
		billing.GetData(),
		getSupportData(&resultT),
		getIncidentData(&resultT),
	}

	if resultT.Error == "" {
		resultT.Status = true
		resultT.Data = &resultSetT
	}

	return resultT
}

func getSMSData(resultT *ResultT) [][]sms.SMSData {
	smsData, err := sms.GetData()
	if err != nil {
		resultT.Error = err.Error()
		return nil
	}

	for i, s := range smsData {
		smsData[i].Сountry = codes.GetName(s.Сountry)
	}

	for i := 0; i < len(smsData); i++ {
		for j := i; j < len(smsData); j++ {
			if smsData[i].Provider > smsData[j].Provider {
				smsData[i], smsData[j] = smsData[j], smsData[i]
			}
		}
	}
	firstSMSData := smsData

	for i := 0; i < len(smsData); i++ {
		for j := i; j < len(smsData); j++ {
			if smsData[i].Сountry > smsData[j].Сountry {
				smsData[i], smsData[j] = smsData[j], smsData[i]
			}
		}
	}
	secondSMSData := smsData

	return [][]sms.SMSData{firstSMSData, secondSMSData}
}

func getMMSData(resultT *ResultT) [][]mms.MMSData {
	mmsData, err := mms.GetData()
	if err != nil {
		resultT.Error = err.Error()
		return nil
	}

	for i, m := range mmsData {
		mmsData[i].Country = codes.GetName(m.Country)
	}

	for i := 0; i < len(mmsData); i++ {
		for j := i; j < len(mmsData); j++ {
			if mmsData[i].Provider > mmsData[j].Provider {
				mmsData[i], mmsData[j] = mmsData[j], mmsData[i]
			}
		}
	}
	firstMMSData := mmsData

	for i := 0; i < len(mmsData); i++ {
		for j := i; j < len(mmsData); j++ {
			if mmsData[i].Country > mmsData[j].Country {
				mmsData[i], mmsData[j] = mmsData[j], mmsData[i]
			}
		}
	}
	secondMMSData := mmsData

	return [][]mms.MMSData{firstMMSData, secondMMSData}
}

func getEmailData(resultT *ResultT) map[string][][]email.EmailData {
	emData, err := email.GetData()
	if err != nil {
		resultT.Error = err.Error()
		return nil
	}

	for i := 0; i < len(emData); i++ {
		for j := i; j < len(emData); j++ {
			if emData[i].DeliveryTime > emData[j].DeliveryTime {
				emData[i], emData[j] = emData[j], emData[i]
			}
		}
	}

	resultMap := make(map[string][][]email.EmailData)

	for _, e := range emData {
		if len(resultMap[e.Country]) == 0 {
			resultMap[e.Country] = append(resultMap[e.Country], []email.EmailData{})
		}
	}

	for _, e := range emData {
		if len(resultMap[e.Country][0]) < 3 {
			resultMap[e.Country][0] = append(resultMap[e.Country][0], e)
		}
	}

	for i := len(emData) - 1; i >= 0; i-- {
		if len(resultMap[emData[i].Country][0]) < 3 {
			resultMap[emData[i].Country][0] = append(resultMap[emData[i].Country][0], emData[i])
		}
	}

	return resultMap
}

func getSupportData(resultT *ResultT) []int {
	suppData, err := support.GetData()
	if err != nil {
		resultT.Error = err.Error()
		return nil
	}

	tikets := 0

	for _, s := range suppData {
		tikets += s.ActiveTickets
	}

	var load int

	if tikets < 9 {
		load = 1
	} else if tikets <= 16 {
		load = 2
	} else {
		load = 3
	}

	time := (60 / 18) * tikets

	return []int{load, time}
}

func getIncidentData(resultT *ResultT) []incident.IncidentData {
	incidData, err := incident.GetData()
	if err != nil {
		resultT.Error = err.Error()
		return nil
	}

	var newIncidentData []incident.IncidentData

	for _, i := range incidData {
		if i.Status == "active" {
			newIncidentData = append(newIncidentData, i)
		}
	}

	for _, i := range incidData {
		if i.Status == "closed" {
			newIncidentData = append(newIncidentData, i)
		}
	}

	return newIncidentData
}
