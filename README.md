slack thread summarizer
===

Simple slack hello bot on Lambda.  
This is a practice to create slack bot written by Go lang and using Lambda.

## Architechture

- Lambda
- Go1.x

## Setup

https://docs.aws.amazon.com/ja_jp/lambda/latest/dg/golang-package.html

```sh
brew install gibo
gibo dump Go > .gitignore
go get github.com/aws/aws-lambda-go/lambda
```

## Setup Lambda

```sh
aws lambda create-function --function-name slack-thread-summarizer --runtime go1.x \
  --zip-file fileb://main.zip --handler main \
  --role arn:aws:iam::XXXXX:role/AWS_Lambda --profile XXXXX
```

## Build function

```sh
GOOS=linux GOARCH=amd64 go build main.go
zip main.zip ./main
```

## Upload function

```sh
aws lambda update-function-code --function-name slack-thread-summarizer --zip-file fileb://main.zip --profile XXX
```

## How to use Makefile

Deploy(delete zip file, build, zip, upload to lambda)

```sh
make deploy PROFILE="AWSPROFILENAME"
```
