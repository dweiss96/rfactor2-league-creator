package main

import (
	"os"
	"io/ioutil"
	"encoding/json"
)

type ApplicationConfiguration struct {
	RaceAppLeagueId uint64 `json:"raceAppLeagueId"`
	CarMappings  map[string]string `json:"carMappings"`
}

func ReadConfigFromJson(path string) ApplicationConfiguration {
	jsonFile, _ := os.Open(path)
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var conf ApplicationConfiguration

	err := json.Unmarshal(byteValue, &conf)
	if(err != nil) {
		panic(err)
	}

	return conf
}
