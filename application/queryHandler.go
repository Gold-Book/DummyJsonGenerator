package main

import (
 "log"
 "bytes"
 "encoding/json"
 "github.com/rs/xid"
 "github.com/aws/aws-lambda-go/events"
 "github.com/aws/aws-lambda-go/lambda"
 "github.com/aws/aws-sdk-go/aws"
 "github.com/aws/aws-sdk-go/aws/endpoints"
 "github.com/aws/aws-sdk-go/aws/session"
 "github.com/aws/aws-sdk-go/service/s3"
)

const bucket = aws.String("test")
const endpoint = "http://localstack:4572"

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
 log.Printf("Processing Lambda request %s\n", request.RequestContext.RequestID)

 var jsonData interface{}
 if err := json.Unmarshal([]byte(request.Body), &jsonData); err != nil {
     log.Print(err)
     return response(400, "requets body is not in json format.", nil)
 }

  if err := putS3(request.Body); err != nil {
      log.Print(err)
      response(500, "Failed to register request.", nil)
  }

 return response(200, "{ }", nil)
}

func putS3(data string) (error) {
    s := session.Must(session.NewSession(&aws.Config {
        S3ForcePathStyle: aws.Bool(true),
        Region: aws.String(endpoints.UsWest2RegionID),
        Endpoint: aws.String(endpoint),
    }))

     c := s3.New(s, &aws.Config{})

    putData := s3.PutObjectInput {
        Bucket: aws.String(bucket),
        Key: aws.String(xid.New().String()),
        ACL: aws.String("public-read"),
        Body: bytes.NewReader([]byte(data)),
        ContentType: aws.String("text/plain"),
        ContentLength: aws.Int64(int64(len([]byte(data)))),
    }

    _, err := c.PutObject(&putData)
    return err
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
