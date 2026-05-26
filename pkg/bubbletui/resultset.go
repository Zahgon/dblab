package bubbletui

import (
	"io"

	"charm.land/bubbles/v2/table"
	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/danvergara/dblab/pkg/client"
	"github.com/danvergara/dblab/pkg/command"
)

// tabStyles is for tab styling.
// The tabs are used to show table metadata.
type tabStyles struct {
	inactiveTab lipgloss.Style
	activeTab   lipgloss.Style
}

// newTabStyles function retuns a pointer to the tabStyles.
// It basically defines the default borders for bot active and inactive tabs.
func newTabStyles() *tabStyles { _ = "STUB: not implemented"; return nil }

type ResultSet struct {
	focused       bool
	tabs          []string
	activeTab     int
	width, height int
	tabStyles     *tabStyles

	bindings *command.TUIKeyMap

	viewport       viewport.Model
	tablesMetadata []table.Model
	dump           io.Writer
}

func NewResultSet(kb *command.TUIKeyMap) ResultSet {
	_ = "STUB: not implemented"
	return *new(ResultSet)
}

func (r *ResultSet) Focus() { _ = "STUB: not implemented"; return }

func (r *ResultSet) Blur() { _ = "STUB: not implemented"; return }

func (r *ResultSet) SetSize(w, h int) { _ = "STUB: not implemented"; return }

func (r *ResultSet) setupTable() { _ = "STUB: not implemented"; return }

func (r ResultSet) Init() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (r ResultSet) Update(msg tea.Msg) (ResultSet, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(ResultSet), *new(tea.Cmd)
}

func (r ResultSet) View() tea.View { _ = "STUB: not implemented"; return *new(tea.View) }

func (r *ResultSet) clearTables() { _ = "STUB: not implemented"; return }

// updateTableMetadataOnChange method is used to print the table metadata retrieved asynchronously.
func (r *ResultSet) updateTableMetadataOnChange(metadata *client.Metadata) {
	_ = "STUB: not implemented"
	return
}

// table data.

// table columns.

// table indexes.

// table constraints.

// tabBorderWithBottom function is used to define the tab borders.
// Borders changes whether the tabs is inacative or inactive.
// Active tab misses the bottom border.
func tabBorderWithBottom(left, middle, right string) lipgloss.Border {
	_ = "STUB: not implemented"
	return *new(lipgloss.Border)
}

// prepare method sets up the client defaults, such as the tables, the editor, the initial queries to show the either the databases or tables the user has access to and the styles.
func setupTable(height, width int) table.Model { _ = "STUB: not implemented"; return *new(table.Model) }

func populateTable(headers []string, data [][]string) ([]table.Column, []table.Row) {
	_ = "STUB: not implemented"
	return nil, nil
}
