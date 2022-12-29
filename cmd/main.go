package main

import (
	"SkillboxDiploma/pkg/SMSData"
	"fmt"
)

func main() {
	data, err := SMSData.GetData()
	if err != nil {
		fmt.Println(err)
	}

	for _, m := range data {
		fmt.Println(m.Provider)
	}
}
