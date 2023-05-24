package model

type Rule struct {
	Domain   string
	Field    string `json:"field"`
	Operator string `json:"operator"`
	Variable string `json:"variable"`
	Action   string `json:"action"`
}

type RuleI interface {
}
