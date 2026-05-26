package form

import (
	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
)

func updateDriver(msg tea.Msg, m *Model) (tea.Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *

	// Is it a key press?
	new(tea.Model), *new(tea.Cmd)
}

// the "up" and "k" keys mve the cursor up.

// the "down" and "j" keys move the cursor down.

func updateStd(msg tea.Msg, m *Model) (tea.Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(tea.Model), *new(tea.Cmd)
}

// Set focused state.

// Remove focused state.

func stdInputs(m *Model) []textinput.Model { _ = "STUB: not implemented"; return nil }

func assignStdInputValues(m *Model, inputs []textinput.Model) { _ = "STUB: not implemented"; return }

func updateInputs(msg tea.Msg, m *Model) (*Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return nil, *new(tea.Cmd)
}

func updateSSLMode(msg tea.Msg, m *Model) (tea.Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *

	// Is it a key press?
	new(tea.Model), *new(tea.Cmd)
}

// These keys should exit the program.
// the "up" and "k" keys mve the cursor up.

// the "down" and "j" keys move the cursor down.

func updateSSLConn(msg tea.Msg, m *Model) (tea.Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(tea.Model), *new(tea.Cmd)
}

// Set focused state.

// Remove focused state.

func sslConnInputs(m *Model) []textinput.Model { _ = "STUB: not implemented"; return nil }

func assignSSLConnInputValues(m *Model, inputs []textinput.Model) {
	_ = "STUB: not implemented"
	return
}

func updateSSLConnInputs(msg tea.Msg, m *Model) (*Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return nil, *new(tea.Cmd)
}
