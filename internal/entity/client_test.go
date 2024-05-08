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
	assert.EqualError(t, err, "name is required")
	assert.Nil(t, client)
}

func TestCreateNewClientWhenEmailIsInvalid(t *testing.T) {
	client, err := NewClient("John Doe", "")

	assert.NotNil(t, err)
	assert.EqualError(t, err, "email is required")
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
	assert.EqualError(t, err, "name is required")
}

func TestAddAccountToClient(t *testing.T) {
	client, _ := NewClient("John Doe", "john@doe.com")
	account, _ := NewAccount(client)

	err := client.AddAccount(account)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(client.Accounts))
}

func TestAddAccountToAnotherClient(t *testing.T) {
	client, _ := NewClient("John Doe", "john@doe.com")
	client2, _ := NewClient("Mary Doe", "mary@doe.com")
	account, _ := NewAccount(client)

	err := client2.AddAccount(account)

	assert.NotNil(t, err)
	assert.EqualError(t, err, "account does not belong to client")
	assert.Equal(t, 0, len(client2.Accounts))
}
