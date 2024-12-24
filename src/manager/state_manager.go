package manager

import (
	"log/slog"

	"github.com/lsproule/dtf/src/rule"
	"github.com/lsproule/dtf/src/validation"
)

type StateManager struct {
	Logger             *slog.Logger
	RuleSetDefinitions []rule.RuleSet
	TestQueue          chan rule.RuleSet
	ValidationQueue    chan validation.ValidationSet
	ValidationSet      chan validation.ValidationSet
}

/*
Process:
- Read in the .yaml
- Parse the necessary information
  - Create rulesets
- Generate what will be used for the tests based off of the rulesets
- Perform tests
- Perform validation
- Store the validation (success or error) into the validation set
*/
