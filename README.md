# notifications-sms-provider-stub

This is a simple Go server that can be used as a stub for Notify SMS providers during a load test or when running the application locally. It accepts the requests for a notification, prints the message contents, returns an HTTP response and after a short delay sends a callback to a pre-configured URL.

## Setting up

### Install go

You need a Go compiler to build the binary (`brew install go`).

### `environment.sh` (optional)

You can tweak the configuration using the following variables.

```shell
echo "
export PORT=6300  # server port

export MMG_MIN_DELAY_MS=100  # min delay before callback is sent in ms
export MMG_MAX_DELAY_MS=1000 # max delay before callback is sent in ms
export MMG_CALLBACK_URL='http://localhost:6011/notifications/sms/mmg'

# similarly for each other provider
"> environment.sh
```

## To run the application

To build and run the server locally:

```shell
make run
```

This will start a server on port 6300, configured to send the callbacks to a local Notify API. To configure Notify API to use the server instead of actual MMG and Firetext set the `environment.sh` variables in the Notify API:

```shell
export MMG_URL='http://localhost:6300/mmg'

# similarly for each other provider
```

## To deploy the application

### How to make the API use this email provider stub

To turn it on for an app running in the PaaS use:

```
cf set-env APP-NAME FIRETEXT_URL http://notify-sms-provider-stub-staging.apps.internal:8080/firetext
cf set-env APP-NAME MMG_URL http://notify-sms-provider-stub-staging.apps.internal:8080/mmg
cf restage APP-NAME
```

and equivalent for other providers. The environment variables will remain set even if you redeploy the app.
