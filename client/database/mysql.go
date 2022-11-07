package database

import "github.com/gocraft/dbr/v2"

type Database struct {
	db     *dbr.Session
	dbRead *dbr.Session
}

func NewMysqlDatabase(host, readHost, username, pwd, database string) *Database {
	return nil
}
