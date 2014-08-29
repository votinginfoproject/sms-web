# SMS-Web

## Description
Receives SMSs from Twilio, pulls out the relevant data, and enqueues
that data in an AWS SQS queue.

## Requirements
- Golang 1.3
- [Godep](https://github.com/tools/godep)
- A .env file with the following items...

~~~~
ACCESS_KEY_ID=YOURACCESSKEY
SECRET_ACCESS_KEY=YOURSECRET
ENVIRONMENT=DEVELOPMENT
~~~~
