SHELL := /bin/bash
CF_MANIFEST_PATH ?= /tmp/manifest.yml
API_PROTOCOL ?= https

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
	$(eval export API_HOSTNAME=notify-api-staging.apps.internal:8080)
	$(eval export API_PROTOCOL=http)
	cf target -s ${CF_SPACE}

.PHONY: generate-manifest
generate-manifest:
	$(if ${CF_SPACE},,$(error Must specify CF_SPACE))
	@sed -e "s/{{CF_SPACE}}/${CF_SPACE}/; s/{{API_HOSTNAME}}/${API_HOSTNAME}/; s/{{API_PROTOCOL}}/${API_PROTOCOL}/" manifest.yml.tpl

.PHONY: cf-push
cf-push:
	$(if ${CF_SPACE},,$(error Must specify CF_SPACE))
	make -s generate-manifest > ${CF_MANIFEST_PATH}
	cf push -f ${CF_MANIFEST_PATH}
