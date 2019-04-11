.PHONY: build
build:
	go build -o bin/sms-provider-stub

run: build
	./bin/sms-provider-stub
