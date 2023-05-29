package model

type Rule struct {
	Field    string `json:"field"`
	Operator string `json:"operator"`
	Variable string `json:"variable"`
	Action   string `json:"action"`
}

type Hits struct {
	Hits []struct {
		Source Rule `json:"_source"`
	} `json:"hits"`
}

type Response struct {
	Hits Hits `json:"hits"`
}

type RuleI interface {
}
