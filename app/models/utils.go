package models

import (
	"database/sql"
	"encoding/json"
	"reflect"

	"github.com/go-sql-driver/mysql"
)

type Modeler interface {
	All(db XODB) (interface{}, error)
	ByID(db XODB, id uint) error
}

type NullInt64 sql.NullInt64
type NullString sql.NullString
type NullBool sql.NullBool
type NullFloat64 sql.NullFloat64
type NullTime mysql.NullTime

// Scan implements the Scanner interface for NullInt64
func (ni *NullInt64) Scan(value interface{}) error {
	var i sql.NullInt64
	if err := i.Scan(value); err != nil {
		return err
	}

	// if nil then make Valid false
	if reflect.TypeOf(value) == nil {
		*ni = NullInt64{i.Int64, false}
	} else {
		*ni = NullInt64{i.Int64, true}
	}
	return nil
}

// Scan implements the Scanner interface for NullBool
func (nb *NullBool) Scan(value interface{}) error {
	var b sql.NullBool
	if err := b.Scan(value); err != nil {
		return err
	}

	// if nil then make Valid false
	if reflect.TypeOf(value) == nil {
		*nb = NullBool{b.Bool, false}
	} else {
		*nb = NullBool{b.Bool, true}
	}

	return nil
}

// Scan implements the Scanner interface for NullFloat64
func (nf *NullFloat64) Scan(value interface{}) error {
	var f sql.NullFloat64
	if err := f.Scan(value); err != nil {
		return err
	}

	// if nil then make Valid false
	if reflect.TypeOf(value) == nil {
		*nf = NullFloat64{f.Float64, false}
	} else {
		*nf = NullFloat64{f.Float64, true}
	}

	return nil
}

// Scan implements the Scanner interface for NullString
func (ns *NullString) Scan(value interface{}) error {
	var s sql.NullString
	if err := s.Scan(value); err != nil {
		return err
	}

	// if nil then make Valid false
	if reflect.TypeOf(value) == nil {
		*ns = NullString{s.String, false}
	} else {
		*ns = NullString{s.String, true}
	}

	return nil
}

// Scan implements the Scanner interface for NullTime
func (nt *NullTime) Scan(value interface{}) error {
	var t mysql.NullTime
	if err := t.Scan(value); err != nil {
		return err
	}

	// if nil then make Valid false
	if reflect.TypeOf(value) == nil {
		*nt = NullTime{t.Time, false}
	} else {
		*nt = NullTime{t.Time, true}
	}

	return nil
}

// MarshalJSON implements Marshler interface
func (v NullInt64) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Int64)
	}
	return json.Marshal(nil)
}

func (v NullFloat64) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Float64)
	}
	return json.Marshal(nil)
}

func (v NullBool) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Bool)
	}
	return json.Marshal(nil)
}

func (v NullString) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.String)
	}
	return json.Marshal(nil)
}

func (v NullTime) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Time)
	}
	return json.Marshal(nil)
}

func (v *NullInt64) UnmarshalJSON(data []byte) error {
	err := json.Unmarshal(data, &v.Int64)
	v.Valid = (err == nil)
	return err
}

func (v *NullFloat64) UnmarshalJSON(data []byte) error {
	err := json.Unmarshal(data, &v.Float64)
	v.Valid = (err == nil)
	return err
}

func (v *NullBool) UnmarshalJSON(data []byte) error {
	err := json.Unmarshal(data, &v.Bool)
	v.Valid = (err == nil)
	return err
}

func (v *NullString) UnmarshalJSON(data []byte) error {
	err := json.Unmarshal(data, &v.String)
	v.Valid = (err == nil)
	return err
}

func (v *NullTime) UnmarshalJSON(data []byte) error {
	err := json.Unmarshal(data, &v.Time)
	v.Valid = (err == nil)
	return err
}
