package database

import (
	"database/sql"

	"github.com/caiofsr/walletcore/internal/entity"
)

type TransferDB struct {
	DB *sql.DB
}

func NewTransferDB(db *sql.DB) *TransferDB {
	return &TransferDB{
		DB: db,
	}
}

func (t *TransferDB) Create(transfer *entity.Transfer) error {
	stmt, err := t.DB.Prepare("INSERT INTO transfers (id, account_id_from, account_id_to, amount, created_at) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(transfer.ID, transfer.AccountFrom.ID, transfer.AccountTo.ID, transfer.Amount, transfer.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}
