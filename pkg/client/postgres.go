package client

import (
	"context"

	"github.com/jmoiron/sqlx"
)

// postgres struct is in charge of perform all the postgres related queries,
// without the client knowing.
type postgres struct {
	db     *sqlx.DB
	dbName string
	schema string
}

// a validation to see if postgres is implementing databaseQuerier.
var _ databaseQuerier = (*postgres)(nil)

// returns a pointer to a postgres, it receives an schema as a parameter.
func newPostgres(dbName, schema string, db *sqlx.DB) *postgres {
	_ = "STUB: not implemented"
	return nil
}

// TableStructure returns a query string to get all the relevant information of a given table,
// under a schema.
func (p *postgres) TableStructure(table TableRef) (string, []any, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// Constraints returns all the constraints of a given table, under a schema.
func (p *postgres) Constraints(table TableRef) (string, []any, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// Indexes returns the indexes of a table, under a schema.
func (p *postgres) Indexes(table TableRef) (string, []any, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func (p *postgres) Catalog(ctx context.Context) (*DBNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *postgres) fetchSchemas(ctx context.Context, parentID string) ([]*DBNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *postgres) fetchTables(ctx context.Context, parentName, parentID string) ([]*DBNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
