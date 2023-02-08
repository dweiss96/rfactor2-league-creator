package main

import (
	"net/http"
	"io"
	"fmt"
	"encoding/json"
)

func getRaceAppSeries(raceAppId uint64) (RaceAppResponse, error) {
	client := &http.Client{}
	var result RaceAppResponse

	req, err := http.NewRequest("GET", fmt.Sprintf("https://raceapp.eu/api/series/%d", raceAppId), nil)
	if(err != nil) {
		return result, err
	}
	
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)
	if(err != nil) {
		return result, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if(err != nil) {
		return result, err
	}

	err = json.Unmarshal(body, &result)
	if(err != nil) {
		return result, err
	}
	return result, nil
}

