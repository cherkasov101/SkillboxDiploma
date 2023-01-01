package incidentData

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type IncidentData struct {
	Topic  string `json:"topic"`
	Status string `json:"status"` // возможные статусы active и closed
}

func GetData() ([]IncidentData, error) {
	var data []IncidentData

	resp, err := http.Get("http://127.0.0.1:8383/accendent")
	if err != nil {
		log.Fatal(err)
		return data, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if resp.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, body)
		return data, err
	}
	if err != nil {
		log.Fatal(err)
		return data, err
	}

	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatal(err)
		return data, err
	}

	return data, nil
}
