package main

import (
	"testAWS/internal/handler"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	eventHandler := &handler.EventHandler{}
	lambda.Start(eventHandler.HandleRequest)
}
