package main

import (
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

func hello() (string, error) {
	log.Printf("Hello World2")
	return "Hello2", nil
}

func main() {
	lambda.Start(hello)
}
