package manager

import (
	"log/slog"

	"github.com/lsproule/dtf/src/rules"
)

type StateManager struct {
	Logger             *slog.Logger
	RuleSetDefinitions []rules.RuleSet
	TestQueue          chan rules.RuleSet
	ValidationQueue    chan rules.RuleSet
}
