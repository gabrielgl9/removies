package postgresql

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type ConfigPostgreDb struct {
	db_user string
	db_pass string
	db_name string
	db_host string
	db_port int
}

func (configDb *ConfigPostgreDb) NewConfigPostgreDb() {
	configDb.db_user = "root"
	configDb.db_pass = "admin"
	configDb.db_name = "removies"
	configDb.db_host = "localhost"
	configDb.db_port = 5432
}

func (configDb *ConfigPostgreDb) SetupDB() (*sql.DB, error) {
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", 
		configDb.db_host, configDb.db_port, configDb.db_user, configDb.db_pass, configDb.db_name)

	db, err := sql.Open("postgres", dbinfo)
	
	if err != nil {
		return nil, err
	}

	// defer db.Close()
	
	return db, err
}