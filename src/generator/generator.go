package generator

import(
	"fmt"
	"os"
	"../config"
)

func generateDirForCar(id string) error {
	return os.MkdirAll(fmt.Sprintf("./_generated_outputs/%s", id), os.ModeDir)
}

func RemoveOutputDir() error {
	return os.RemoveAll(fmt.Sprintf("./_generated_outputs"))
}

func Generate(conf config.ApplicationConfiguration, version string) {
	for _, car := range conf.CarConfig {
		err := generateDirForCar(car.Id)
		if(err != nil) {
			fmt.Printf("ERROR CREATE_DIR FOR CAR %s \n", car.Id)
		}
	}

	for _, driver := range conf.LeagueConfig.Driver {
		copyDriverFilesToDir(driver, conf)
	}

	for _, car := range conf.CarConfig {
		copyCarFilesToDir(car, conf)
		generateRF2Files(car.Id, version)
	}
}
