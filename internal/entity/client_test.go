package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewClient(t *testing.T) {
	client, err := NewClient("John Doe", "john@doe.com")

	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "John Doe", client.Name)
	assert.Equal(t, "john@doe.com", client.Email)
}

func TestCreateNewClientWhenNameIsInvalid(t *testing.T) {
	client, err := NewClient("", "john@doe.com")

	assert.NotNil(t, err)
	assert.Error(t, err, "name is required")
	assert.Nil(t, client)
}

func TestCreateNewClientWhenEmailIsInvalid(t *testing.T) {
	client, err := NewClient("John Doe", "")

	assert.NotNil(t, err)
	assert.Error(t, err, "email is required")
	assert.Nil(t, client)
}

func TestUpdateClient(t *testing.T) {
	client, _ := NewClient("John Doe", "john@doe.com")
	err := client.Update("John Doe Update", "j@doe.com")

	assert.Nil(t, err)
	assert.Equal(t, "John Doe Update", client.Name)
	assert.Equal(t, "j@doe.com", client.Email)
}

func TestUpdateClientWithInvalidArgs(t *testing.T) {
	client, _ := NewClient("John Doe", "john@doe.com")
	err := client.Update("", "j@doe.com")

	assert.NotNil(t, err)
	assert.Error(t, err, "name is required")
}
