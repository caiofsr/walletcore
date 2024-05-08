package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewTransfer(t *testing.T) {
	client1, _ := NewClient("John Doe", "john@doe.com")
	client2, _ := NewClient("Mary Doe", "mary@doe.com")
	account1, _ := NewAccount(client1)
	account2, _ := NewAccount(client2)

	account1.Credit(1000)
	account2.Credit(1000)

	transfer, err := NewTransfer(account1, account2, 100)

	assert.Nil(t, err)
	assert.NotNil(t, transfer)
	assert.Equal(t, float64(1100), account2.Balance)
	assert.Equal(t, float64(900), account1.Balance)
}

func TestCreateNewTransferWhenAmountIsBelowZero(t *testing.T) {
	client1, _ := NewClient("John Doe", "john@doe.com")
	client2, _ := NewClient("Mary Doe", "mary@doe.com")
	account1, _ := NewAccount(client1)
	account2, _ := NewAccount(client2)

	account1.Credit(1000)
	account2.Credit(1000)

	transfer, err := NewTransfer(account1, account2, 0)

	assert.NotNil(t, err)
	assert.EqualError(t, err, "amount must be greater than zero")
	assert.Nil(t, transfer)
	assert.Equal(t, float64(1000), account1.Balance)
	assert.Equal(t, float64(1000), account2.Balance)
}

func TestCreateNewTransferIfBalanceIsLessThanTheAmount(t *testing.T) {
	client1, _ := NewClient("John Doe", "john@doe.com")
	client2, _ := NewClient("Mary Doe", "mary@doe.com")
	account1, _ := NewAccount(client1)
	account2, _ := NewAccount(client2)

	transfer, err := NewTransfer(account1, account2, 100)

	assert.NotNil(t, err)
	assert.EqualError(t, err, "insufficient fund")
	assert.Nil(t, transfer)
	assert.Equal(t, float64(0), account1.Balance)
	assert.Equal(t, float64(0), account2.Balance)
}

func TestCreateNewTransferWhenSameAccountFromTo(t *testing.T) {
	client, _ := NewClient("John Doe", "john@doe.com")
	account, _ := NewAccount(client)

	account.Credit(200)

	transfer, err := NewTransfer(account, account, 100)

	assert.NotNil(t, err)
	assert.EqualError(t, err, "is not possible to transfer to same account")
	assert.Nil(t, transfer)
	assert.Equal(t, float64(200), account.Balance)
}
