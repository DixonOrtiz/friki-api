package authinfra

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetupConfig(t *testing.T) {
	os.Setenv(GOOGLE_CLIENT_ID, testGoogleClientID)
	os.Setenv(GOOGLE_CLIENT_SECRET, testGoogleClientSecret)
	os.Setenv(REDIRECT_URL, testRedirectUrl)
	defer os.Clearenv()

	config := SetupConfig()

	assert.Equal(t, config, expectedConfig)
}
