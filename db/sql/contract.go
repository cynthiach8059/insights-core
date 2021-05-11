package sql

import (
	_ "github.com/go-sql-driver/mysql"
)

type StorageDB interface {
	Connect() error
	IsConnected() bool
	Disconnect()
}
