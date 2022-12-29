package main

import (
	"SkillboxDiploma/pkg/billingData"
	"fmt"
)

func main() {
	data := billingData.GetData()

	fmt.Println(data)
}
