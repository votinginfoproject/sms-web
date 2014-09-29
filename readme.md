# SMS-Web

## Description
Receives SMSs from Twilio, pulls out the relevant data, and enqueues
that data in an AWS SQS queue.

## Requirements
- Golang 1.3
- [Godep](https://github.com/tools/godep)
- A .env file with the following items...

~~~~
ACCESS_KEY_ID=
SECRET_ACCESS_KEY=
ENVIRONMENT=DEVELOPMENT
QUEUE_PREFIX=vip-sms-app
PROCS=24
LOGGLY_TOKEN=
NEWRELIC_TOKEN=
TWILIO_SID=
~~~~
