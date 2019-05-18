package config_test

import (
	"github.com/bilfash/protoman/config"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var yamlFile = "../application.yml"

func TestConfigAppName(t *testing.T){
	t.Run("os_environment", func(t *testing.T) {
		failTestIfError(t, os.Setenv("APP_NAME", "protoman"))

		appName := config.AppName()

		assert.Equal(t, "protoman", appName)
	})
	t.Run("yaml_file", func(t *testing.T) {
		yamlString := "APP_NAME: protoman"
		setupEnvYamlFile(t, yamlString)
		defer removeEnvYamlFile()

		appName := config.AppName()

		assert.Equal(t, "protoman", appName)
	})
}

func setupEnvYamlFile(t *testing.T, yamlString string) {
	os.Clearenv()
	f, err := os.Create(yamlFile)
	failTestIfError(t, err)
	_, err = f.WriteString(yamlString)
	failTestIfError(t, err)
}

func removeEnvYamlFile() {
	_ = os.Remove(yamlFile)
}

func failTestIfError(t *testing.T, err error) {
	if err != nil {
		t.Failed()
	}
}