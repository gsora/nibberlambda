# `nibberlambda` - a Telegram bot to write like a ni🅱️🅱️a, leveraging AWS Lambda

This branch uses the official AWS Lambda Go SDK!

## Usage

To build `nibberlambda`, make sure to `go get github.com/aws/aws-lambda-go/lambda` first, then `go build` in this directory, `zip nibberlambda.zip nibberlambda`.

Next, create a new Lambda function, upload `nibberlambda.zip` and setup an appropriate POST endpoint using API Gateway.

`nibberlambda` requires two environment variables to be set:

 - `POST_URL`, which is the URL you should have got from API Gateway
 - `BOT_KEY`, your Telegram bot API key

This is mostly an experiment, but you can find a working instance of this bot at `@nibber_bot`!
