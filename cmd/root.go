package cmd

import (
	"github.com/spf13/cobra"
)

// Define all the global flags.
var (
	cfg         bool
	cfgName     string
	driver      string
	url         string
	host        string
	port        string
	user        string
	pass        string
	schema      string
	db          string
	ssl         string
	limit       uint
	socket      string
	sslcert     string
	sslkey      string
	sslpassword string
	sslrootcert string

	// SSH Tunnel.
	sshHost          string
	sshPort          string
	sshUser          string
	sshPass          string
	sshKey           string
	sshKeyPassphrase string

	// oracle specific.
	traceFile string
	sslVerify string
	wallet    string

	// sql server.
	encrypt                string
	trustServerCertificate string
	connectionTimeout      string

	// keybindings.
	keybindings bool
)

// NewRootCmd returns the root command.
func NewRootCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// Default keybindings.

// If the --keybindings flag is set, fill the keybindings with the ones fonud in the config file.
// This is safe to do even if they're missing in the config files, because the config package has default values for it.

// Set the keybindings values, either the default ones or the found in the config file.

// rootCmd represents the base command when called without any subcommands.
var rootCmd = NewRootCmd()

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() { _ = "STUB: not implemented"; return }

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// config file flag.
	rootCmd.PersistentFlags().
		BoolVarP(&cfg, "config", "", false, "Get the connection data from a config file (default locations are: current directory, $HOME/.dblab.yaml or $XDG_CONFIG_HOME/.dblab.yaml)")

	// keybindings flag.
	rootCmd.PersistentFlags().
		BoolVarP(&keybindings, "keybindings", "k", false, "Get the keybindings configuration from the config file (default locations are: current directory, $HOME/.dblab.yaml or $XDG_CONFIG_HOME/.dblab.yaml)")

	// cfg-name is used to indicate the name of the config section to be used to establish a
	// connection with desired database.
	// default: if empty, the first item of the databases options is gonna be selected.
	rootCmd.Flags().StringVarP(&cfgName, "cfg-name", "", "", "Database config name section")

	// global flags used to open a database connection.
	rootCmd.Flags().StringVarP(&driver, "driver", "", "", "Database driver")
	rootCmd.Flags().StringVarP(&url, "url", "u", "", "Database connection string")
	rootCmd.Flags().StringVarP(&host, "host", "", "", "Server host name or IP")
	rootCmd.Flags().StringVarP(&port, "port", "", "", "Server port")
	rootCmd.Flags().StringVarP(&user, "user", "", "", "Database user")
	rootCmd.Flags().StringVarP(&pass, "pass", "", "", "Password for user")
	rootCmd.Flags().StringVarP(&db, "db", "", "", "Database name")
	rootCmd.Flags().
		StringVarP(&schema, "schema", "", "", "Database schema (optional for postgres and oracle only)")
	rootCmd.Flags().StringVarP(&ssl, "ssl", "", "", "SSL mode")
	rootCmd.Flags().
		UintVarP(&limit, "limit", "", 100, "Size of the result set from the table content query (should be greater than zero, otherwise the app will error out)")
	rootCmd.Flags().StringVarP(&socket, "socket", "", "", "Path to a Unix socket file")
	rootCmd.Flags().StringVarP(
		&sslcert,
		"sslcert",
		"",
		"",
		"This parameter specifies the file name of the client SSL certificate, replacing the default ~/.postgresql/postgresql.crt",
	)
	rootCmd.Flags().StringVarP(
		&sslkey,
		"sslkey",
		"",
		"",
		"This parameter specifies the location for the secret key used for the client certificate. It can either specify a file name that will be used instead of the default ~/.postgresql/postgresql.key, or it can specify a key obtained from an external “engine”",
	)
	rootCmd.Flags().
		StringVarP(&sslpassword, "sslpassword", "", "", "This parameter specifies the password for the secret key specified in sslkey")
	rootCmd.Flags().StringVarP(
		&sslrootcert,
		"sslrootcert",
		"",
		"",
		"This parameter specifies the name of a file containing SSL certificate authority (CA) certificate(s) The default is ~/.postgresql/root.crt",
	)
	rootCmd.Flags().
		StringVarP(&sslVerify, "ssl-verify", "", "", "[enable|disable] or [true|false] enable ssl verify for the server")
	rootCmd.Flags().StringVarP(&traceFile, "trace-file", "", "", "File name for trace log")
	rootCmd.Flags().StringVarP(&wallet, "wallet", "", "", "Path for auto-login oracle wallet")

	rootCmd.Flags().
		StringVarP(&encrypt, "encrypt", "", "", "[strict|disable|false|true] data sent between client and server is encrypted or not")
	rootCmd.Flags().
		StringVarP(&trustServerCertificate, "trust-server-certificate", "", "", "[false|true] server certificate is checked or not")
	rootCmd.Flags().
		StringVarP(&connectionTimeout, "timeout", "", "", "in seconds (default is 0 for no timeout), set to 0 for no timeout. Recommended to set to 0 and use context to manage query and connection timeouts")
	rootCmd.Flags().StringVarP(&sshHost, "ssh-host", "", "", "SSH Server Hostname/IP")
	rootCmd.Flags().StringVarP(&sshPort, "ssh-port", "", "", "SSH Port")
	rootCmd.Flags().StringVarP(&sshUser, "ssh-user", "", "", "SSH User")
	rootCmd.Flags().
		StringVarP(&sshPass, "ssh-pass", "", "", "SSH Password (Empty string for no password)")
	rootCmd.Flags().
		StringVarP(&sshKey, "ssh-key", "", "", "File with private key for SSH authentication")
	rootCmd.Flags().
		StringVarP(&sshKeyPassphrase, "ssh-key-pass", "", "", "Supports connections with protected private keys with passphrase")

	// rootCmd.Flags().
	// 	StringVarP(&sshKeyAlgo, "ssh-key-algo", "", "", "Publick Key Algorithm")
}
