package sshdb

import (
	"database/sql/driver"
	"net"
	"os"
	"path/filepath"
	"time"

	"golang.org/x/crypto/ssh"
)

// default path to the known_hosts file.
var defaultKnownHostsPath = filepath.Join(os.Getenv("HOME"), ".ssh")

// createKnownHosts function creates known_hosts if does not exist.
// It uses the os package which has an OpenFile function, this function accepts 3 arguments:
// 1. the file path
// 2. the flag (e.g. os.O_CREATE|os.O_APPEND creates the file if not exists, if exists, appends to the file)
// 3. the last argument is the permission.
func createKnownHosts(knownHostsPath string) (err error) { _ = "STUB: not implemented"; return nil }

// checkKnownHosts fucntion creates a know_hosts callback function with the New function.
// This callback function can be used to check if the host exists in the known_hosts file.
func checkKnownHosts(knownHostsPath string) (ssh.HostKeyCallback, error) {
	_ = "STUB: not implemented"
	return *new(ssh.HostKeyCallback), nil
}

// keyString create human-readable SSH-key strings.
func keyString(k ssh.PublicKey) string { _ = "STUB: not implemented"; return "" }

// e.g. "ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTY...."

// addHostKey adds the host key to known_hosts file by using Normalize and Line functions of knownhosts package.
// This functions implements the ssh.HostKeyCallback type wiich is a function type which signature goes like this:
// type HostKeyCallback func(hostname string, remote net.Addr, key PublicKey) error.
func addHostKey(_ string, remote net.Addr, pubKey ssh.PublicKey, knownHostsPath string) error {
	_ = "STUB: not implemented"
	return nil
}

// PostgresViaSSHDialer implements the driver.Driver interface to register the connection to the database via the ssh tunnel.
type PostgresViaSSHDialer struct {
	client *ssh.Client
}

func (sd *PostgresViaSSHDialer) Open(s string) (_ driver.Conn, err error) {
	_ = "STUB: not implemented"
	return *new(driver.Conn), nil
}

func (sd *PostgresViaSSHDialer) Dial(network, address string) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

func (sd *PostgresViaSSHDialer) DialTimeout(
	network, address string,
	timeout time.Duration,
) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

// MySQLViaSSHDialer used to register the database connection via the ssh tunnel.
type MySQLViaSSHDialer struct {
	client *ssh.Client
}

func (m *MySQLViaSSHDialer) Dial(addr string) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

// SSHConfig struct setup the ssh tunnel to connect with a given database.
type SSHConfig struct {
	sshUser        string
	sshPass        string
	sshKeyFile     string
	sshKeyPass     string
	sshHost        string
	sshPort        string
	sshClient      *ssh.Client
	dbDriver       string
	dbURL          string
	knownHostsPath string
}

type Option func(*SSHConfig)

func New(opts ...Option) *SSHConfig { _ = "STUB: not implemented"; return nil }

func WithSSHUser(sshUser string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithPass(sshPass string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithSSHKeyFile(sshKeyFile string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithSSHKeyPass(sshKeyPass string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithSShHost(sshHost string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithSShPort(sshPort string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithDBDriver(driver string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithDBDURL(url string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithKnownHostsPath(knownHostsPath string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// SSHTunnel method sets up the ssh tunnel and does a number of things:
// Create a ssh client config object that witht he user.
// Define a HostKeyCallback to ensures known ssh server is the actual server.
// If host key checking is ignore then any server that has the same FQDN or IP address can impersonate the actual ssh server.
// Define the authentication method to perform the ssh tunnel (passsword or private key).
// Register the ViaSSHDialer with the ssh connection as a parameter.
func (c *SSHConfig) SSHTunnel() error {
	_ = "STUB: not implemented"
	// Reference: https://github.com/melbahja/goph/blob/master/client.go
	// Reference: https://github.com/melbahja/goph/blob/master/hosts.go
	// Study the client.go and hosts.go to understand how to write host key call back.
	return nil
}

// Reference: https://www.godoc.org/golang.org/x/crypto/ssh/knownhosts#KeyError
// if keyErr.Want slice is empty then host is unknown, if keyErr.Want is not empty
// and if host is known then there is key mismatch the connection is then rejected.

// host key not found in known_hosts then give a warning and continue to connect.

// Load the private key for SSH authentication.

// Parse the private using a passphrase if required.

// Close method closes the tcp connection.
func (c *SSHConfig) Close() error { _ = "STUB: not implemented"; return nil }
