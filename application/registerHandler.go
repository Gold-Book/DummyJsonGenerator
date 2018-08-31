package main

import (
 "log"
 "github.com/aws/aws-lambda-go/events"
 "github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

 log.Printf("Processing Lambda request %s\n", request.RequestContext.RequestID)

 return events.APIGatewayProxyResponse{
  Body:       "Hello " + request.Body,
  Headers:    map[string]string{ "x-custom-header" : "my custom header value" },
  StatusCode: 200,
 }, nil

}

func main() {
 lambda.Start(Handler)
}
