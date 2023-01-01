package support

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type SupportData struct {
	Topic         string `json:"topic"`
	ActiveTickets int    `json:"active_tickets"`
}

func GetData() ([]SupportData, error) {
	var data []SupportData

	resp, err := http.Get("http://127.0.0.1:8383/support")
	if err != nil {
		log.Printf(err.Error())
		return data, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if resp.StatusCode > 299 {
		log.Printf("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, body)
		return data, err
	}
	if err != nil {
		log.Printf(err.Error())
		return data, err
	}

	if err := json.Unmarshal(body, &data); err != nil {
		log.Printf(err.Error())
		return data, err
	}

	return data, nil
}
