package database

import (
	"database/sql"
	"testing"

	"github.com/caiofsr/walletcore/internal/entity"
	"github.com/stretchr/testify/suite"
)

type TransferDBTestSuite struct {
	suite.Suite
	db            *sql.DB
	client        *entity.Client
	client2       *entity.Client
	accountFrom   *entity.Account
	accountTo     *entity.Account
	transactionDB *TransferDB
}

func (s *TransferDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("Create table clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	db.Exec("Create table accounts (id varchar(255), client_id varchar(255), balance int, created_at date)")
	db.Exec("Create table transfers (id varchar(255), account_id_from varchar(255), account_id_to varchar(255), amount int, created_at date)")
	client, err := entity.NewClient("John", "j@j.com")
	s.Nil(err)
	s.client = client
	client2, err := entity.NewClient("John2", "jj@j.com")
	s.Nil(err)
	s.client2 = client2
	//creating accounts
	accountFrom, _ := entity.NewAccount(s.client)
	accountFrom.Balance = 1000
	s.accountFrom = accountFrom
	accountTo, _ := entity.NewAccount(s.client2)
	accountTo.Balance = 1000
	s.accountTo = accountTo
	s.transactionDB = NewTransferDB(db)
}

func (s *TransferDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients")
	s.db.Exec("DROP TABLE accounts")
	s.db.Exec("DROP TABLE transfers")
}

func TestTransactionDBTestSuite(t *testing.T) {
	suite.Run(t, new(TransferDBTestSuite))
}

func (s *TransferDBTestSuite) TestCreate() {
	transaction, err := entity.NewTransfer(s.accountFrom, s.accountTo, 100)
	s.Nil(err)
	err = s.transactionDB.Create(transaction)
	s.Nil(err)
}
