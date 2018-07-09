package configuration


import (
	"os"
	"encoding/json"
)

// Configuration es la estructura que encarga de obtener la copia de los datos del archivo de configuración config.json
type Configuration struct {
	Engine   string
	Server   string
	Port     string
	User     string
	Password string
	Database string
}
// GetConfiguration se encarga de realizar la copia de los datos de configuración del archivo config.json en la estructura Configuration
func GetConfiguration() (Configuration, error){
	var config Configuration

	file, err := os.Open("./config.json")

	if err != nil {
		return config, err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&config)

	if err != nil {
		return config, err
	}
	return config, nil
}
