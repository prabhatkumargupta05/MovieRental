package repository

import (
	"movierental/configs"
	"testing"
)

func TestShouldReturnHelloworldMessage(t *testing.T) {
	config := configs.Config{}
	configs.GetConfigs(&config)
	repository := NewRepository(nil)
	message, err := repository.GetEndPoint()
	expectedMessage := "helloworld..."
	var expectedError error
	if message != expectedMessage || err != expectedError {
		t.Errorf("Expected %s and %s , got %s and %s", expectedMessage, expectedError, message, err)
	}
}
