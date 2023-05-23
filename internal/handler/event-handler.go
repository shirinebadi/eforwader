package handler

import (
	"context"
	"net/http"
	"strings"
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

	mailgunEventParsed, err := mailgunEventParser.MailgunEventParser(body)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest}, err
	}

	return events.APIGatewayProxyResponse{StatusCode: http.StatusOK, Body: string("from: " + mailgunEventParsed.From + "subject: " + mailgunEventParsed.Subject + "body: " + mailgunEventParsed.BodyPlain)}, nil

}
