// Package models contains the types for schema 'ccdb_dupl'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"
)

// Partner represents a row from 'ccdb_dupl.partner'.
type Partner struct {
	ID      uint           `json:"_id"`     // _id
	Name    sql.NullString `json:"name"`    // name
	Type    sql.NullString `json:"type"`    // type
	Country sql.NullString `json:"country"` // country

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Partner exists in the database.
func (p *Partner) Exists() bool {
	return p._exists
}

// Deleted provides information if the Partner has been deleted from the database.
func (p *Partner) Deleted() bool {
	return p._deleted
}

// Insert inserts the Partner to the database.
func (p *Partner) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if p._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO ccdb_dupl.partner (` +
		`name, type, country` +
		`) VALUES (` +
		`?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, p.Name, p.Type, p.Country)
	res, err := db.Exec(sqlstr, p.Name, p.Type, p.Country)
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

// Update updates the Partner in the database.
func (p *Partner) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !p._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if p._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE ccdb_dupl.partner SET ` +
		`name = ?, type = ?, country = ?` +
		` WHERE _id = ?`

	// run query
	XOLog(sqlstr, p.Name, p.Type, p.Country, p.ID)
	_, err = db.Exec(sqlstr, p.Name, p.Type, p.Country, p.ID)
	return err
}

// Save saves the Partner to the database.
func (p *Partner) Save(db XODB) error {
	if p.Exists() {
		return p.Update(db)
	}

	return p.Insert(db)
}

// Delete deletes the Partner from the database.
func (p *Partner) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !p._exists {
		return nil
	}

	// if deleted, bail
	if p._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM ccdb_dupl.partner WHERE _id = ?`

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

// PartnerByID retrieves a row from 'ccdb_dupl.partner' as a Partner.
//
// Generated from index 'partner__id_pkey'.
func PartnerByID(db XODB, id uint) (*Partner, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`_id, name, type, country ` +
		`FROM ccdb_dupl.partner ` +
		`WHERE _id = ?`

	// run query
	XOLog(sqlstr, id)
	p := Partner{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&p.ID, &p.Name, &p.Type, &p.Country)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func Partners(db XODB) ([]*Partner, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`* ` +
		`FROM ccdb_dupl.partner `

	// run query
	XOLog(sqlstr)
	q, err := db.Query(sqlstr)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Partner{}
	for q.Next() {
		p := Partner{
			_exists: true,
		}

		// scan
		err = q.Scan(&p.ID, &p.Name, &p.Type, &p.Country)
		if err != nil {
			return nil, err
		}

		res = append(res, &p)
	}

	return res, nil
}
