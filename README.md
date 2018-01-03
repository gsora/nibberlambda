# `nibberlambda` - a Telegram bot to write like a ni🅱️🅱️a, leveraging AWS Lambda

## Usage

To build `nibberlambda`, make sure to setup your environment as written [here](https://github.com/eawsy/aws-lambda-go-shim#quick-hands-on) - **Docker is required**.

Then create a new Lambda function, upload `handler.zip` and setup an appropriate POST endpoint using API Gateway.

`nibberlambda` requires two environment variables to be set:

 - `POST_URL`, which is the URL you should have got from API Gateway
 - `BOT_KEY`, your Telegram bot API key

This is mostly an experiment, but you can find a working instance of this bot at `@nibber_bot`!
