slack_bot_hellp
===

## Architechture

- Lambda
- Go1.x

## Setup

- https://docs.aws.amazon.com/ja_jp/lambda/latest/dg/golang-package.html

```sh
brew install gibo
gibo dump Go > .gitignore
go get github.com/aws/aws-lambda-go/lambda
```

Setup Lambda
```sh
aws lambda create-function --function-name slack-bot-hello --runtime go1.x \
  --zip-file fileb://hello.zip --handler main \
  --role arn:aws:iam::XXXXX:role/AWS_Lambda --profile XXXXX

```

## Build function

```sh
GOOS=linux GOARCH=amd64 go build main.go
zip main.zip ./main
```

## Upload function

```sh
aws lambda update-function-code --function-name slack-bot-hello --zip-file fileb://main.zip --profile XXX
```
