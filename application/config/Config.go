package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

//This holds all the other configurations.
type Config struct {
	AppConfig ApplicationConfig
}

//Application config holds all the app related configurations
type ApplicationConfig struct {
	Name 		 string 	`yaml:"name"`
	Host		 string 	`yaml:"host"`
	Port 		 int    	`yaml:"port"`
	AuthToken    string 	`yaml:"auth_token"`
	SecretKey    string 	`yaml:"secret_key"`
}

//ParseConfig function makes all the config as single object
func ParseConfig(path string) *Config {
	//set config directory
	dir := getConfigurationPath(path)

	return &Config{
		AppConfig: parseApplicationConfig(dir),
	}
}

func getConfigurationPath(path string) string {
	//to get last character of the path
	char := path[len(path)-1]
	if os.IsPathSeparator(char) {
		return path
	}

	return path + string(os.PathSeparator)
}

func parseApplicationConfig(path string) ApplicationConfig {
	con := ApplicationConfig{}
	getConfigStruct(path+"app.yaml", &con)
	return con
}

func getConfigStruct(file string, unpacker interface{}) {
	content := readFile(file)
	err := yaml.Unmarshal(content, unpacker)
	if err != nil {
		panic(fmt.Sprintf("error: %v", err))
	}
}

func readFile(file string) []byte {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		panic(fmt.Sprintf("error: %v", err))
	}

	return content
}
