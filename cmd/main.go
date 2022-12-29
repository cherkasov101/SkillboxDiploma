package main

import (
	"SkillboxDiploma/pkg/emailData"
	"fmt"
)

func main() {
	data, err := emailData.GetData()
	if err != nil {
		fmt.Println(err)
	}

	for _, m := range data {
		fmt.Println(m.Provider)
	}
}
