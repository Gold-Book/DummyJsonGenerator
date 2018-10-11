package main

import (
 "log"
 "bytes"
 "encoding/json"
 "github.com/aws/aws-lambda-go/events"
 "github.com/aws/aws-lambda-go/lambda"
 "github.com/aws/aws-sdk-go/aws"
 "github.com/aws/aws-sdk-go/aws/endpoints"
 "github.com/aws/aws-sdk-go/aws/session"
 "github.com/aws/aws-sdk-go/service/s3"
 "github.com/aws/aws-sdk-go/aws/credentials"
 // "github.com/koron/go-dproxy"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
 log.Printf("Processing Lambda request %s\n", request.RequestContext.RequestID)

 var jsonData interface{}
 if err := json.Unmarshal([]byte(request.Body), &jsonData); err != nil {
     log.Print(err)
     return response(400, "requets body is not in json format.", nil)
 }

 s := session.Must(session.NewSession(&aws.Config{
     Credentials:      credentials.NewStaticCredentials("foo", "var", ""),
     S3ForcePathStyle: aws.Bool(true),
     Region:           aws.String(endpoints.UsWest2RegionID),
     Endpoint:         aws.String("http://localstack:4572"),
 }))

 c := s3.New(s, &aws.Config{})

 p := s3.PutObjectInput{
     Bucket: aws.String("test"),
     Key:    aws.String("key.txt"),
     ACL:    aws.String("public-read"),
     Body: bytes.NewReader([]byte("test")),
 }

 _, err := c.PutObject(&p)
 if err != nil {
     log.Print(err)
     return response(500, "Failed to register request.", nil)
 }

 return response(200, "{ }", nil)
}

func response(code int, message string, err error) (events.APIGatewayProxyResponse, error) {
    return events.APIGatewayProxyResponse {
     Body: message,
     StatusCode: code,
    }, err
}

func main() {
 lambda.Start(Handler)
}
