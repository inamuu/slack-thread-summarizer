
help: ## show help text
	@echo "[Usage]"
	@echo "make [help|deploy]"

deploy: ## uploda a zip file to lambda
	test -f main.zip && rm -f main.zip
	GOOS=linux GOARCH=amd64 go build main.go
	zip main.zip ./main
	aws lambda update-function-code --function-name slack-thread-summarizer --zip-file fileb://main.zip --profile ${PROFILE}
