package database

import (
	"database/sql"
	"testing"

	"github.com/marcelofelixsalgado/fullcycle-eda/balances/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type BalanceDBTestSuite struct {
	suite.Suite
	db        *sql.DB
	balanceDB *BalanceDB
}

func (s *BalanceDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("Create table balances (id varchar(255), account_id varchar(255), amount float, updated_at date)")
	s.balanceDB = NewBalanceDB(db)
}

func (s *BalanceDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE balances")
}

func TestBalanceDBTestSuite(t *testing.T) {
	suite.Run(t, new(BalanceDBTestSuite))
}

func (s *BalanceDBTestSuite) TestUpdate() {
	balance := &entity.Balance{
		AccountId: "11",
		Amount:    10,
	}
	err := s.balanceDB.Update(balance.AccountId, balance.Amount)
	s.Nil(err)
}

func (s *BalanceDBTestSuite) TestGet() {
	balance, _ := entity.NewBalance("11", 1000)
	s.balanceDB.Insert(balance)

	balanceDB, err := s.balanceDB.Get(balance.AccountId)
	s.Nil(err)
	s.Equal(balance.ID, balanceDB.ID)
	s.Equal(balance.AccountId, balanceDB.AccountId)
	s.Equal(balance.Amount, balanceDB.Amount)
}
