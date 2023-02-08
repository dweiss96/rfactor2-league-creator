package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"os"
	"io/ioutil"
)

func generateDir() error {
	return os.MkdirAll(fmt.Sprintf("./_generated_outputs"), os.ModeDir)
}

func removeOutputDir() error {
	return os.RemoveAll(fmt.Sprintf("./_generated_outputs"))
}

func generate(config ApplicationConfiguration) error {
	err := generateDir()
	if(err != nil) {
		return err
	}
	// fetch raceapp api
	raceAppData, err := getRaceAppSeries(config.RaceAppLeagueId)
	if(err != nil) {
		return err
	}
	// generate 
	var drivers []DriverDefinition
	for _, driver := range raceAppData.Settings.SoloBookings {
		if(len(driver.Drivers) > 0) {
			drivers = append(drivers, DriverDefinition{
				Name: driver.Drivers[0].Name,
				TeamName: driver.Drivers[0].Name,
				Number: driver.Number,
				Car: config.CarMappings[driver.Model],
			})
		}
	}
	for _, driver := range raceAppData.Settings.TeamBookings {
		if(len(driver.Drivers) > 0) {
			drivers = append(drivers, DriverDefinition{
				Name: driver.Drivers[0].Name,
				TeamName: driver.TeamName,
				Number: driver.Number,
				Car: config.CarMappings[driver.Model],
			})
		}
	}
	fmt.Printf("%d driver converted\n", len(drivers))

	resultJson, err := json.Marshal(drivers)
	if(err != nil) {
		return err
	}

	writePhotopeaCommands(string(resultJson))
	writeError := ioutil.WriteFile("series_config.json", resultJson, 0644)
	if(writeError != nil) {
		return writeError
	}

	return nil
}

func writePhotopeaCommands(drivers string) error {
	numplateCmdTemplate := `
	let drivers = {{DRIVER_DEFS}}
	let numbers = Object.keys(driverNames)
	for(i=0; i<drivers.length; i++) {
		let number = drivers[i].number
		let name = drivers[i].name
	
		app.activeDocument.layers.getByName("Number1").textItem.contents = ""+number;
		app.activeDocument.layers.getByName("Number2").textItem.contents = ""+number;
	
		var opts = new ExportOptionsSaveForWeb();
			opts.format = SaveDocumentType.PNG;
			opts.PNG8 = false;
			opts.quality = 100;
			
			pngFile = new File("num"+number+".png");
			app.activeDocument.exportDocument(pngFile, ExportType.SAVEFORWEB, opts);
	}
	`
	windowCmdTemplate := `
	let drivers = {{DRIVER_DEFS}}
	let numbers = Object.keys(driverNames)
	for(i=0; i<drivers.length; i++) {
		let number = drivers[i].number
		let name = drivers[i].name
	
		app.activeDocument.layers.getByName("Number1").textItem.contents = ""+number;
		app.activeDocument.layers.getByName("Number2").textItem.contents = ""+number;
		app.activeDocument.layers.getByName("DriverName1").textItem.contents = ""+name;
		app.activeDocument.layers.getByName("DriverName2").textItem.contents = ""+name;
	
		var opts = new ExportOptionsSaveForWeb();
			opts.format = SaveDocumentType.PNG;
			opts.PNG8 = false;
			opts.quality = 100;
			
			pngFile = new File("windshield"+number+".png");
			app.activeDocument.exportDocument(pngFile, ExportType.SAVEFORWEB, opts);
	}
	`

	numplateCmd := strings.ReplaceAll(numplateCmdTemplate, "{{DRIVER_DEFS}}", drivers)
	windowCmd := strings.ReplaceAll(windowCmdTemplate, "{{DRIVER_DEFS}}", drivers)

	writeError := ioutil.WriteFile("numplateCmd.jsx", []byte(numplateCmd), 0644)
	if(writeError != nil) {
		return writeError
	}
	writeError = ioutil.WriteFile("windowCmd.jsx", []byte(windowCmd), 0644)
	if(writeError != nil) {
		return writeError
	}

	return nil
}
