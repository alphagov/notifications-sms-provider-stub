SHELL := /bin/bash

.PHONY: build
build:
	go build -o bin/sms-provider-stub

.PHONY: run
run: build
	./bin/sms-provider-stub

.PHONY: preview
preview:
	$(eval export CF_SPACE=preview)
	$(eval export API_HOSTNAME=notify.works)
	cf target -s ${CF_SPACE}

.PHONY: staging
staging:
	$(eval export CF_SPACE=staging)
	$(eval export API_HOSTNAME=staging-notify.works)
	cf target -s ${CF_SPACE}

.PHONY: generate-manifest
generate-manifest:
	$(if ${CF_SPACE},,$(error Must specify CF_SPACE))
	@sed -e "s/{{CF_SPACE}}/${CF_SPACE}/; s/{{API_HOSTNAME}}/${API_HOSTNAME}/" manifest.yml.tpl

.PHONY: cf-push
cf-push:
	$(if ${CF_SPACE},,$(error Must specify CF_SPACE))
	cf push -f <(make -s generate-manifest)
