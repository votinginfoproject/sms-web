# SMS-Web

## Description
Receives SMSs from Twilio, pulls out the relevant data, and enqueues
that data in an AWS SQS queue.

## Requirements

- [Docker][docker]

**OR**

- Golang 1.3
- Ruby 2.1.2
- [Godep](https://github.com/tools/godep)
- A .env file with the following items...
    - The first three items in the example .env file below (AWS credentials and
      environment) MUST be in that order at the top of the file.

~~~~
ACCESS_KEY_ID=
SECRET_ACCESS_KEY=
ENVIRONMENT=development
QUEUE_PREFIX=vip-sms-app
PROCS=24
LOGGLY_TOKEN=
NEWRELIC_TOKEN=
TWILIO_SID=
~~~~

## Development System Setup
To set your system up to develop this application...

1. Make sure you have everything from the requirements section
2. Run `bundle`

## Docker Development System Setup

It is potentially easier to develop with [Docker][docker].

To compile and run the project, the typical docker `build` and `run`
commands will work. When running, you will need to have the
environment variables above set, because the docker version does not
use the .env file.

With the environment variables set, the commands are:

```
$ docker build -t sms-web .
$ docker run -p 8080:8080 sms-web
```

[docker]: https://www.docker.com/

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
