package app

import (
	"github.com/danvergara/dblab/pkg/bubbletui"
	"github.com/danvergara/dblab/pkg/client"
	"github.com/danvergara/dblab/pkg/command"
	"github.com/danvergara/dblab/pkg/sshdb"
)

// App Struct.
type App struct {
	c  *client.Client
	sc *sshdb.SSHConfig
	m  *bubbletui.Model
}

// New bootstrap a new application.
func New(opts command.Options) (*App, error) { _ = "STUB: not implemented"; return nil, nil }

// Run runs the application.
func (a *App) Run() error { _ = "STUB: not implemented"; return nil }
