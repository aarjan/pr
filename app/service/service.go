package service

import (
	"database/sql"
)

type AppServicer interface {
	// All(db *sql.DB) (interface{}, error)
}
type AppServer struct {
	DB *sql.DB
}
