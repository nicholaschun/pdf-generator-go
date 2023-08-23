package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Request events.APIGatewayProxyRequest
type Response events.APIGatewayProxyResponse

func generatePdf(request Request)(Response, error){
	// implement pdf generation logic
	fmt.Println("Received body", request.Body)
	return Response{Body: request.Body, StatusCode: 200}, nil
}

func main(){
	lambda.Start(generatePdf)
}