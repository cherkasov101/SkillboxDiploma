package result

import (
	"SkillboxDiploma/internal/MMSData"
	"SkillboxDiploma/internal/SMSData"
	"SkillboxDiploma/internal/billingData"
	"SkillboxDiploma/internal/emailData"
	"SkillboxDiploma/internal/incidentData"
	"SkillboxDiploma/internal/stateCodes"
	"SkillboxDiploma/internal/supportData"
	"SkillboxDiploma/internal/voiceCall"
)

type ResultT struct {
	Status bool        `json:"status"` // True, если все этапы сбора данных прошли успешно, False во всех остальных случаях
	Data   *ResultSetT `json:"data"`   // Заполнен, если все этапы сбора  данных прошли успешно, nil во всех остальных случаях
	Error  string      `json:"error"`  // Пустая строка, если все этапы сбора данных прошли успешно, в случае ошибки заполнено текстом ошибки
}

type ResultSetT struct {
	SMS       [][]SMSData.SMSData                `json:"sms"`
	MMS       [][]MMSData.MMSData                `json:"mms"`
	VoiceCall []voiceCall.VoiceData              `json:"voice_call"`
	Email     map[string][][]emailData.EmailData `json:"email"`
	Billing   billingData.BillingData            `json:"billing"`
	Support   []int                              `json:"support"`
	Incidents []incidentData.IncidentData        `json:"incident"`
}

func GetResultData() ResultT {
	resultT := ResultT{
		false,
		nil,
		"",
	}

	voiceCall, err := voiceCall.GetData()
	if err != nil {
		resultT.Error = err.Error()
		return resultT
	}

	resultSetT := ResultSetT{
		getSMSData(&resultT),
		getMMSData(&resultT),
		voiceCall,
		getEmailData(&resultT),
		billingData.GetData(),
		getSupportData(&resultT),
		getIncidentData(&resultT),
	}

	if resultT.Error == "" {
		resultT.Status = true
		resultT.Data = &resultSetT
	}

	return resultT
}

func getSMSData(resultT *ResultT) [][]SMSData.SMSData {
	smsData, err := SMSData.GetData()
	if err != nil {
		resultT.Error = err.Error()
		return nil
	}

	for i, s := range smsData {
		smsData[i].Сountry = stateCodes.GetName(s.Сountry)
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

	return [][]SMSData.SMSData{firstSMSData, secondSMSData}
}

func getMMSData(resultT *ResultT) [][]MMSData.MMSData {
	mmsData, err := MMSData.GetData()
	if err != nil {
		resultT.Error = err.Error()
		return nil
	}

	for i, m := range mmsData {
		mmsData[i].Country = stateCodes.GetName(m.Country)
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

	return [][]MMSData.MMSData{firstMMSData, secondMMSData}
}

func getEmailData(resultT *ResultT) map[string][][]emailData.EmailData {
	emData, err := emailData.GetData()
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

	resultMap := make(map[string][][]emailData.EmailData)

	for _, e := range emData {
		if len(resultMap[e.Country]) == 0 {
			resultMap[e.Country] = append(resultMap[e.Country], []emailData.EmailData{})
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
	suppData, err := supportData.GetData()
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

func getIncidentData(resultT *ResultT) []incidentData.IncidentData {
	incidData, err := incidentData.GetData()
	if err != nil {
		resultT.Error = err.Error()
		return nil
	}

	var newIncidentData []incidentData.IncidentData

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
