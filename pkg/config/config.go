package config

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/spf13/cobra"

	"github.com/danvergara/dblab/pkg/command"
)

// Config struct is used to store the db connection data.
type Config struct {
	Database []Database
	User     string
	Pswd     string
	Host     string
	Port     string
	DBName   string
	Driver   string
	Limit    uint `fig:"limit" default:"100"`
}

type KeyMapConfig struct {
	KeyBindings KeyBindings
}

type Database struct {
	Name string `fig:"name"`

	Host     string
	Port     string
	DB       string
	User     string
	Password string
	Driver   string `validate:"required"`
	Schema   string

	// SSH Tunnel.
	SSHHost          string `fig:"ssh-host"`
	SSHPort          string `fig:"ssh-port"`
	SSHUser          string `fig:"ssh-user"`
	SSHPass          string `fig:"ssh-pass"`
	SSHKeyFile       string `fig:"ssh-key-file"`
	SSHKeyPassphrase string `fig:"ssh-key-pass"`

	// SSL connection params.
	SSL string `default:"disable"`

	SSLCert     string `fig:"sslcert"`
	SSLKey      string `fig:"sslkey"`
	SSLPassword string `fig:"sslpassword"`
	SSLRootcert string `fig:"sslrootcert"`

	// oracle specific.
	TraceFile string `fig:"trace"`
	SSLVerify string `fig:"ssl-verify"`
	Wallet    string `fig:"wallet"`

	// sql server.
	Encrypt                string `fig:"encrypt"`
	TrustServerCertificate string `fig:"trust-server-certificate"`
	ConnectionTimeout      string `fig:"connection-timeout"`
}

type KeyBindings struct {
	NextTab         string `fig:"next-tab"      default:"tab"`
	PrevTab         string `fig:"prev-tab"      default:"shift+tab"`
	PageTop         string `fig:"page-top"      default:"g"`
	PageBottom      string `fig:"page-bottom"   default:"G"`
	EndOfLine       string `fig:"end-of-line"   default:"$"`
	BeginningOfLine string `fig:"beginning-of-line"   default:"0"`
	Navigation      NavigationBindgins
	Editor          EditorKeyMap
}

type EditorKeyMap struct {
	// Normal Mode Navigation.
	Up    string `fig:"up" default:"k"`
	Down  string `fig:"down" default:"j"`
	Left  string `fig:"left" default:"h"`
	Right string `fig:"right" default:"l"`

	// Mode Switching.
	Insert string `fig:"insert" default:"i"`
	Normal string `fig:"normal" default:"esc"`

	// Actions.
	ExecuteQuery string `fig:"execute-query" default:"ctrl+e"`
}

type NavigationBindgins struct {
	Up    string `fig:"up"    default:"ctrl+k"`
	Down  string `fig:"down"  default:"ctrl+j"`
	Left  string `fig:"left"  default:"ctrl+h"`
	Right string `fig:"right" default:"ctrl+l"`
}

// New returns a config instance the with db connection data inplace based on the flags of a cobra command.
func New(cmd *cobra.Command) *Config { _ = "STUB: not implemented"; return nil }

// Init reads in config file and returns a commands/Options instance.
func Init(configName string) (command.Options, error) {
	_ = "STUB: not implemented"
	return *new(command.Options), nil
}

func SetupKeyMap() (command.TUIKeyMap, error) {
	_ = "STUB: not implemented"
	return *new(command.TUIKeyMap), nil
}

// Open returns a db connection using the data from the config object.
func (c *Config) Open() (*sql.DB, error) { _ = "STUB: not implemented"; return nil, nil }

// MigrateInstance returns a migrate instance based on the given driver.
func (c *Config) MigrateInstance() (*migrate.Migrate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get returns a config object with the db connection data already in place.
func Get() *Config { _ = "STUB: not implemented"; return nil }

// GetDBConnStr returns the connection string.
func (c *Config) GetDBConnStr() string { _ = "STUB: not implemented"; return "" }

// GetSQLXDBConnStr returns the connection string.
func (c *Config) GetSQLXDBConnStr() string { _ = "STUB: not implemented"; return "" }

// getDBConnStr returns the connection string based on the provided host and db name.
func (c *Config) getDBConnStr(dbhost, dbname string) string { _ = "STUB: not implemented"; return "" }

// getSQLXConnStr returns the connection string based on the provided host and db name.
func (c *Config) getSQLXConnStr(dbhost, dbname string) string { _ = "STUB: not implemented"; return "" }
