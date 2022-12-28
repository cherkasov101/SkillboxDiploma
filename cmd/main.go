package main

import (
	"SkillboxDiploma/pkg/SMSData"
	"fmt"
)

func main() {
	smsData, err := SMSData.GetData()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, s := range smsData {
		fmt.Println(s.Сountry)
	}
}
