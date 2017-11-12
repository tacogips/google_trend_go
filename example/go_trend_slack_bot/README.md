## Slack Google Hot trend stories push bot on GAE Example

## Install and set up

`go get -u github.com/tacogips/google_trend_go/example/go_trend_slack_bot`

this example work on google app engine with go1.8.

[check google app engine tutorial](https://cloud.google.com/appengine/docs/standard/go/runtime)

## Basic Usage

All you need to edit is env_variables section in app/app.yaml

1. Create your own Slack Bot user and API Token

2. Write Slack token to `SLACK_TOKEN` in `app/app.yaml`

3. Create channel(s) to push "hot trend" and "newest stories" (it's can be a single channel ).
   Then Invite Bot User to the channels.

4. Write channel names to `HOT_TREND_CHANNEL` and `STORIES_CHANNEL` in  `app/app.yaml`

5. `make serve` to work on local app engine . `make deploy` to deploy.

##  Change Geolocation ,Language and Timezone
Geo, Language, and timezone to pass to google trend api set as Japanese(Asia/Tokyo).
To change geo and timezone in `hot-trend.go` `stories.go`

