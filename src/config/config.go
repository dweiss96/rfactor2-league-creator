package config

import (
	"fmt"
	"strings"
	"os"
	"io/ioutil"
	"encoding/json"
	"errors"
)

type ApplicationConfiguration struct {
    ModMgrPath    string `json:"modmgrPath"`
	WorkshopPath  string `json:"workshopPath"`
	CarConfig []CarConfigModel `json:"cars"`
	LeagueConfig LeagueConfigModel `json:"league"`
}

func (c *ApplicationConfiguration) FindCar(id string) (CarConfigModel, error) {
	for _, car := range c.CarConfig {
		if(car.Id == id) {
			return car, nil
		}
	}

	return c.CarConfig[0], errors.New("No Car Found")
}

type CarConfigModel struct {
    Id    string `json:"id"`
	WorkshopID  string `json:"workshopId"`
	Postfixes  []string `json:"possiblePostfixes"`
}

type LeagueConfigModel struct {
    Name    string `json:"name"`
	Driver  []DriverModel `json:"driver"`
}

type DriverModel struct {
	Name string `json:"name"`
	Team string `json:"team"`
	Number int16 `json:"number"`
	Car string `json:"car"`
}

func (d *DriverModel) GetLiveryIdentifier() string {
	namePart := ""
	for _, p := range strings.Split(d.Name, " ") {
		namePart = fmt.Sprintf("%s%c", namePart, p[0])
	}
	teamPart := ""
	for _, p := range strings.Split(d.Team, " ") {
		teamPart = fmt.Sprintf("%s%c", teamPart, p[0])
	}
	id := fmt.Sprintf("%d_%s_%s", d.Number, namePart, teamPart)
	return strings.ToUpper(id)
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

