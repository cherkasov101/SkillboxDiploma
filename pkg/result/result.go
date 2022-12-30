package result

import (
	"SkillboxDiploma/pkg/MMSData"
	"SkillboxDiploma/pkg/SMSData"
	"SkillboxDiploma/pkg/billingData"
	"SkillboxDiploma/pkg/emailData"
	"SkillboxDiploma/pkg/incidentData"
	"SkillboxDiploma/pkg/voiceCall"
)

type ResultT struct {
	Status bool       `json:"status"` // True, если все этапы сбора данных прошли успешно, False во всех остальных случаях
	Data   ResultSetT `json:"data"`   // Заполнен, если все этапы сбора  данных прошли успешно, nil во всех остальных случаях
	Error  string     `json:"error"`  // Пустая строка, если все этапы сбора данных прошли успешно, в случае ошибки заполнено текстом ошибки
}

type ResultSetT struct {
	SMS       [][]SMSData.SMSData                `json:"sms"`
	MMS       [][]MMSData.MMSData                `json:"mms"`
	VoiceCall []voiceCall.VoiceData              `json:"voice_call"`
	Email     map[string][][]emailData.EmailData `json: email"`
	Billing   billingData.BillingData            `json: billing"`
	Support   []int                              `json: support"`
	Incidents []incidentData.IncidentData        `json:"incident"`
}
