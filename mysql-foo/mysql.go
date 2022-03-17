package mysqlfoo

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

const (
	MYSQL_MAX_OPEN_CONNECTIONS = 1000
	MYSQL_MAX_IDLE_CONNECTIONS = 100 // 并发协程数 = 保活链接数 / 10
)

func OpenMySQLDatabase(URI string) (*sql.DB, error) {
	db, openError := sql.Open("mysql", URI)
	if db == nil || openError != nil {
		return nil, openError
	}

	db.SetMaxOpenConns(MYSQL_MAX_OPEN_CONNECTIONS)
	db.SetMaxIdleConns(MYSQL_MAX_IDLE_CONNECTIONS)

	if pingError := db.Ping(); pingError != nil {
		return nil, pingError
	}

	return db, nil
}
