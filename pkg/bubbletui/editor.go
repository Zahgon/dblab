package bubbletui

import (
	"io"

	"charm.land/bubbles/v2/textarea"
	tea "charm.land/bubbletea/v2"
	"github.com/danvergara/dblab/pkg/command"
)

type Mode int

const (
	NormalMode Mode = iota
	InsertMode
)

type executeQueryMsg struct {
	Query string
}

type Editor struct {
	editor     textarea.Model
	bindings   *command.TUIKeyMap
	mode       Mode
	register   string
	pendingCmd string
	dump       io.Writer
}

func NewEditor(kb *command.TUIKeyMap) Editor { _ = "STUB: not implemented"; return *new(Editor) }

func (e *Editor) SetWidth(w int) { _ = "STUB: not implemented"; return }

func (e *Editor) SetHeight(h int) { _ = "STUB: not implemented"; return }

func (e *Editor) Blur() { _ = "STUB: not implemented"; return }

func (e *Editor) Focus() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (e Editor) Init() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (e Editor) Update(msg tea.Msg) (Editor, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(Editor), *new(tea.Cmd)
}

func (e Editor) View() tea.View { _ = "STUB: not implemented"; return *new(tea.View) }

func (e *Editor) yankCurrentLine() { _ = "STUB: not implemented"; return }

func (e *Editor) deleteCurrentLine() { _ = "STUB: not implemented"; return }

func (e *Editor) pasteAfter() { _ = "STUB: not implemented"; return }
