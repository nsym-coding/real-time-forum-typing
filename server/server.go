package server

import "database/sql"

type MyServer struct {
	Db *sql.DB
}
