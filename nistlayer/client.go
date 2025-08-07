package nistlayer

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func GETCVEs(httpClient http.Client, cpe string) (LongCveResponse, error) {
	url := fmt.Sprintf("https://services.nvd.nist.gov/rest/json/cves/2.0?cpeName=%s", cpe)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LongCveResponse{}, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("accept", "application/json")

	res, err := httpClient.Do(req)
	if err != nil {
		return LongCveResponse{}, fmt.Errorf("error sending request: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LongCveResponse{}, fmt.Errorf("error reading response body: %w", err)
	}

	log.Printf("Reponse Body: %v", body)

	var resJson LongCveResponse
	err = json.Unmarshal(body, &resJson)
	if err != nil {
		fmt.Println(err)
		return LongCveResponse{}, fmt.Errorf("error unmarshaling response body: %w", err)
	}

	return resJson, nil

}
