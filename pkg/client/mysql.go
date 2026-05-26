package client

import (
	"context"

	"github.com/jmoiron/sqlx"
)

// mysql struct is in charge of perform all the mysql related queries.
type mysql struct {
	db     *sqlx.DB
	dbName string
}

// a validation to see if mysql is implementing databaseQuerier.
var _ databaseQuerier = (*mysql)(nil)

// returns a pointer to a mysql.
func newMySQL(dbName string, db *sqlx.DB) *mysql { _ = "STUB: not implemented"; return nil }

// TableStructure returns a query string to retrieve all the relevant information of a given table.
func (m *mysql) TableStructure(table TableRef) (string, []any, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// Constraints returns all the constraints of a given table.
func (m *mysql) Constraints(table TableRef) (string, []any, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// Indexes returns a query to get all the indexes of a table.
func (m *mysql) Indexes(table TableRef) (string, []any, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func (m *mysql) Catalog(ctx context.Context) (*DBNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *mysql) fetchTables(ctx context.Context, parentName, parentID string) ([]*DBNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
