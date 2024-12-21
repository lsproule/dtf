package rules

type RuleType struct{}

type REST_API struct {
	RuleType
	Method  string
	Headers HeaderList
}

type HeaderList struct {
	Method              string
	Credentials         string
	Headers             map[string]string
	RequestStructure    map[string]string
	ValidationStructure map[string]string
	Iterations          uint
	DependsOn           []string
}
