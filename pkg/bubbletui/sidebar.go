package bubbletui

import (
	"context"
	"io"

	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"github.com/Digital-Shane/treeview/v2"

	"github.com/danvergara/dblab/pkg/client"
	"github.com/danvergara/dblab/pkg/command"
)

func dbObjectHasType(nodeType string) func(*treeview.Node[*client.DBNode]) bool {
	_ = "STUB: not implemented"
	return nil
}

type selectTableMsg struct {
	Schema string
	Table  string
}

type SidebarViewport struct {
	c        *client.Client
	bindings *command.TUIKeyMap

	sidebarViewport viewport.Model
	dbTree          *treeview.TuiTreeModel[*client.DBNode]
	width, height   int

	selected bool
	dump     io.Writer
}

type DBGraphTreeBuilderProvider struct{}

func (d DBGraphTreeBuilderProvider) ID(do *client.DBNode) string {
	_ = "STUB: not implemented"
	return ""
}

func (d *DBGraphTreeBuilderProvider) Name(do *client.DBNode) string {
	_ = "STUB: not implemented"
	return ""
}

func (p *DBGraphTreeBuilderProvider) Children(do *client.DBNode) []*client.DBNode {
	_ = "STUB: not implemented"
	return nil
}

func NewSidebarViewport(ctx context.Context, c *client.Client, kb *command.TUIKeyMap) (SidebarViewport, error) {
	_ = "STUB: not implemented"
	return *new(SidebarViewport), nil
}

func (s *SidebarViewport) SetSize(w, h int) { _ = "STUB: not implemented"; return }

func (s *SidebarViewport) newTuiTreeModel(tree *treeview.Tree[*client.DBNode], width, height int) *treeview.TuiTreeModel[*client.DBNode] {
	_ = "STUB: not implemented"
	// Create custom key map to avoid key conflicts
	return nil
}

func (s SidebarViewport) Init() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (s SidebarViewport) Update(msg tea.Msg) (SidebarViewport, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(SidebarViewport), *new(tea.Cmd)
}

func (s SidebarViewport) View() string { _ = "STUB: not implemented"; return "" }

func createCyberpunkProvider() *treeview.DefaultNodeProvider[*client.DBNode] {
	_ = "STUB: not implemented"
	// Icons for database objects.
	return nil
}
