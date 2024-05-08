package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewAccount(t *testing.T) {
	client, _ := NewClient("John Doe", "john@doe.com")
	account, err := NewAccount(client)

	assert.Nil(t, err)
	assert.Equal(t, client.ID, account.Client.ID)
}

func TestCreateNewAccountWithNilClient(t *testing.T) {
	account, err := NewAccount(nil)

	assert.NotNil(t, err)
	assert.EqualError(t, err, "client is required")
	assert.Nil(t, account)
}

func TestCreditAccount(t *testing.T) {
	client, _ := NewClient("John Doe", "john@doe.com")
	account, _ := NewAccount(client)

	account.Credit(100)

	assert.Equal(t, float64(100), account.Balance)
}

func TestDebitAccount(t *testing.T) {
	client, _ := NewClient("John Doe", "john@doe.com")
	account, _ := NewAccount(client)

	account.Credit(150)
	account.Debit(100)

	assert.Equal(t, float64(50), account.Balance)
}
