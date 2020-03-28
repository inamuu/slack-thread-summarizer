
package main

import (
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

func hello() (string, error) {
	log.Printf("Hello World")
	return "Hello", nil
}

func main() {
	lambda.Start(hello)
}
