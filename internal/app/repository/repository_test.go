package repository

import (
	"movierental/configs"
	"movierental/database"
	"testing"
)

func TestShouldReturnHelloworldMessage(t *testing.T) {
	config := configs.Config{}
	configs.GetConfigs(&config)
	db := database.CreateConnection(config.Database)
	repository := NewRepository(db)
	message, err := repository.GetEndPoint()
	expectedMessage := "helloworld..."
	var expectedError error
	if message != expectedMessage || err != expectedError {
		t.Errorf("Expected %s and %s , got %s and %s", expectedMessage, expectedError, message, err)
	}
}
