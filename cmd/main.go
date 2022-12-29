package main

import (
	"SkillboxDiploma/pkg/incidentData"
	"fmt"
)

func main() {
	data, err := incidentData.GetData()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, i := range data {
		fmt.Println(i.Status)
	}
}
