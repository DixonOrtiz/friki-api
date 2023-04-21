package gormdb

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetConfigWithError(t *testing.T) {
	os.Setenv(DB_PORT, "not_a_port")
	defer os.Clearenv()

	config, err := GetConfig()

	assert.Equal(t, config, Config{})
	assert.EqualError(t, err, "the string 'not_a_port' is not a number")
}

func TestGetConfig(t *testing.T) {
	os.Setenv(DB_USER, testDBUser)
	os.Setenv(DB_PASSWORD, testDBPassword)
	os.Setenv(DB_HOST, testDBHost)
	os.Setenv(DB_NAME, testDBName)
	os.Setenv(DB_PORT, testDBPort)
	defer os.Clearenv()

	config, err := GetConfig()

	assert.Equal(t, config, expectedConfig)
	assert.Nil(t, err)
}
