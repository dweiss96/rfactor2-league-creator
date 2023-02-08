package main

type DriverDefinition struct {
	Name string `json:"name"`
	Number  string `json:"number"`
	Car  string `json:"car"`
	TeamName  string `json:"teamName"`
}

type RaceAppResponse struct {
	Settings RaceAppSettings `json:"Settings"`
}
type RaceAppSettings struct {
	SoloBookings []RaceAppSoloBooking `json:"CarSoloBookings"`
	TeamBookings []RaceAppTeamBooking `json:"CarTeamBookings"`
}

type RaceAppSoloBooking struct {
	Drivers []RaceAppDriver `json:"Drivers"`
	Model  string `json:"Model"`
	Number  string `json:"CarNumber"`
}

type RaceAppTeamBooking struct {
	Drivers []RaceAppDriver `json:"Drivers"`
	Model  string `json:"Model"`
	Number  string `json:"CarNumber"`
	TeamName  string `json:"TeamName"`
}

type RaceAppDriver struct {
	Name string `json:"Name"`
}