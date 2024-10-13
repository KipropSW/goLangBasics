package domain

import (
	"github.com/jmoiron/sqlx"
	"server/errs"
	"server/logger"
	"strconv"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (d AccountRepositoryDb) SaveTransaction(t Transaction) (*Transaction, *errs.AppError) {
	tx, err := d.client.Begin()
	if err != nil {
		logger.Error("Error while starting a new Transaction for bank account transaction")
		return nil, errs.NewUnexpectedError("Unexpected DB error")
	}

	sqlInsert := "INSERT INTO `transactions`(account_id, transaction_type, amount, transaction_date) VALUES (?,?,?,?)"
	result, err := tx.Exec(sqlInsert, t.AccountId, t.TransactionType, t.Amount, t.TransactionDate)
	if err != nil {
		logger.Error("Error while inserting a new record" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}
	if t.IsWithdrawal() {
		_, err = tx.Exec(`UPDATE accounts SET amount = amount - ? WHERE account_id = ?`, t.Amount, t.AccountId)
	} else {
		_, err = tx.Exec(`UPDATE accounts SET amount = amount + ? WHERE account_id = ?`, t.Amount, t.AccountId)
	}

	if err != nil {
		tx.Rollback()
		logger.Error("Error while saving transaction" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}
	err = tx.Commit()
	if err != nil {
		logger.Error("Error while commiting transaction" + err.Error())
	}
	transactionId, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last inserted ID" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}
	account, appError := d.FindBy(t.AccountId)
	if appError != nil {
		return nil, appError
	}
	t.TransactionId = strconv.FormatInt(transactionId, 10)
	t.Amount = account.Amount
	return &t, nil
}

func (d AccountRepositoryDb) Save(a Account) (*Account, *errs.AppError) {
	sqlInsert := "INSERT INTO accounts (listing_id, opening_date, account_type, amount, status) values (?, ?, ?, ?, ?)"
	result, err := d.client.Exec(sqlInsert, a.ListingId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("Error while creating new account" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}
	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last inserted id" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}
	a.AccountId = strconv.FormatUint(uint64(id), 10)
	return &a, nil
}

func (d AccountRepositoryDb) FindBy(accountId string) (*Account, *errs.AppError) {
	sqlGetAccount := "SELECT account_id, customer_id, opening_date, account_type, amount from accounts where account_id = ?"
	var account Account
	err := d.client.Get(&account, sqlGetAccount, accountId)
	if err != nil {
		logger.Error("Error while fetching account information: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	return &account, nil
}

func NewAccountRepositoryDb(db *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{client: db}
}
