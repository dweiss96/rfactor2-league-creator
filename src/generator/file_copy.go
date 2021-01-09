package generator

import(
	"fmt"
	"strings"
	"../config"
	"../tools"
)

func copyDriverFilesToDir(driver config.DriverModel, config config.ApplicationConfiguration) {
	skinPath := fmt.Sprintf("./templates/vehicles/%s/skins/%d", driver.Car, driver.Number)
	outputPath := fmt.Sprintf("./_generated_outputs/%s", driver.Car)

	skinFiles := []string {
		"skin.dds",
		"skin_region.dds",
		"skin.png",
		"skin.json",
	}
	optionalSkinFiles := []string {
		"skin_helmet.dds",
		"skin_driver.dds",
	}

	carConfig, configErr := config.FindCar(driver.Car)
	if(configErr == nil) {
		for _, postfix := range carConfig.Postfixes {
			optionalSkinFiles = append(optionalSkinFiles, fmt.Sprintf("skin%s.dds", postfix))
		}
	}
	
	for _, skinFile := range skinFiles {
		newSkinFilePrefix := strings.Replace(skinFile, "skin", driver.GetLiveryIdentifier(), 1)
		inputFile := fmt.Sprintf("%s/%s", skinPath, skinFile)
		outputFile := fmt.Sprintf("%s/%s", outputPath, newSkinFilePrefix)
		_,copyErr := tools.CopyFile(inputFile, outputFile)
		if(copyErr != nil) {
			panic(copyErr)
		}
	}
	for _, skinFile := range optionalSkinFiles {
		newSkinFilePrefix := strings.Replace(skinFile, "skin", driver.GetLiveryIdentifier(), 1)
		inputFile := fmt.Sprintf("%s/%s", skinPath, skinFile)
		outputFile := fmt.Sprintf("%s/%s", outputPath, newSkinFilePrefix)
		_, copyErr := tools.CopyFile(inputFile, outputFile)
		if(copyErr != nil) {
			fmt.Printf("Error copying optional file %s for %s #%d\n", skinFile, driver.Car, driver.Number)
		}
	}
	
	vehTemplatePath := fmt.Sprintf("./templates/vehicles/%s/_vehicle.veh", driver.Car)
	vehOutputPath := fmt.Sprintf("%s/%s.veh", outputPath, driver.GetLiveryIdentifier())
	vehErr := fillInVehFile(vehTemplatePath, vehOutputPath, driver)
	if(vehErr != nil) {
		panic(vehErr)
	}
}

func copyCarFilesToDir(car config.CarConfigModel, config config.ApplicationConfiguration) {
	skinPath := fmt.Sprintf("./templates/vehicles/%s", car.Id)
	outputPath := fmt.Sprintf("./_generated_outputs/%s", car.Id)

	files := []string {
		"LEAGUE_MOD_Upgrades.ini",
		"brand_logo.png",
	}

	for _, skinFile := range files {
		inputFile := fmt.Sprintf("%s/%s", skinPath, skinFile)
		outputFile := fmt.Sprintf("%s/%s", outputPath, skinFile)
		_,copyErr := tools.CopyFile(inputFile, outputFile)
		if(copyErr != nil) {
			panic(copyErr)
		}
	}
}