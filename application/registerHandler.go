package main

import (
 "log"
 "github.com/aws/aws-lambda-go/events"
 "github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

 log.Printf("Processing Lambda request %s\n", request.RequestContext.RequestID)

 return events.APIGatewayProxyResponse {
  Body:       "Hello " + request.Body,
  StatusCode: 200,
 }, error

}

func main() {
 lambda.Start(Handler)
}
