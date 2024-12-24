package rule

type RuleSetName string

type RuleSets struct {
	RuleSet []RuleSet `mapstructure:"rule_set"`
}

type RuleSet struct {
	Name     RuleSetName `mapstructure:"name"`
	RuleType RuleType    `mapstructure:"rule_type"`
}
