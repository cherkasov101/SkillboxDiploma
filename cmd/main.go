package main

import (
	"SkillboxDiploma/pkg/supportData"
	"fmt"
)

func main() {
	data, err := supportData.GetData()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, i := range data {
		fmt.Println(i.ActiveTickets)
	}
}
