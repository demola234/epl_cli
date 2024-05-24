package ui

// Config contains TUI-specific configuration.
type Config struct {
	ShowAllFiles    bool
	Gopath          string `env:"GOPATH"`
	HomeDir         string `env:"HOME"`
	GlamourMaxWidth uint
	GlamourStyle    string
	EnableMouse     bool

	
}
