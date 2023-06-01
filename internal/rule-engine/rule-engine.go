package ruleengine

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
	"strings"
	"testAWS/internal/model"
	"testAWS/internal/utils/mailgun"

	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
)

type RuleEngine struct {
	DataContext   ast.IDataContext
	KnowledgeBase *ast.KnowledgeBase
}

func NewRuleEngine(rule *model.Rule, email *mailgun.MailgunEventParsed) (*RuleEngine, error) {
	ruleEngine := &RuleEngine{}
	emailJson, err := json.Marshal(email)
	if err != nil {
		return ruleEngine, err
	}

	// Load Email as a fact
	fact := []byte(emailJson)
	ruleEngine.DataContext = ast.NewDataContext()
	if err := ruleEngine.DataContext.AddJSON("Fact", fact); err != nil {
		return nil, err
	}

	//Create the rule
	rulesJSON, err := ioutil.ReadFile("rules.json")
	rulesJSONStr := string(rulesJSON)
	rulesJSONStr = strings.ReplaceAll(rulesJSONStr, "$operator", rule.Operator)
	rulesJSONStr = strings.ReplaceAll(rulesJSONStr, "$field", reflect.ValueOf(email).FieldByName(rule.Variable).String())
	rulesJSONStr = strings.ReplaceAll(rulesJSONStr, "$variable", rule.Variable)
	//rulesJSONStr = strings.ReplaceAll(rulesJSONStr, "$action", rule.Action)
	rulesJSON = []byte(rulesJSONStr)

	knowledgeLibrary := ast.NewKnowledgeLibrary()
	ruleBuilder := builder.NewRuleBuilder(knowledgeLibrary)
	if err := json.Unmarshal(rulesJSON, &ruleBuilder.KnowledgeLibrary); err != nil {
		panic(err)
	}

	bs := pkg.NewBytesResource([]byte(rulesJSON))
	if err := ruleBuilder.BuildRuleFromResource("EmailRules", "0.0.1", bs); err != nil {
		panic(err)
	}

	ruleEngine.KnowledgeBase = knowledgeLibrary.NewKnowledgeBaseInstance("EmailRules", "0.0.1")

	return ruleEngine, nil
}

func (ruleEngine *RuleEngine) ExecuteRule() error {
	engine := engine.NewGruleEngine()
	err := engine.Execute(ruleEngine.DataContext, ruleEngine.KnowledgeBase)
	if err != nil {
		return err
	}

	return nil
}
