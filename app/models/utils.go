package models

import (
	"database/sql"
	"encoding/json"

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

func (v NullInt64) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Int64)
	} else {
		return json.Marshal(nil)
	}
}

func (v NullFloat64) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Float64)
	} else {
		return json.Marshal(nil)
	}
}

func (v NullBool) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Bool)
	} else {
		return json.Marshal(nil)
	}
}

func (v NullString) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.String)
	} else {
		return json.Marshal(nil)
	}
}

func (v NullTime) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Time)
	} else {
		return json.Marshal(nil)
	}
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
