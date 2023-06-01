package handler

import (
	"context"
	"log"
	"net/http"
	"strings"
	elasticsearch "testAWS/connection/elastic-search"
	"testAWS/internal/config"
	ruleengine "testAWS/internal/rule-engine"
	"testAWS/internal/utils/mailgun"

	"github.com/aws/aws-lambda-go/events"
)

type EventHandler struct {
}

func (h *EventHandler) HandleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	boundary := strings.Split(req.Headers["content-type"], "=")[1]
	body := req.Body

	mailgunEventParser := &mailgun.MailgunEventParser{
		Boundary: boundary,
	}

	elasticConfig := config.ElasticConfig{
		Hosts:    "https://search-rulestest-w3xr76s5lrzjqwsicj66oywwyy.eu-central-1.es.amazonaws.com/",
		Username: "shirine",
		Password: "Sh123456irin!",
	}

	es := &elasticsearch.ElasticSearchClient{
		Config: elasticConfig,
	}

	mailgunEventParsed, err := mailgunEventParser.MailgunEventParser(body)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest}, err
	}

	log.Println(mailgunEventParsed)

	rule, err := es.QueryRule("rule")
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest}, err
	}

	log.Println(rule)

	ruleEngine, err := ruleengine.NewRuleEngine(rule, mailgunEventParsed)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest}, err
	}

	err = ruleEngine.ExecuteRule()
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest}, err
	}

	return events.APIGatewayProxyResponse{StatusCode: http.StatusOK}, nil
}
