package client

import (
	"context"

	"github.com/jmoiron/sqlx"
	_ "github.com/sijms/go-ora/v2"
)

// oracle struct is in charge of perform all the oracle related queries.
type oracle struct {
	db     *sqlx.DB
	dbName string
	schema string
}

// a validation to see if oracle is implementing databaseQuerier.
var _ databaseQuerier = (*oracle)(nil)

// returns a pointer to a oracle struct, it receives an schema as a parameter.
func newOracle(dbName, schema string, db *sqlx.DB) *oracle {
	_ = "STUB: not implemented"
	// If the schema is no empty, the client queries against the ALL_* tables,
	// where the OWNER is equal to the schema/user the dblab user has access to.
	// Otherwise, the client will query against the USER_* tables,
	// meaning it only cares about what the user has direct access to.
	return nil
}

// TableStructure returns a query string to get all the relevant information of a given table.
func (o *oracle) TableStructure(table TableRef) (string, []interface{}, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// Constraints returns all the constraints of a given table.
func (o *oracle) Constraints(table TableRef) (string, []interface{}, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// Indexes returns the indexes of a table.
func (o *oracle) Indexes(table TableRef) (string, []interface{}, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func (o *oracle) Catalog(ctx context.Context) (*DBNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (o *oracle) fetchSchemas(ctx context.Context, parentID string) ([]*DBNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (o *oracle) fetchTables(ctx context.Context, parentName, parentID string) ([]*DBNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
