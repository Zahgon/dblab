package seeds

import (

	// mysql driver.
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	// postgres driver.
	_ "github.com/lib/pq"

	// sqlite driver.
	_ "modernc.org/sqlite"
)

// Seed type.
type Seed struct {
	db     *sqlx.DB
	driver string
}

// Execute will executes the given seeder method.
func Execute(db *sqlx.DB, driver string, seedMethodNames ...string) {
	_ = "STUB: not implemented"
	return
}

// Executes all seeders if no method is given.

// We are looping over the method on a Seed struct.

// Get the method in the current iteration.

// Execute seeder.

// Execute only the given method names

func seed(s Seed, seedMethodName string) {
	_ = "STUB: not implemented"
	// Get the reflect value of the method.
	return
}

// Exit if the method doesn't exist.

// Execute the method.
