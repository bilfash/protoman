package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)


var (
	configFile = "../application.yml"
	conf map[string]interface{}
)

func AppName() string {
	return getMandatoryConfigString("APP_NAME")
}

func getMandatoryConfigString(key string) string {
	appName := os.Getenv(key)
	if appName == "" {
		initiateYamlFile()
		if _, ok := conf[key]; !ok {
			log.Fatalf("%s config required", key)
		}
		appName = conf[key].(string)
	}
	return appName
}

func initiateYamlFile() {
	if len(conf) > 0 {
		return
	}
	yamlFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatalf("%s.Get err   #%v ", configFile, err)
	}
	fmt.Println(yamlFile)
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		log.Fatalf("%s Unmarshal: %v", configFile, err)
	}
}