package client

import (
	"context"

	"github.com/jmoiron/sqlx"
)

// sqlite  is in charge of perform all the sqlite related queries,
// without the client knowing.
type sqlite struct {
	db     *sqlx.DB
	dbName string
}

// a validation to see if sqlite is implementing databaseQuerier.
var _ databaseQuerier = (*sqlite)(nil)

// returns a pointer to a sqlite.
func newSQLite(dbName string, db *sqlx.DB) *sqlite { _ = "STUB: not implemented"; return nil }

// TableStructure returns a query string to retrieve all the relevant information of a given table.
func (s *sqlite) TableStructure(table TableRef) (string, []interface{}, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// Constraints returns all the constraints of a given table.
func (s *sqlite) Constraints(table TableRef) (string, []interface{}, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// Indexes returns a query to get all the indexes of a table.
func (s *sqlite) Indexes(table TableRef) (string, []interface{}, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func (s *sqlite) Catalog(ctx context.Context) (*DBNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *sqlite) fetchTables(ctx context.Context, parentName, parentID string) ([]*DBNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
