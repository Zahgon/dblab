package connection

import (
	"errors"
	"regexp"

	"github.com/danvergara/dblab/pkg/command"
)

var (
	// pattern used to parse an incoming dsn for mysql connection.
	dsnPattern *regexp.Regexp
	// ErrCantDetectUSer is the error used to notify that a default username is not found
	// in the system to be used as database username.
	ErrCantDetectUSer = errors.New("could not detect default username")
	// ErrInvalidPostgresURLFormat is the error used to notify that the postgres given url is not valid.
	ErrInvalidPostgresURLFormat = errors.New(
		"invalid url - valid format: postgres://user:password@host:port/db?sslmode=mode",
	)
	// ErrInvalidMySQLURLFormat is the error used to notify that the given mysql url is not valid.
	ErrInvalidMySQLURLFormat = errors.New(
		"invalid url - valid format: mysql://user:password@tcp(host:port)/db",
	)
	// ErrInvalidOracleURLFormat is the error used to notify the user that the oracle url is invalid.
	ErrInvalidOracleURLFormat = errors.New(
		"invalid url - valid format: oracle://user:pass@server/service_name",
	)
	// ErrInvalidURLFormat is used to notify the url is invalid.
	ErrInvalidURLFormat = errors.New("invalid url")
	// ErrInvalidDriver is used to notify that the provided driver is not supported.
	ErrInvalidDriver = errors.New("invalid driver")
	// ErrSocketFileDoNotExist indicates that the given path to the socket files leads to no file.
	ErrSocketFileDoNotExist = errors.New("socket file does not exist")
	// ErrInvalidSocketFile indicates that the socket file must end with .sock as suffix.
	ErrInvalidSocketFile = errors.New("invalid socket file - must end with .sock")
	// ErrInvalidOraclePort indicates that the port passed is not a proper integer.
	ErrInvalidOraclePort = errors.New("invalid oracle port")
)

func init() {
	dsnPattern = regexp.MustCompile(
		`^(?:(?P<user>.*?)(?::(?P<passwd>.*))?@)?` + // [user[:password]@]
			`(?:(?P<net>[^\(]*)(?:\((?P<addr>[^\)]*)\))?)?` + // [net[(addr)]]
			`\/(?P<dbname>.*?)` + // /dbname
			`(?:\?(?P<params>[^\?]*))?$`) // [?param1=value1&paramN=valueN]
}

// BuildConnectionFromOpts return the connection uri string given the options passed by the uses.
func BuildConnectionFromOpts(opts command.Options) (string, command.Options, error) {
	_ = "STUB: not implemented"
	return "", *new(command.Options), nil
}

// This options is for sqlite.
// For more information see https://github.com/mattn/go-sqlite3#connection-string.
// The sqlite driver is modernc.org/sqlite now, so the query params changed.
// mattn's driver and modernc.org's differ in the query paremeters they use to perform a connection.
// To know more about the connection and query paremeters see: https://pkg.go.dev/modernc.org/sqlite#Driver.Open

func currentUser() (string, error) { _ = "STUB: not implemented"; return "", nil }

// formatPostgresURL returns valid uri for postgres connection.
func formatPostgresURL(opts command.Options) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// formatMySQLURL returns valid uri for mysql connection.
func formatMySQLURL(opts command.Options) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// removes the mysql:// scheme since mysql does not need this.
// we need it to know what database we're trying to connect to.

// checks if *url.Error is the type of the error.
// if the url is a dsn for mysql connection
// the most likely is this is gonna be true.

// formatOracleURL returns valid uri for oracle connection.
func formatOracleURL(opts command.Options) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// formatSQLServerURL returns valid uri for sql server connection.
func formatSQLServerURL(opts command.Options) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// validates if dsn pattern match with the parameter.
func parseDSN(dsn string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// hasValidPostgresPrefix checks if a given url has the driver name in it.
func hasValidPostgresPrefix(rawurl string) bool { _ = "STUB: not implemented"; return false }

// hasValidMySQLPrefix checks if a given url has the driver name in it.
func hasValidMySQLPrefix(rawurl string) bool { _ = "STUB: not implemented"; return false }

// hasValidOraclePrefix checks if a given url has the driver name in it.
func hasValidOraclePrefix(rawurl string) bool { _ = "STUB: not implemented"; return false }

// hasValidSQLServerPrefix checks if a given url has the driver name in it.
func hasValidSQLServerPrefix(rawurl string) bool { _ = "STUB: not implemented"; return false }

func socketFileExists(socketFile string) bool { _ = "STUB: not implemented"; return false }

func validSocketFile(socketFile string) bool { _ = "STUB: not implemented"; return false }
