package ent

import (
	"database/sql"
	"sync"

	entSql "entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
)

var onceDB sync.Once
var DB *sql.DB

func (c *Client) DB() *sql.DB {
	onceDB.Do(func() {
		DB = c.driver.(*entSql.Driver).DB()
	})
	return DB
}
