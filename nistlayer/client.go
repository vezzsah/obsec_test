package nistlayer

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GETCVEs(httpClient http.Client, cpe string) (ShortCVEResponse, error) {
	url := fmt.Sprintf("https://services.nvd.nist.gov/rest/json/cves/2.0?cpeName=%s", cpe)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ShortCVEResponse{}, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("accept", "application/dns-json")

	res, err := httpClient.Do(req)
	if err != nil {
		return ShortCVEResponse{}, fmt.Errorf("error sending request: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return ShortCVEResponse{}, fmt.Errorf("error reading response body: %w", err)
	}

	var result ShortCVEResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println(err)
		return ShortCVEResponse{}, fmt.Errorf("error unmarshaling response body: %w", err)
	}

	return result, nil

}
