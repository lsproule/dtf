package rules

type RuleType struct{}

type REST_API struct {
	RuleType
	Method  string     `mapstructure:"method"`
	Headers HeaderList `mapstructure:"headers"`
}

type HeaderList struct {
	Method              string                 `mapstructure:"method"`
	Credentials         string                 `mapstructure:"credentials"`
	Headers             map[string]string      `mapstructure:"headers"`
	RequestStructure    map[string]string      `mapstructure:"request_structure"`
	ValidationStructure map[string]interface{} `mapstructure:"validation_structure"`
	Iterations          uint                   `mapstructure:"iterations"`
	DependsOn           []string               `mapstructure:"depends_on"`
}
