// Package models contains the types for schema 'ccdb_dupl'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"
)

// Payment represents a row from 'ccdb_dupl.payment'.
type Payment struct {
	ID                                    uint                `json:"_id"`                                         // _id
	SubscriptionID                        NullInt64           `json:"subscription_id"`                             // subscription_id
	TransactionDate                       NullTime            `json:"transaction_date"`                            // transaction_date
	NatureOfTransaction                   NatureOfTransaction `json:"nature_of_transaction"`                       // nature_of_transaction
	ContractStatus                        ContractStatus      `json:"contract_status"`                             // contract_status
	CashRecieved                          NullString          `json:"cash_recieved"`                               // cash_recieved
	Currency                              string              `json:"currency"`                                    // currency
	FxRateToHkd                           float32             `json:"fx_rate_to_HKD"`                              // fx_rate_to_HKD
	CashRecievedHkd                       NullFloat64         `json:"cash_recieved_HKD"`                           // cash_recieved_HKD
	PaymentMode                           PaymentMode         `json:"payment_mode"`                                // payment_mode
	FirstMonthOfAccural                   NullTime            `json:"first_month_of_accural"`                      // first_month_of_accural
	MonthsOfAccural                       NullInt64           `json:"months_of_accural"`                           // months_of_accural
	LastMonthOfAccural                    NullTime            `json:"last_month_of_accural"`                       // last_month_of_accural
	MonthlyRevenueAllocation              NullFloat64         `json:"monthly_revenue_allocation"`                  // monthly_revenue_allocation
	CashReceivedForPaymentsToThirdParties NullFloat64         `json:"cash_received_for_payments_to_third_parties"` // cash_received_for_payments_to_third_parties
	TypeOfThirdParties                    NullString          `json:"type_of_third_parties"`                       // type_of_third_parties
	NameOfThirdParty                      NullString          `json:"name_of_third_party"`                         // name_of_third_party

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Payment exists in the database.
func (p *Payment) Exists() bool {
	return p._exists
}

// Deleted provides information if the Payment has been deleted from the database.
func (p *Payment) Deleted() bool {
	return p._deleted
}

// Insert inserts the Payment to the database.
func (p *Payment) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if p._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO ccdb_dupl.payment (` +
		`subscription_id, transaction_date, nature_of_transaction, contract_status, cash_recieved, currency, fx_rate_to_HKD, cash_recieved_HKD, payment_mode, first_month_of_accural, months_of_accural, last_month_of_accural, monthly_revenue_allocation, cash_received_for_payments_to_third_parties, type_of_third_parties, name_of_third_party` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, p.SubscriptionID, p.TransactionDate, p.NatureOfTransaction, p.ContractStatus, p.CashRecieved, p.Currency, p.FxRateToHkd, p.CashRecievedHkd, p.PaymentMode, p.FirstMonthOfAccural, p.MonthsOfAccural, p.LastMonthOfAccural, p.MonthlyRevenueAllocation, p.CashReceivedForPaymentsToThirdParties, p.TypeOfThirdParties, p.NameOfThirdParty)
	res, err := db.Exec(sqlstr, p.SubscriptionID, p.TransactionDate, p.NatureOfTransaction, p.ContractStatus, p.CashRecieved, p.Currency, p.FxRateToHkd, p.CashRecievedHkd, p.PaymentMode, p.FirstMonthOfAccural, p.MonthsOfAccural, p.LastMonthOfAccural, p.MonthlyRevenueAllocation, p.CashReceivedForPaymentsToThirdParties, p.TypeOfThirdParties, p.NameOfThirdParty)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	p.ID = uint(id)
	p._exists = true

	return nil
}

// Update updates the Payment in the database.
func (p *Payment) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	// if !p._exists {
	// 	return errors.New("update failed: does not exist")
	// }

	// // if deleted, bail
	// if p._deleted {
	// 	return errors.New("update failed: marked for deletion")
	// }

	// sql query
	const sqlstr = `UPDATE ccdb_dupl.payment SET ` +
		`subscription_id = ?, transaction_date = ?, nature_of_transaction = ?, contract_status = ?, cash_recieved = ?, currency = ?, fx_rate_to_HKD = ?, cash_recieved_HKD = ?, payment_mode = ?, first_month_of_accural = ?, months_of_accural = ?, last_month_of_accural = ?, monthly_revenue_allocation = ?, cash_received_for_payments_to_third_parties = ?, type_of_third_parties = ?, name_of_third_party = ?` +
		` WHERE _id = ?`

	// run query
	XOLog(sqlstr, p.SubscriptionID, p.TransactionDate, p.NatureOfTransaction, p.ContractStatus, p.CashRecieved, p.Currency, p.FxRateToHkd, p.CashRecievedHkd, p.PaymentMode, p.FirstMonthOfAccural, p.MonthsOfAccural, p.LastMonthOfAccural, p.MonthlyRevenueAllocation, p.CashReceivedForPaymentsToThirdParties, p.TypeOfThirdParties, p.NameOfThirdParty, p.ID)
	_, err = db.Exec(sqlstr, p.SubscriptionID, p.TransactionDate, p.NatureOfTransaction, p.ContractStatus, p.CashRecieved, p.Currency, p.FxRateToHkd, p.CashRecievedHkd, p.PaymentMode, p.FirstMonthOfAccural, p.MonthsOfAccural, p.LastMonthOfAccural, p.MonthlyRevenueAllocation, p.CashReceivedForPaymentsToThirdParties, p.TypeOfThirdParties, p.NameOfThirdParty, p.ID)
	return err
}

// Save saves the Payment to the database.
func (p *Payment) Save(db XODB) error {
	if p.Exists() {
		return p.Update(db)
	}

	return p.Insert(db)
}

// Delete deletes the Payment from the database.
func (p *Payment) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	// if !p._exists {
	// 	return nil
	// }

	// // if deleted, bail
	// if p._deleted {
	// 	return nil
	// }

	// sql query
	const sqlstr = `DELETE FROM ccdb_dupl.payment WHERE _id = ?`

	// run query
	XOLog(sqlstr, p.ID)
	_, err = db.Exec(sqlstr, p.ID)
	if err != nil {
		return err
	}

	// set deleted
	p._deleted = true

	return nil
}

// Subscription returns the Subscription associated with the Payment's SubscriptionID (subscription_id).
//
// Generated from foreign key 'payment_ibfk_1'.
func (p *Payment) Subscription(db XODB) (*Subscription, error) {
	s := &Subscription{}
	err := s.ByID(db, uint(p.SubscriptionID.Int64))
	return s, err
}

// PaymentByID retrieves a row from 'ccdb_dupl.payment' as a Payment.
//
// Generated from index 'payment__id_pkey'.
func (p *Payment) ByID(db XODB, id uint) error {
	// sql query
	const sqlstr = `SELECT ` +
		`_id, subscription_id, transaction_date, nature_of_transaction, contract_status, cash_recieved, currency, fx_rate_to_HKD, cash_recieved_HKD, payment_mode, first_month_of_accural, months_of_accural, last_month_of_accural, monthly_revenue_allocation, cash_received_for_payments_to_third_parties, type_of_third_parties, name_of_third_party ` +
		`FROM ccdb_dupl.payment ` +
		`WHERE _id = ?`

	// run query
	XOLog(sqlstr, id)
	p._exists = true

	return db.QueryRow(sqlstr, id).Scan(&p.ID, &p.SubscriptionID, &p.TransactionDate, &p.NatureOfTransaction, &p.ContractStatus, &p.CashRecieved, &p.Currency, &p.FxRateToHkd, &p.CashRecievedHkd, &p.PaymentMode, &p.FirstMonthOfAccural, &p.MonthsOfAccural, &p.LastMonthOfAccural, &p.MonthlyRevenueAllocation, &p.CashReceivedForPaymentsToThirdParties, &p.TypeOfThirdParties, &p.NameOfThirdParty)
}

// PaymentsBySubscriptionID retrieves a row from 'ccdb_dupl.payment' as a Payment.
//
// Generated from index 'subscription_id'.
func PaymentsBySubscriptionID(db XODB, subscriptionID sql.NullInt64) ([]*Payment, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`_id, subscription_id, transaction_date, nature_of_transaction, contract_status, cash_recieved, currency, fx_rate_to_HKD, cash_recieved_HKD, payment_mode, first_month_of_accural, months_of_accural, last_month_of_accural, monthly_revenue_allocation, cash_received_for_payments_to_third_parties, type_of_third_parties, name_of_third_party ` +
		`FROM ccdb_dupl.payment ` +
		`WHERE subscription_id = ?`

	// run query
	XOLog(sqlstr, subscriptionID)
	q, err := db.Query(sqlstr, subscriptionID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Payment{}
	for q.Next() {
		p := Payment{
			_exists: true,
		}

		// scan
		err = q.Scan(&p.ID, &p.SubscriptionID, &p.TransactionDate, &p.NatureOfTransaction, &p.ContractStatus, &p.CashRecieved, &p.Currency, &p.FxRateToHkd, &p.CashRecievedHkd, &p.PaymentMode, &p.FirstMonthOfAccural, &p.MonthsOfAccural, &p.LastMonthOfAccural, &p.MonthlyRevenueAllocation, &p.CashReceivedForPaymentsToThirdParties, &p.TypeOfThirdParties, &p.NameOfThirdParty)
		if err != nil {
			return nil, err
		}

		res = append(res, &p)
	}

	return res, nil
}

func (p *Payment) All(db XODB) (interface{}, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`* ` +
		`FROM ccdb_dupl.payment `

	// run query
	XOLog(sqlstr)
	q, err := db.Query(sqlstr)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Payment{}
	for q.Next() {
		p := Payment{
			_exists: true,
		}

		// scan
		err = q.Scan(&p.ID, &p.SubscriptionID, &p.TransactionDate, &p.NatureOfTransaction, &p.ContractStatus, &p.CashRecieved, &p.Currency, &p.FxRateToHkd, &p.CashRecievedHkd, &p.PaymentMode, &p.FirstMonthOfAccural, &p.MonthsOfAccural, &p.LastMonthOfAccural, &p.MonthlyRevenueAllocation, &p.CashReceivedForPaymentsToThirdParties, &p.TypeOfThirdParties, &p.NameOfThirdParty)
		if err != nil {
			return nil, err
		}

		res = append(res, &p)
	}

	return res, nil
}
