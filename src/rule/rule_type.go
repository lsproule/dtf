package rule

type RuleType struct{}

type Server string

// Closely follow the OpenAPI standard
type REST_API struct {
	RuleType
	Servers []string `mapstructure:"servers"`
	Method  string   `mapstructure:"method"`
	//Headers HeaderList `mapstructure:"headers"`
}

type Path struct {
	PathSegment string
	Method      Method
}

// type HttpMethod string

// const (
// 	Get HttpMethod = "GET"
// )

type Method struct {
	Method      string
	Description string
	OperationId string
	Parameters  Parameters
}

type Parameters struct {
	Name        string
	In          string
	Description string
	Required    bool
}
