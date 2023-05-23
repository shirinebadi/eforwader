package main

import (
	"context"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type MailgunEvent struct {
	BodyPlain string   `json:"body-plain"`
	From      string   `json:"from"`
	Subject   string   `json:"subject"`
	To        []string `json:"to"`
}

func HandleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	contentType := req.Headers["content-type"]
	boundary := strings.Split(contentType, "=")
	body := req.Body

	log.Println("Content Type", contentType)
	log.Println("Boundry", boundary[1])
	log.Println("Body", req.Body)
	log.Println("UGHHHH")

	reader := multipart.NewReader(strings.NewReader(body), boundary[1])

	// Iterate over each part in the multipart form
	for {
		part, err := reader.NextPart()
		if err != nil {
			log.Println(err)
			break
		}

		partName := part.FormName()

		// Process the desired parts of the email
		switch partName {
		case "Subject":
			subjectBytes, _ := ioutil.ReadAll(part)
			log.Println("Subject:", string(subjectBytes))
		case "From":
			fromBytes, _ := ioutil.ReadAll(part)
			log.Println("From:", string(fromBytes))
		case "To":
			toBytes, _ := ioutil.ReadAll(part)
			log.Println("To:", string(toBytes))
		case "body-plain":
			bodyBytes, _ := ioutil.ReadAll(part)
			log.Println("Body:", string(bodyBytes))

		}
	}

	return events.APIGatewayProxyResponse{StatusCode: http.StatusOK}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
