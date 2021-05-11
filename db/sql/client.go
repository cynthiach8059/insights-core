package sql

import (
	"database/sql"
	"errors"
	"log"
	"os"
	"sync"

	_ "github.com/denisenkom/go-mssqldb"
)

var (
	configOptions    []Config
	DbErrNoDocuments = errors.New("No documents matched")
)

type Client struct {
	id          string
	isConnected bool
	connStr     string
	reconnects  int
	sync.RWMutex
	*sql.DB
}

func (c *Client) Connect() error {
	if c.DB != nil {
		return nil
	}
	var err interface{}
	constr := os.Getenv("INSIGHTS_DB")
	c.Lock()

	c.DB, err = sql.Open("mssql", constr)
	if err != nil {
		log.Fatal(err)
	}

	c.isConnected = true
	c.Unlock()

	return nil
}

func (c *Client) IsConnected() bool {
	if c.DB != nil {
		err := c.DB.Ping()
		return err == nil
	}
	return false
}

type Config func(*Client)

func Init(configs ...Config) {
	configOptions = configs
}

func SetClientID(id string) Config {
	return func(c *Client) {
		c.id = id
	}
}

func SetConnectionString(conn string) Config {
	return func(c *Client) {
		c.connStr = conn
	}
}
