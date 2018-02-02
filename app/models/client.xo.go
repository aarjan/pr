// Package models contains the types for schema 'ccdb_dupl'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"
	"time"

	"github.com/go-sql-driver/mysql"
)

// Client represents a row from 'ccdb_dupl.client'.
type Client struct {
	ID                         uint               `json:"_id"`                            // _id
	AppUUID                    sql.NullString     `json:"app_uuid"`                       // app_uuid
	Name                       sql.NullString     `json:"name"`                           // name
	Type                       sql.NullString     `json:"type"`                           // type
	LegalEntityBooking         LegalEntityBooking `json:"legal_entity_booking"`           // legal_entity_booking
	NumberOfEmployees          NumberOfEmployee   `json:"number_of_employees"`            // number_of_employees
	BusinessSince              sql.NullInt64      `json:"business_since"`                 // business_since
	Segment                    sql.NullString     `json:"segment"`                        // segment
	LeadSource                 sql.NullString     `json:"lead_source"`                    // lead_source
	ReasonForAllocationToSales sql.NullString     `json:"reason_for_allocation_to_sales"` // reason_for_allocation_to_sales
	ResponsibleRm              sql.NullString     `json:"responsible_RM"`                 // responsible_RM
	AcquisitionDate            mysql.NullTime     `json:"acquisition_date"`               // acquisition_date
	AcquisitionCohort          mysql.NullTime     `json:"acquisition_cohort"`             // acquisition_cohort
	ClientChurnDate            mysql.NullTime     `json:"client_churn_date"`              // client_churn_date
	Timestamp                  time.Time          `json:"timestamp"`                      // timestamp
	SalesPersonID              sql.NullInt64      `json:"sales_person_id"`                // sales_person_id

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Client exists in the database.
func (c *Client) Exists() bool {
	return c._exists
}

// Deleted provides information if the Client has been deleted from the database.
func (c *Client) Deleted() bool {
	return c._deleted
}

// Insert inserts the Client to the database.
func (c *Client) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if c._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO ccdb_dupl.client (` +
		`app_uuid, name, type, legal_entity_booking, number_of_employees, business_since, segment, lead_source, reason_for_allocation_to_sales, responsible_RM, acquisition_date, acquisition_cohort, client_churn_date, timestamp, sales_person_id` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, c.AppUUID, c.Name, c.Type, c.LegalEntityBooking, c.NumberOfEmployees, c.BusinessSince, c.Segment, c.LeadSource, c.ReasonForAllocationToSales, c.ResponsibleRm, c.AcquisitionDate, c.AcquisitionCohort, c.ClientChurnDate, c.Timestamp, c.SalesPersonID)
	res, err := db.Exec(sqlstr, c.AppUUID, c.Name, c.Type, c.LegalEntityBooking, c.NumberOfEmployees, c.BusinessSince, c.Segment, c.LeadSource, c.ReasonForAllocationToSales, c.ResponsibleRm, c.AcquisitionDate, c.AcquisitionCohort, c.ClientChurnDate, c.Timestamp, c.SalesPersonID)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	c.ID = uint(id)
	c._exists = true

	return nil
}

// Update updates the Client in the database.
func (c *Client) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !c._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if c._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE ccdb_dupl.client SET ` +
		`app_uuid = ?, name = ?, type = ?, legal_entity_booking = ?, number_of_employees = ?, business_since = ?, segment = ?, lead_source = ?, reason_for_allocation_to_sales = ?, responsible_RM = ?, acquisition_date = ?, acquisition_cohort = ?, client_churn_date = ?, timestamp = ?, sales_person_id = ?` +
		` WHERE _id = ?`

	// run query
	XOLog(sqlstr, c.AppUUID, c.Name, c.Type, c.LegalEntityBooking, c.NumberOfEmployees, c.BusinessSince, c.Segment, c.LeadSource, c.ReasonForAllocationToSales, c.ResponsibleRm, c.AcquisitionDate, c.AcquisitionCohort, c.ClientChurnDate, c.Timestamp, c.SalesPersonID, c.ID)
	_, err = db.Exec(sqlstr, c.AppUUID, c.Name, c.Type, c.LegalEntityBooking, c.NumberOfEmployees, c.BusinessSince, c.Segment, c.LeadSource, c.ReasonForAllocationToSales, c.ResponsibleRm, c.AcquisitionDate, c.AcquisitionCohort, c.ClientChurnDate, c.Timestamp, c.SalesPersonID, c.ID)
	return err
}

// Save saves the Client to the database.
func (c *Client) Save(db XODB) error {
	if c.Exists() {
		return c.Update(db)
	}

	return c.Insert(db)
}

// Delete deletes the Client from the database.
func (c *Client) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !c._exists {
		return nil
	}

	// if deleted, bail
	if c._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM ccdb_dupl.client WHERE _id = ?`

	// run query
	XOLog(sqlstr, c.ID)
	_, err = db.Exec(sqlstr, c.ID)
	if err != nil {
		return err
	}

	// set deleted
	c._deleted = true

	return nil
}

// SalesPerson returns the SalesPerson associated with the Client's SalesPersonID (sales_person_id).
//
// Generated from foreign key 'client_ibfk_1'.
func (c *Client) SalesPerson(db XODB) (*SalesPerson, error) {
	return SalesPersonByID(db, uint(c.SalesPersonID.Int64))
}

// ClientByID retrieves a row from 'ccdb_dupl.client' as a Client.
//
// Generated from index 'client__id_pkey'.
func ClientByID(db XODB, id uint) (*Client, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`_id, app_uuid, name, type, legal_entity_booking, number_of_employees, business_since, segment, lead_source, reason_for_allocation_to_sales, responsible_RM, acquisition_date, acquisition_cohort, client_churn_date, timestamp, sales_person_id ` +
		`FROM ccdb_dupl.client ` +
		`WHERE _id = ?`

	// run query
	XOLog(sqlstr, id)
	c := Client{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&c.ID, &c.AppUUID, &c.Name, &c.Type, &c.LegalEntityBooking, &c.NumberOfEmployees, &c.BusinessSince, &c.Segment, &c.LeadSource, &c.ReasonForAllocationToSales, &c.ResponsibleRm, &c.AcquisitionDate, &c.AcquisitionCohort, &c.ClientChurnDate, &c.Timestamp, &c.SalesPersonID)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

// ClientsBySalesPersonID retrieves a row from 'ccdb_dupl.client' as a Client.
//
// Generated from index 'client_ibfk_1'.
func ClientsBySalesPersonID(db XODB, salesPersonID sql.NullInt64) ([]*Client, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`_id, app_uuid, name, type, legal_entity_booking, number_of_employees, business_since, segment, lead_source, reason_for_allocation_to_sales, responsible_RM, acquisition_date, acquisition_cohort, client_churn_date, timestamp, sales_person_id ` +
		`FROM ccdb_dupl.client ` +
		`WHERE sales_person_id = ?`

	// run query
	XOLog(sqlstr, salesPersonID)
	q, err := db.Query(sqlstr, salesPersonID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Client{}
	for q.Next() {
		c := Client{
			_exists: true,
		}

		// scan
		err = q.Scan(&c.ID, &c.AppUUID, &c.Name, &c.Type, &c.LegalEntityBooking, &c.NumberOfEmployees, &c.BusinessSince, &c.Segment, &c.LeadSource, &c.ReasonForAllocationToSales, &c.ResponsibleRm, &c.AcquisitionDate, &c.AcquisitionCohort, &c.ClientChurnDate, &c.Timestamp, &c.SalesPersonID)
		if err != nil {
			return nil, err
		}

		res = append(res, &c)
	}

	return res, nil
}

func Clients(db XODB) ([]*Client, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`* ` +
		`FROM ccdb_dupl.client `

	// run query
	XOLog(sqlstr)
	q, err := db.Query(sqlstr)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Client{}
	for q.Next() {
		c := Client{
			_exists: true,
		}

		// scan
		err = q.Scan(&c.ID, &c.AppUUID, &c.Name, &c.Type, &c.LegalEntityBooking, &c.NumberOfEmployees, &c.BusinessSince, &c.Segment, &c.LeadSource, &c.ReasonForAllocationToSales, &c.ResponsibleRm, &c.AcquisitionDate, &c.AcquisitionCohort, &c.ClientChurnDate, &c.Timestamp, &c.SalesPersonID)
		if err != nil {
			return nil, err
		}

		res = append(res, &c)
	}

	return res, nil
}
