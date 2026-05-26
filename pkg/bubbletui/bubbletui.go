package bubbletui

import (
	"io"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"

	"github.com/danvergara/dblab/pkg/client"
	"github.com/danvergara/dblab/pkg/command"
)

type focusState int

var (
	// colors.
	green      = lipgloss.Color("#1fb009") // Normal green
	purple     = lipgloss.Color("#800080")
	cyberGreen = lipgloss.Color("#39FF14") // High-visibility neon green
	hiMagenta  = lipgloss.Color("#FF00FF") // High-visibility Magenta
	mutedGreen = lipgloss.Color("#2ECC71") // Softer green for standard text
	neonPurple = lipgloss.Color("#BF40BF") // Bright purple for highlights
	darkPurple = lipgloss.Color("#4B0082") // Deep violet for backgrounds
	whiteText  = lipgloss.Color("#E0E0E0") // Off-white for readability
	black      = lipgloss.Color("#000000")
)

const (
	// focus state management.
	focusEditor focusState = iota
	focusList
	focusTable
)

var (
	baseStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			Padding(0, 1)

	tablesListStyle = baseStyle
	editorStyle     = baseStyle
	resultSetStyle  = baseStyle

	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(green).
			Foreground(purple).
			AlignVertical(lipgloss.Center).
			Align(lipgloss.Center)

	footerStyle = lipgloss.NewStyle().
			Foreground(green)

	errorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF0000")).
			Bold(true).
			Padding(1, 2)
)

// metadataSucessMsg struct used to retrieve a given table's metadata asynchronously.
type metadataSuccessMsg struct {
	metadata *client.Metadata
}

// metadataErrMsg struct used to report error to user at the time to retrieve metadata.
type metadataErrMsg struct{ err error }

// querySuccessMsg struct used to get result sets from executed queries asynchronously.
// Sometimes, tables can be created, altered of deleted, so the this returns a refreshed list of tables.
type querySuccessMsg struct {
	columns []string
	rows    [][]string
	tables  []string
}

// queryErrMsg struct used to report when the query execution fails.
type queryErrMsg struct{ err error }

type Model struct {
	// database client.
	c *client.Client

	// models.
	editor          Editor
	sidebarViewport SidebarViewport
	resulstset      ResultSet

	// Manages the focus on the app.
	focus focusState

	// widget dimensions.
	width                 int
	height                int
	leftWidth             int
	rightWidth            int
	titleHeight           int
	titleWidth            int
	sidebarViewportHeight int
	sidebarViewportWidth  int
	resultSetHeight       int
	resultSetWidth        int
	editorHeight          int
	editorWidth           int

	// Key Bindings.
	bindings *command.TUIKeyMap

	// constant text on the client.
	footer        string
	renderedTitle string

	dump io.Writer
}

func NewModel(c *client.Client, kb *command.TUIKeyMap) (*Model, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m Model) Init() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(tea.Model), *new(tea.Cmd)
}

func (m Model) View() tea.View { _ = "STUB: not implemented"; return *new(tea.View) }

func (m *Model) Run() error { _ = "STUB: not implemented"; return nil }

// runTableMetadata gets the given table's metadata asynchronously.
// If the query succeeds, it returns metadataSucessMsg with the metadata, otherwise it returns metadataErrMsg with the error.
func (m *Model) runTableMetadata(table client.TableRef) tea.Cmd {
	_ = "STUB: not implemented"
	return *new(tea.Cmd)
}

// executeQueryCmd method executes queryes asynchronously, so it does not block the bubbletea execution.
// If it succeeds, returns a querySuccessMsg with the resultset. Otherwise, it returns queryErrMsg with the error.
func (m *Model) executeQueryCmd(query string) tea.Cmd {
	_ = "STUB: not implemented"
	return *new(tea.Cmd)
}
