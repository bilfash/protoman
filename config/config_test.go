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
		err := os.Setenv("APP_NAME", "protoman")

		appName := config.AppName()

		assert.NoError(t, err)
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
	failTest(t, err)
	_, err = f.WriteString(yamlString)
	failTest(t, err)
}

func removeEnvYamlFile() {
	_ = os.Remove(yamlFile)
}

func failTest(t *testing.T, err error) {
	if err != nil {
		t.Failed()
	}
}