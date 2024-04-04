SHELL := /bin/bash

.PHONY: build-with-docker
build-with-docker:
	docker build -f docker/Dockerfile -t sms-provider-stub .

.PHONY: run-with-docker
run-with-docker: build-with-docker
	docker run \
		-p 6300:6300 \
		-e FIRETEXT_CALLBACK_URL="http://host.docker.internal:6011/notifications/sms/firetext" \
		-e MMG_CALLBACK_URL="http://host.docker.internal:6011/notifications/sms/mmg" \
		sms-provider-stub

.PHONY: build
build:
	go build -o bin/sms-provider-stub

.PHONY: run
run: build
	./bin/sms-provider-stub
