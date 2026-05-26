package form

import (
	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"charm.land/lipgloss/v2/compat"
	"github.com/muesli/termenv"

	"github.com/danvergara/dblab/pkg/command"
)

const (
	defaultLimit = 100
)

var (
	isDark       = compat.HasDarkBackground
	focusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	noStyle      = lipgloss.NewStyle()
	term         = termenv.ColorProfile()
)

// Model is a meta-model.
type Model struct {
	// menu management.
	cursor int
	steps  int

	// driver.
	drivers []string
	driver  string

	// ssl connection params.
	sslCertInput     textinput.Model
	sslKeyInput      textinput.Model
	sslPasswordInput textinput.Model
	sslRootcertInput textinput.Model

	// oracle specific.
	traceFileInput textinput.Model
	sslVerifyInput textinput.Model
	walletInput    textinput.Model

	// sql server.
	trustServerCertificateInput textinput.Model

	// std data.
	hostInput     textinput.Model
	portInput     textinput.Model
	userInput     textinput.Model
	passwordInput textinput.Model
	databaseInput textinput.Model
	filePathInput textinput.Model
	limitInput    textinput.Model

	// ssl.
	postgreSQLSSLModes []string
	mySQLSSLModes      []string
	oracleSSLModes     []string
	sqlServerSSLModes  []string
	sqliteSSLModes     []string
	sslMode            string
}

// Init initialize the meta-model.
func (m *Model) Init() tea.Cmd {
	_ = "STUB: not implemented"
	return *

	// Update update the view of the meta-model.
	new(tea.Cmd)
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_ = "STUB: not implemented"
	// if the pressed keys are esc or ctrl + c, finish the execution.
	return *new(tea.Model), *new(tea.Cmd)
}

// View displays the content on the terminal.
func (m *Model) View() tea.View { _ = "STUB: not implemented"; return *new(tea.View) }

// Host returns the host value.
func (m *Model) Host() string { _ = "STUB: not implemented"; return "" }

// Port returns the Port value.
func (m *Model) Port() string { _ = "STUB: not implemented"; return "" }

// User returns the user value.
func (m *Model) User() string { _ = "STUB: not implemented"; return "" }

// Password returns the password value.
func (m *Model) Password() string { _ = "STUB: not implemented"; return "" }

// Database returns the database name value.
func (m *Model) Database() string { _ = "STUB: not implemented"; return "" }

// SSLMode returns the ssl mode name value.
func (m *Model) SSLMode() string { _ = "STUB: not implemented"; return "" }

func (m *Model) SSLCert() string { _ = "STUB: not implemented"; return "" }

func (m *Model) SSLKey() string { _ = "STUB: not implemented"; return "" }

func (m *Model) SSLPassword() string { _ = "STUB: not implemented"; return "" }

func (m *Model) SSLRootcert() string { _ = "STUB: not implemented"; return "" }

func (m *Model) SSLVerify() string { _ = "STUB: not implemented"; return "" }

func (m *Model) TraceFile() string { _ = "STUB: not implemented"; return "" }

func (m *Model) Wallet() string { _ = "STUB: not implemented"; return "" }

func (m *Model) TrustServerCertificate() string { _ = "STUB: not implemented"; return "" }

// Limit returns the limit input value from the user.
func (m *Model) Limit() (uint, error) {
	_ = "STUB: not implemented"
	// if the user skipped the question, resort to default value
	return 0, nil
}

// FilePath returns the path to the database file (just in sqlite3) value.
func (m *Model) FilePath() string { _ = "STUB: not implemented"; return "" }

func checkbox(label string, checked bool) string { _ = "STUB: not implemented"; return "" }

// Color a string's foreground with the given value.
func colorFg(val, color string) string { _ = "STUB: not implemented"; return "" }

func initModel() Model { _ = "STUB: not implemented"; return *new(Model) }

// the supported drivers by the client.

// our default value.

// Run runs the menus programs to introduced the required data to connect with a database.
func Run() (command.Options, error) { _ = "STUB: not implemented"; return *new(command.Options), nil }

// IsEmpty checks if the given options objects is empty.
func IsEmpty(opts command.Options) bool { _ = "STUB: not implemented"; return false }
