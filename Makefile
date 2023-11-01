SHELL := /bin/bash
CF_MANIFEST_PATH ?= /tmp/manifest.yml


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

.PHONY: preview
preview:
	$(eval export CF_SPACE=preview)
	$(eval export API_HOSTNAME=api.notify.works)
	cf target -s ${CF_SPACE}

.PHONY: staging
staging:
	$(eval export CF_SPACE=staging)
	$(eval export API_HOSTNAME=api.staging-notify.works)
	cf target -s ${CF_SPACE}

.PHONY: generate-manifest
generate-manifest:
	$(if ${CF_SPACE},,$(error Must specify CF_SPACE))
	@sed -e "s/{{CF_SPACE}}/${CF_SPACE}/; s/{{API_HOSTNAME}}/${API_HOSTNAME}/" manifest.yml.tpl

.PHONY: cf-push
cf-push:
	$(if ${CF_SPACE},,$(error Must specify CF_SPACE))
	make -s generate-manifest > ${CF_MANIFEST_PATH}
	cf push -f ${CF_MANIFEST_PATH}
