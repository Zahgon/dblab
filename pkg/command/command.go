package command

import (
	"charm.land/bubbles/v2/key"
)

// Options is a struct that stores the provided commands by the user.
type Options struct {
	Driver string
	URL    string
	Host   string
	Port   string
	User   string
	Pass   string
	DBName string
	// PostgreSQL only.
	Schema string
	Limit  uint
	Socket string
	SSL    string
	// SSH.
	SSHHost          string
	SSHPort          string
	SSHUser          string
	SSHPass          string
	SSHKeyFile       string
	SSHKeyPassphrase string
	// SSL connection params.
	SSLCert     string
	SSLKey      string
	SSLPassword string
	SSLRootcert string
	// oracle specific.
	TraceFile string
	SSLVerify string
	Wallet    string
	// sql server.
	Encrypt                string
	TrustServerCertificate string
	ConnectionTimeout      string
	// TUI keybidings.
	TUIKeyBindings TUIKeyMap
}

// UpdateKeybindings method updates the TUIKeyBindings field, since the keybidings configuration parted ways with the connection configuration.
func (o *Options) UpdateKeybindings(k TUIKeyMap) { _ = "STUB: not implemented"; return }

type TUIKeyMap struct {
	NextTab         key.Binding
	PrevTab         key.Binding
	PageTop         key.Binding
	PageBottom      key.Binding
	EndOfLine       key.Binding
	BeginningOfLine key.Binding
	Navigation      TUINavigationKeyMap
	Editor          EditorKeyMap
}

type EditorKeyMap struct {
	// Normal Mode Navigation.
	Up    key.Binding
	Down  key.Binding
	Left  key.Binding
	Right key.Binding

	// Mode Switching.
	Insert key.Binding
	Normal key.Binding

	// Actions.
	ExecuteQuery key.Binding
}

type TUINavigationKeyMap struct {
	Up    key.Binding
	Down  key.Binding
	Left  key.Binding
	Right key.Binding
}

func DefaultKeyMap() TUIKeyMap { _ = "STUB: not implemented"; return *new(TUIKeyMap) }

// Capital 'G' for shift+g

// --- Mode Switching ---

// --- Actions ---

// SetDefault returns a Options struct and fills the empty
// values with environment variables if any.
func SetDefault(opts Options) Options { _ = "STUB: not implemented"; return *new(Options) }
