package generator

import (
	"io/ioutil"
	"strings"
	"../config"
	"fmt"
)

func fillInVehFile(templatePath string, outputPath string, driver config.DriverModel) error {
	bytes, readError := ioutil.ReadFile(templatePath)
	if(readError != nil) {
		return readError
	}
	
	content := string(bytes)

	content = strings.ReplaceAll(content, "{{LIVERY}}", driver.GetLiveryIdentifier())
	content = strings.ReplaceAll(content, "{{NUMBER}}", fmt.Sprintf("%d", driver.Number))
	content = strings.ReplaceAll(content, "{{DRIVER}}", driver.Name)
	content = strings.ReplaceAll(content, "{{TEAMNAME}}", driver.Team)
	writeError := ioutil.WriteFile(outputPath, []byte(content), 0644)
	if(writeError != nil) {
		return writeError
	}
	return nil
}