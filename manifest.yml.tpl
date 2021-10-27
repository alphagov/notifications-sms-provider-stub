---
applications:
- name: notify-sms-provider-stub

  memory: 256M
  instances: 1

  buildpacks:
    - go_buildpack

  routes:
    - route: notify-sms-provider-stub-{{CF_SPACE}}.cloudapps.digital

  env:
    GOVERSION: go1.12
    MMG_CALLBACK_URL: https://{{API_HOSTNAME}}/notifications/sms/mmg
    FIRETEXT_CALLBACK_URL: https://{{API_HOSTNAME}}/notifications/sms/firetext
