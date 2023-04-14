package database

import (
	"database/sql"
	"time"

	"github.com.br/marcelofelixsalgado/fullcycle-eda/balances/internal/entity"
)

type BalanceDB struct {
	DB *sql.DB
}

func NewBalanceDB(db *sql.DB) *BalanceDB {
	return &BalanceDB{
		DB: db,
	}
}

func (a *BalanceDB) Insert(balance *entity.Balance) error {
	stmt, err := a.DB.Prepare("INSERT INTO balances (id, account_id, amount, updated_at) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(balance.ID, balance.AccountId, balance.Amount, balance.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (c *BalanceDB) Get(account_id string) (*entity.Balance, error) {
	balance := &entity.Balance{}
	stmt, err := c.DB.Prepare("SELECT id, account_id, amount, updated_at FROM balances WHERE account_id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(account_id)
	if err := row.Scan(&balance.ID, &balance.AccountId, &balance.Amount, &balance.UpdatedAt); err != nil {
		return nil, err
	}
	return balance, nil
}

func (c *BalanceDB) Update(accountId string, amount float64) error {

	stmt, err := c.DB.Prepare("UPDATE balances SET amount = ?, updated_at = ? where account_id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(amount, time.Now(), accountId)
	if err != nil {
		return err
	}

	return nil
}
