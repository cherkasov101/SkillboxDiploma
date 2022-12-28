package main

import (
	"SkillboxDiploma/pkg/MMSData"
	"fmt"
)

func main() {
	data, err := MMSData.GetData()
	if err != nil {
		fmt.Println(err)
	}

	for _, m := range data {
		fmt.Println(m.Provider)
	}
}
