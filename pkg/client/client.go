package client

import (
	"context"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	_ "github.com/microsoft/go-mssqldb"
	_ "github.com/sijms/go-ora/v2"
	_ "modernc.org/sqlite"

	"github.com/danvergara/dblab/pkg/command"
	"github.com/danvergara/dblab/pkg/pagination"
)

type TableRef struct {
	Schema string
	Name   string
}

type DBNode struct {
	ID         string
	Name       string
	Type       string
	ParentID   string
	ParentName string
	Children   []*DBNode
}

// databaseQuerier is an interface that indicates the methods
// a given type has to implement to interact with a database,
// to get specific data.
// This allows us to decouple the client from the database implementation and
// make adding new databases easier.
type databaseQuerier interface {
	TableStructure(table TableRef) (string, []any, error)
	Constraints(table TableRef) (string, []any, error)
	Indexes(table TableRef) (string, []any, error)
	Catalog(context.Context) (*DBNode, error)
}

// Client is used to store the pool of db connection.
type Client struct {
	db                *sqlx.DB
	dbName            string
	databaseQuerier   databaseQuerier
	driver, schema    string
	host              string
	paginationManager *pagination.Manager
	limit             uint
}

// New return an instance of the client.
func New(opts command.Options) (*Client, error) { _ = "STUB: not implemented"; return nil, nil }

// This is where an implementation of databaseQuerier is getting picked up.

// DB Return the db attribute.
func (c *Client) DB() *sqlx.DB {
	_ = "STUB: not implemented"

	// Driver returns the driver of the database.
	return nil
}

func (c *Client) Driver() string { _ = "STUB: not implemented"; return "" }

func (c *Client) Host() string {
	_ = "STUB: not implemented"

	// Query returns performs the query and returns the result set and the column names.
	return ""
}

func (c *Client) Query(q string, args ...any) ([][]string, []string, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil

	// Runs the query extracting the content of the view calling the Buffer method.
}

// Gets the names of the columns of the result set.

// cols is an []any of all of the column results.

// Convert []any into []string.

// Table represents a SQL table.
type Table struct {
	name    string
	Rows    [][]string
	Columns []string
}

func (t *Table) Name() string {
	_ = "STUB: not implemented"

	// Metadata sums up the most relevant data from a table.
	return ""
}

type Metadata struct {
	TableContent Table
	Structure    Table
	Constraints  Table
	Indexes      Table
	TotalPages   int
}

// Metadata returns the most relevant data from a given table.
func (c *Client) Metadata(table TableRef) (*Metadata, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TableContent returns all the rows of a table.
func (c *Client) tableContent(table TableRef) ([][]string, []string, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// tableStructure returns the structure of the table columns.
func (c *Client) tableStructure(table TableRef) ([][]string, []string, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// constraints returns the resultet of from information_schema.table_constraints.
func (c *Client) constraints(table TableRef) ([][]string, []string, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// indexes returns a resulset with the information of the indexes given a table name.
func (c *Client) indexes(table TableRef) ([][]string, []string, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (c *Client) Catalog(ctx context.Context) (*DBNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
