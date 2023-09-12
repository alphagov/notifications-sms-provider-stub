---
applications:
- name: notify-sms-provider-stub

  memory: 512M
  instances: 1

  buildpacks:
    - go_buildpack

  routes:
    - route: notify-sms-provider-stub-{{CF_SPACE}}.cloudapps.digital
    - route: notify-sms-provider-stub-{{CF_SPACE}}.apps.internal

  env:
    GOVERSION: go1.19
    MMG_CALLBACK_URL: {{API_PROTOCOL}}://{{API_HOSTNAME}}/notifications/sms/mmg
    FIRETEXT_CALLBACK_URL: {{API_PROTOCOL}}://{{API_HOSTNAME}}/notifications/sms/firetext
