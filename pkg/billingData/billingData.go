package billingData

import (
	"log"
	"os"
	"strconv"
)

var fileName = "../../skillbox-diploma/billing.data"

type BillingData struct {
	CreateCustomer bool
	Purchase       bool
	Payout         bool
	Recurring      bool
	FraudControl   bool
	CheckoutPage   bool
}

func GetData() BillingData {
	bytesData, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	var statuses []bool

	for i := len(bytesData) - 1; i >= 0; i-- {
		status, err := strconv.ParseBool(string(bytesData[i]))
		if err != nil {
			log.Fatal(err)
		}
		statuses = append(statuses, status)
	}

	billing := BillingData{
		statuses[0],
		statuses[1],
		statuses[2],
		statuses[3],
		statuses[4],
		statuses[5],
	}

	return billing
}
