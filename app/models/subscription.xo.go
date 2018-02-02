// Package models contains the types for schema 'ccdb_dupl'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/go-sql-driver/mysql"
)

// Subscription represents a row from 'ccdb_dupl.subscription'.
type Subscription struct {
	ID                                   uint            `json:"_id"`                                       // _id
	ClientID                             sql.NullInt64   `json:"client_id"`                                 // client_id
	PackageID                            uint            `json:"package_id"`                                // package_id
	PartnerID                            sql.NullInt64   `json:"partner_id"`                                // partner_id
	StartDate                            mysql.NullTime  `json:"start_date"`                                // start_date
	EndDate                              mysql.NullTime  `json:"end_date"`                                  // end_date
	MonthsFree                           sql.NullInt64   `json:"months_free"`                               // months_free
	Status                               sql.NullBool    `json:"status"`                                    // status
	PaymentTerms                         PaymentTerm     `json:"payment_terms"`                             // payment_terms
	ContractStatus                       ContractStatus  `json:"contract_status"`                           // contract_status
	ContractStartDate                    mysql.NullTime  `json:"contract_start_date"`                       // contract_start_date
	ContractEndDate                      mysql.NullTime  `json:"contract_end_date"`                         // contract_end_date
	MonthlyPayment                       sql.NullFloat64 `json:"monthly_payment"`                           // monthly_payment
	AnnualPayment                        sql.NullFloat64 `json:"annual_payment"`                            // annual_payment
	ManagedAccount                       sql.NullBool    `json:"managed_account"`                           // managed_account
	ChangeOfPlanAtRenewal                sql.NullString  `json:"change_of_plan_at_renewal"`                 // change_of_plan_at_renewal
	ChurnRationale                       sql.NullString  `json:"churn_rationale"`                           // churn_rationale
	MonthsToChurn                        sql.NullInt64   `json:"months_to_churn"`                           // months_to_churn
	Reconcillation                       sql.NullString  `json:"reconcillation"`                            // reconcillation
	SubscriptionAddonPrice               sql.NullFloat64 `json:"subscription_addon_price"`                  // subscription_addon_price
	ServicePaidInInstallments            sql.NullString  `json:"service_paid_in_installments"`              // service_paid_in_installments
	SetupFee                             sql.NullFloat64 `json:"setup_fee"`                                 // setup_fee
	NetServicePriceInc                   sql.NullFloat64 `json:"net_service_price_INC"`                     // net_service_price_INC
	NetServicePriceTr                    sql.NullFloat64 `json:"net_service_price_TR"`                      // net_service_price_TR
	NetServicePriceOthers                sql.NullFloat64 `json:"net_service_price_others"`                  // net_service_price_others
	DiscountsOnSubsPrice                 sql.NullFloat64 `json:"discounts_on_subs_price"`                   // discounts_on_subs_price
	DiscountsOnSubsAddon                 sql.NullFloat64 `json:"discounts_on_subs_addon"`                   // discounts_on_subs_addon
	DiscountsOnServicePaidInInstallments sql.NullFloat64 `json:"discounts_on_service_paid_in_installments"` // discounts_on_service_paid_in_installments
	DiscountsOnServiceInc                sql.NullFloat64 `json:"discounts_on_service_INC"`                  // discounts_on_service_INC
	DiscountsOnServiceTr                 sql.NullFloat64 `json:"discounts_on_service_TR"`                   // discounts_on_service_TR
	DiscountsOnServiceOthers             sql.NullFloat64 `json:"discounts_on_service_others"`               // discounts_on_service_others
	DiscountsOnSetupFee                  sql.NullFloat64 `json:"discounts_on_setup_fee"`                    // discounts_on_setup_fee
	OtherPromotionsWithOutlay            sql.NullString  `json:"other_promotions_with_outlay"`              // other_promotions_with_outlay
	MonthsOfService                      sql.NullInt64   `json:"months_of_service"`                         // months_of_service
	SpecialPromotion                     sql.NullString  `json:"special_promotion"`                         // special_promotion
	NewSalesCommitment                   sql.NullInt64   `json:"new_sales_commitment"`                      // new_sales_commitment

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Subscription exists in the database.
func (s *Subscription) Exists() bool {
	return s._exists
}

// Deleted provides information if the Subscription has been deleted from the database.
func (s *Subscription) Deleted() bool {
	return s._deleted
}

// Insert inserts the Subscription to the database.
func (s *Subscription) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if s._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO ccdb_dupl.subscription (` +
		`client_id, package_id, partner_id, start_date, end_date, months_free, status, payment_terms, contract_status, contract_start_date, contract_end_date, monthly_payment, annual_payment, managed_account, change_of_plan_at_renewal, churn_rationale, months_to_churn, reconcillation, subscription_addon_price, service_paid_in_installments, setup_fee, net_service_price_INC, net_service_price_TR, net_service_price_others, discounts_on_subs_price, discounts_on_subs_addon, discounts_on_service_paid_in_installments, discounts_on_service_INC, discounts_on_service_TR, discounts_on_service_others, discounts_on_setup_fee, other_promotions_with_outlay, months_of_service, special_promotion, new_sales_commitment` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, s.ClientID, s.PackageID, s.PartnerID, s.StartDate, s.EndDate, s.MonthsFree, s.Status, s.PaymentTerms, s.ContractStatus, s.ContractStartDate, s.ContractEndDate, s.MonthlyPayment, s.AnnualPayment, s.ManagedAccount, s.ChangeOfPlanAtRenewal, s.ChurnRationale, s.MonthsToChurn, s.Reconcillation, s.SubscriptionAddonPrice, s.ServicePaidInInstallments, s.SetupFee, s.NetServicePriceInc, s.NetServicePriceTr, s.NetServicePriceOthers, s.DiscountsOnSubsPrice, s.DiscountsOnSubsAddon, s.DiscountsOnServicePaidInInstallments, s.DiscountsOnServiceInc, s.DiscountsOnServiceTr, s.DiscountsOnServiceOthers, s.DiscountsOnSetupFee, s.OtherPromotionsWithOutlay, s.MonthsOfService, s.SpecialPromotion, s.NewSalesCommitment)
	res, err := db.Exec(sqlstr, s.ClientID, s.PackageID, s.PartnerID, s.StartDate, s.EndDate, s.MonthsFree, s.Status, s.PaymentTerms, s.ContractStatus, s.ContractStartDate, s.ContractEndDate, s.MonthlyPayment, s.AnnualPayment, s.ManagedAccount, s.ChangeOfPlanAtRenewal, s.ChurnRationale, s.MonthsToChurn, s.Reconcillation, s.SubscriptionAddonPrice, s.ServicePaidInInstallments, s.SetupFee, s.NetServicePriceInc, s.NetServicePriceTr, s.NetServicePriceOthers, s.DiscountsOnSubsPrice, s.DiscountsOnSubsAddon, s.DiscountsOnServicePaidInInstallments, s.DiscountsOnServiceInc, s.DiscountsOnServiceTr, s.DiscountsOnServiceOthers, s.DiscountsOnSetupFee, s.OtherPromotionsWithOutlay, s.MonthsOfService, s.SpecialPromotion, s.NewSalesCommitment)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	s.ID = uint(id)
	s._exists = true

	return nil
}

// Update updates the Subscription in the database.
func (s *Subscription) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !s._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if s._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE ccdb_dupl.subscription SET ` +
		`client_id = ?, package_id = ?, partner_id = ?, start_date = ?, end_date = ?, months_free = ?, status = ?, payment_terms = ?, contract_status = ?, contract_start_date = ?, contract_end_date = ?, monthly_payment = ?, annual_payment = ?, managed_account = ?, change_of_plan_at_renewal = ?, churn_rationale = ?, months_to_churn = ?, reconcillation = ?, subscription_addon_price = ?, service_paid_in_installments = ?, setup_fee = ?, net_service_price_INC = ?, net_service_price_TR = ?, net_service_price_others = ?, discounts_on_subs_price = ?, discounts_on_subs_addon = ?, discounts_on_service_paid_in_installments = ?, discounts_on_service_INC = ?, discounts_on_service_TR = ?, discounts_on_service_others = ?, discounts_on_setup_fee = ?, other_promotions_with_outlay = ?, months_of_service = ?, special_promotion = ?, new_sales_commitment = ?` +
		` WHERE _id = ?`

	// run query
	XOLog(sqlstr, s.ClientID, s.PackageID, s.PartnerID, s.StartDate, s.EndDate, s.MonthsFree, s.Status, s.PaymentTerms, s.ContractStatus, s.ContractStartDate, s.ContractEndDate, s.MonthlyPayment, s.AnnualPayment, s.ManagedAccount, s.ChangeOfPlanAtRenewal, s.ChurnRationale, s.MonthsToChurn, s.Reconcillation, s.SubscriptionAddonPrice, s.ServicePaidInInstallments, s.SetupFee, s.NetServicePriceInc, s.NetServicePriceTr, s.NetServicePriceOthers, s.DiscountsOnSubsPrice, s.DiscountsOnSubsAddon, s.DiscountsOnServicePaidInInstallments, s.DiscountsOnServiceInc, s.DiscountsOnServiceTr, s.DiscountsOnServiceOthers, s.DiscountsOnSetupFee, s.OtherPromotionsWithOutlay, s.MonthsOfService, s.SpecialPromotion, s.NewSalesCommitment, s.ID)
	_, err = db.Exec(sqlstr, s.ClientID, s.PackageID, s.PartnerID, s.StartDate, s.EndDate, s.MonthsFree, s.Status, s.PaymentTerms, s.ContractStatus, s.ContractStartDate, s.ContractEndDate, s.MonthlyPayment, s.AnnualPayment, s.ManagedAccount, s.ChangeOfPlanAtRenewal, s.ChurnRationale, s.MonthsToChurn, s.Reconcillation, s.SubscriptionAddonPrice, s.ServicePaidInInstallments, s.SetupFee, s.NetServicePriceInc, s.NetServicePriceTr, s.NetServicePriceOthers, s.DiscountsOnSubsPrice, s.DiscountsOnSubsAddon, s.DiscountsOnServicePaidInInstallments, s.DiscountsOnServiceInc, s.DiscountsOnServiceTr, s.DiscountsOnServiceOthers, s.DiscountsOnSetupFee, s.OtherPromotionsWithOutlay, s.MonthsOfService, s.SpecialPromotion, s.NewSalesCommitment, s.ID)
	return err
}

// Save saves the Subscription to the database.
func (s *Subscription) Save(db XODB) error {
	if s.Exists() {
		return s.Update(db)
	}

	return s.Insert(db)
}

// Delete deletes the Subscription from the database.
func (s *Subscription) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !s._exists {
		return nil
	}

	// if deleted, bail
	if s._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM ccdb_dupl.subscription WHERE _id = ?`

	// run query
	XOLog(sqlstr, s.ID)
	_, err = db.Exec(sqlstr, s.ID)
	if err != nil {
		return err
	}

	// set deleted
	s._deleted = true

	return nil
}

// Client returns the Client associated with the Subscription's ClientID (client_id).
//
// Generated from foreign key 'subscription_ibfk_1'.
// func (s *Subscription) Client(db XODB) (*Client, error) {
// 	return ClientByID(db, uint(s.ClientID.Int64))
// }

// Package returns the Package associated with the Subscription's PackageID (package_id).
//
// Generated from foreign key 'subscription_ibfk_2'.
// func (s *Subscription) Package(db XODB) (*Package, error) {
// 	return PackageByID(db, s.PackageID)
// }

// Partner returns the Partner associated with the Subscription's PartnerID (partner_id).
//
// Generated from foreign key 'subscription_ibfk_3'.
// func (s *Subscription) Partner(db XODB) (*Partner, error) {
// 	return PartnerByID(db, uint(s.PartnerID.Int64))
// }

// SubscriptionsByClientID retrieves a row from 'ccdb_dupl.subscription' as a Subscription.
//
// Generated from index 'client_id'.
func SubscriptionsByClientID(db XODB, clientID sql.NullInt64) ([]*Subscription, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`_id, client_id, package_id, partner_id, start_date, end_date, months_free, status, payment_terms, contract_status, contract_start_date, contract_end_date, monthly_payment, annual_payment, managed_account, change_of_plan_at_renewal, churn_rationale, months_to_churn, reconcillation, subscription_addon_price, service_paid_in_installments, setup_fee, net_service_price_INC, net_service_price_TR, net_service_price_others, discounts_on_subs_price, discounts_on_subs_addon, discounts_on_service_paid_in_installments, discounts_on_service_INC, discounts_on_service_TR, discounts_on_service_others, discounts_on_setup_fee, other_promotions_with_outlay, months_of_service, special_promotion, new_sales_commitment ` +
		`FROM ccdb_dupl.subscription ` +
		`WHERE client_id = ?`

	// run query
	XOLog(sqlstr, clientID)
	q, err := db.Query(sqlstr, clientID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Subscription{}
	for q.Next() {
		s := Subscription{
			_exists: true,
		}

		// scan
		err = q.Scan(&s.ID, &s.ClientID, &s.PackageID, &s.PartnerID, &s.StartDate, &s.EndDate, &s.MonthsFree, &s.Status, &s.PaymentTerms, &s.ContractStatus, &s.ContractStartDate, &s.ContractEndDate, &s.MonthlyPayment, &s.AnnualPayment, &s.ManagedAccount, &s.ChangeOfPlanAtRenewal, &s.ChurnRationale, &s.MonthsToChurn, &s.Reconcillation, &s.SubscriptionAddonPrice, &s.ServicePaidInInstallments, &s.SetupFee, &s.NetServicePriceInc, &s.NetServicePriceTr, &s.NetServicePriceOthers, &s.DiscountsOnSubsPrice, &s.DiscountsOnSubsAddon, &s.DiscountsOnServicePaidInInstallments, &s.DiscountsOnServiceInc, &s.DiscountsOnServiceTr, &s.DiscountsOnServiceOthers, &s.DiscountsOnSetupFee, &s.OtherPromotionsWithOutlay, &s.MonthsOfService, &s.SpecialPromotion, &s.NewSalesCommitment)
		if err != nil {
			return nil, err
		}

		res = append(res, &s)
	}

	return res, nil
}

// SubscriptionsByPackageID retrieves a row from 'ccdb_dupl.subscription' as a Subscription.
//
// Generated from index 'package_id'.
func SubscriptionsByPackageID(db XODB, packageID uint) ([]*Subscription, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`_id, client_id, package_id, partner_id, start_date, end_date, months_free, status, payment_terms, contract_status, contract_start_date, contract_end_date, monthly_payment, annual_payment, managed_account, change_of_plan_at_renewal, churn_rationale, months_to_churn, reconcillation, subscription_addon_price, service_paid_in_installments, setup_fee, net_service_price_INC, net_service_price_TR, net_service_price_others, discounts_on_subs_price, discounts_on_subs_addon, discounts_on_service_paid_in_installments, discounts_on_service_INC, discounts_on_service_TR, discounts_on_service_others, discounts_on_setup_fee, other_promotions_with_outlay, months_of_service, special_promotion, new_sales_commitment ` +
		`FROM ccdb_dupl.subscription ` +
		`WHERE package_id = ?`

	// run query
	XOLog(sqlstr, packageID)
	q, err := db.Query(sqlstr, packageID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Subscription{}
	for q.Next() {
		s := Subscription{
			_exists: true,
		}

		// scan
		err = q.Scan(&s.ID, &s.ClientID, &s.PackageID, &s.PartnerID, &s.StartDate, &s.EndDate, &s.MonthsFree, &s.Status, &s.PaymentTerms, &s.ContractStatus, &s.ContractStartDate, &s.ContractEndDate, &s.MonthlyPayment, &s.AnnualPayment, &s.ManagedAccount, &s.ChangeOfPlanAtRenewal, &s.ChurnRationale, &s.MonthsToChurn, &s.Reconcillation, &s.SubscriptionAddonPrice, &s.ServicePaidInInstallments, &s.SetupFee, &s.NetServicePriceInc, &s.NetServicePriceTr, &s.NetServicePriceOthers, &s.DiscountsOnSubsPrice, &s.DiscountsOnSubsAddon, &s.DiscountsOnServicePaidInInstallments, &s.DiscountsOnServiceInc, &s.DiscountsOnServiceTr, &s.DiscountsOnServiceOthers, &s.DiscountsOnSetupFee, &s.OtherPromotionsWithOutlay, &s.MonthsOfService, &s.SpecialPromotion, &s.NewSalesCommitment)
		if err != nil {
			return nil, err
		}

		res = append(res, &s)
	}

	return res, nil
}

// SubscriptionByID retrieves a row from 'ccdb_dupl.subscription' as a Subscription.
//
// Generated from index 'subscription__id_pkey'.
func SubscriptionByID(db XODB, id uint) (*Subscription, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`_id, client_id, package_id, partner_id, start_date, end_date, months_free, status, payment_terms, contract_status, contract_start_date, contract_end_date, monthly_payment, annual_payment, managed_account, change_of_plan_at_renewal, churn_rationale, months_to_churn, reconcillation, subscription_addon_price, service_paid_in_installments, setup_fee, net_service_price_INC, net_service_price_TR, net_service_price_others, discounts_on_subs_price, discounts_on_subs_addon, discounts_on_service_paid_in_installments, discounts_on_service_INC, discounts_on_service_TR, discounts_on_service_others, discounts_on_setup_fee, other_promotions_with_outlay, months_of_service, special_promotion, new_sales_commitment ` +
		`FROM ccdb_dupl.subscription ` +
		`WHERE _id = ?`

	// run query
	XOLog(sqlstr, id)
	s := Subscription{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&s.ID, &s.ClientID, &s.PackageID, &s.PartnerID, &s.StartDate, &s.EndDate, &s.MonthsFree, &s.Status, &s.PaymentTerms, &s.ContractStatus, &s.ContractStartDate, &s.ContractEndDate, &s.MonthlyPayment, &s.AnnualPayment, &s.ManagedAccount, &s.ChangeOfPlanAtRenewal, &s.ChurnRationale, &s.MonthsToChurn, &s.Reconcillation, &s.SubscriptionAddonPrice, &s.ServicePaidInInstallments, &s.SetupFee, &s.NetServicePriceInc, &s.NetServicePriceTr, &s.NetServicePriceOthers, &s.DiscountsOnSubsPrice, &s.DiscountsOnSubsAddon, &s.DiscountsOnServicePaidInInstallments, &s.DiscountsOnServiceInc, &s.DiscountsOnServiceTr, &s.DiscountsOnServiceOthers, &s.DiscountsOnSetupFee, &s.OtherPromotionsWithOutlay, &s.MonthsOfService, &s.SpecialPromotion, &s.NewSalesCommitment)
	if err != nil {
		return nil, err
	}

	return &s, nil
}

// SubscriptionsByPartnerID retrieves a row from 'ccdb_dupl.subscription' as a Subscription.
//
// Generated from index 'subscription_ibfk_3'.
func SubscriptionsByPartnerID(db XODB, partnerID sql.NullInt64) ([]*Subscription, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`_id, client_id, package_id, partner_id, start_date, end_date, months_free, status, payment_terms, contract_status, contract_start_date, contract_end_date, monthly_payment, annual_payment, managed_account, change_of_plan_at_renewal, churn_rationale, months_to_churn, reconcillation, subscription_addon_price, service_paid_in_installments, setup_fee, net_service_price_INC, net_service_price_TR, net_service_price_others, discounts_on_subs_price, discounts_on_subs_addon, discounts_on_service_paid_in_installments, discounts_on_service_INC, discounts_on_service_TR, discounts_on_service_others, discounts_on_setup_fee, other_promotions_with_outlay, months_of_service, special_promotion, new_sales_commitment ` +
		`FROM ccdb_dupl.subscription ` +
		`WHERE partner_id = ?`

	// run query
	XOLog(sqlstr, partnerID)
	q, err := db.Query(sqlstr, partnerID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Subscription{}
	for q.Next() {
		s := Subscription{
			_exists: true,
		}

		// scan
		err = q.Scan(&s.ID, &s.ClientID, &s.PackageID, &s.PartnerID, &s.StartDate, &s.EndDate, &s.MonthsFree, &s.Status, &s.PaymentTerms, &s.ContractStatus, &s.ContractStartDate, &s.ContractEndDate, &s.MonthlyPayment, &s.AnnualPayment, &s.ManagedAccount, &s.ChangeOfPlanAtRenewal, &s.ChurnRationale, &s.MonthsToChurn, &s.Reconcillation, &s.SubscriptionAddonPrice, &s.ServicePaidInInstallments, &s.SetupFee, &s.NetServicePriceInc, &s.NetServicePriceTr, &s.NetServicePriceOthers, &s.DiscountsOnSubsPrice, &s.DiscountsOnSubsAddon, &s.DiscountsOnServicePaidInInstallments, &s.DiscountsOnServiceInc, &s.DiscountsOnServiceTr, &s.DiscountsOnServiceOthers, &s.DiscountsOnSetupFee, &s.OtherPromotionsWithOutlay, &s.MonthsOfService, &s.SpecialPromotion, &s.NewSalesCommitment)
		if err != nil {
			return nil, err
		}

		res = append(res, &s)
	}

	return res, nil
}

func (s Subscription) All(db XODB) (interface{}, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`* ` +
		`FROM ccdb_dupl.subscription `

	// run query
	XOLog(sqlstr)
	q, err := db.Query(sqlstr)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Subscription{}
	for q.Next() {
		s := Subscription{
			_exists: true,
		}

		// scan
		err = q.Scan(&s.ID, &s.ClientID, &s.PackageID, &s.PartnerID, &s.StartDate, &s.EndDate, &s.MonthsFree, &s.Status, &s.PaymentTerms, &s.ContractStatus, &s.ContractStartDate, &s.ContractEndDate, &s.MonthlyPayment, &s.AnnualPayment, &s.ManagedAccount, &s.ChangeOfPlanAtRenewal, &s.ChurnRationale, &s.MonthsToChurn, &s.Reconcillation, &s.SubscriptionAddonPrice, &s.ServicePaidInInstallments, &s.SetupFee, &s.NetServicePriceInc, &s.NetServicePriceTr, &s.NetServicePriceOthers, &s.DiscountsOnSubsPrice, &s.DiscountsOnSubsAddon, &s.DiscountsOnServicePaidInInstallments, &s.DiscountsOnServiceInc, &s.DiscountsOnServiceTr, &s.DiscountsOnServiceOthers, &s.DiscountsOnSetupFee, &s.OtherPromotionsWithOutlay, &s.MonthsOfService, &s.SpecialPromotion, &s.NewSalesCommitment)
		if err != nil {
			return nil, err
		}

		res = append(res, &s)
	}

	return res, nil
}
func (s Subscription) ByID(db XODB, id uint) (interface{}, error) {
	return nil, nil
}
