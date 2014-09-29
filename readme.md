# SMS-Web

## Description
Receives SMSs from Twilio, pulls out the relevant data, and enqueues
that data in an AWS SQS queue.

## Requirements
- Golang 1.3
- Ruby 2.1.2
- [Godep](https://github.com/tools/godep)
- A .env file with the following items...
    - The first three items in the example .env file below (AWS credentials and
      environment) MUST be in that order at the top of the file.

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

## Commands
### Run Tests
~~~~
godep go test ./...
~~~~

### Deploy
~~~~
rake deploy\[environment\]
~~~~

- Build the binary
- Upload the binary to S3
- Upload all but the first THREE lines of the .env file to S3
- Restart the sms-web process on all instances

### Send Test Message
~~~~
rake test\[environment,number,message\]
~~~~

- Send a test SMS from the specified number
