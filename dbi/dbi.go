//Package dbi provide simple tools to interact with database.
package dbi

import (
	"database/sql"

	"golang.org/x/net/context"
)

var (
	//the key for the map is id for registered db.
	ds map[string]db
)

type DBExecer interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

//RegisterDB connect to the database and store the db on gloabl variables to be reuse.
//Register will return an error if the id already exist or can't open the database.
func RegisterDB(driverName, dataSourceName, id string) error {
	return nil
}

//NewDBContext create a context from registered db that carrying db.
func NewDBContext(ctx context.Context) DBContext {
	return nil
}

type DBContext interface {
	context.Context
	Execer(ctx context.Context) DBExecer
}

type db struct {
	*sql.DB
	builder QueryBuilder
}

type QueryBuilder interface {
	Select() (query string, args []interface{}, err error)
	Insert() string
	Delete() string
	Update() string
}

//RunInTransaction create a DBContext that carrying tx,than call fn with the context.
//If fn return an error RunInTransaction will do a rollback, otherwise will commit the transaction.
func RunInTransaction(db *sql.DB, fn func(ctx DBContext) error) error {
	return nil
}

//CreateTable create table on the database.
func CreateTable(ctx DBContext) error {
	return nil
}

//DropTable drop table from the database.
func DropTable(ctx DBContext) error {
	return nil
}

//AlterTable alter table from the database.
func AlterTable(ctx DBContext) error {
	return nil
}
