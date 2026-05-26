package client

import (
	"context"

	"github.com/jmoiron/sqlx"
)

// mssql struct is in charge of perform all the SQL Server related queries.
type mssql struct {
	db     *sqlx.DB
	dbName string
}

var _ databaseQuerier = (*mssql)(nil)

// returns a pointer to a mysql.
func newMSSQL(dbName string, db *sqlx.DB) *mssql { _ = "STUB: not implemented"; return nil }

// TableStructure returns a query string to retrieve all the relevant information of a given table.
func (m *mssql) TableStructure(table TableRef) (string, []interface{}, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// Constraints returns all the constraints of a given table.
func (m *mssql) Constraints(table TableRef) (string, []interface{}, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// Indexes returns the indexes of a table.
func (m *mssql) Indexes(table TableRef) (string, []interface{}, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func (m *mssql) Catalog(ctx context.Context) (*DBNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *mssql) fetchTables(ctx context.Context, parentName, parentID string) ([]*DBNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
