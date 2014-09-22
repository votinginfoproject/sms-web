require 'aws-sdk'
require 'dotenv/tasks'
require 'net/http'
require 'uri'
require_relative 'lib/helper'

desc 'Build and deploy sms-web'
task :deploy, [:environment] => :aws_auth do |_, args|
  environment = args.environment or
    fail 'You must specify an environment type (development, staging, or production): `rake deploy[ENVIRONMENT]`'

  environment.downcase!
  fail 'please supply a valid environment value (development, staging, or production)' unless ['development', 'staging', 'production'].include?(environment)

  puts 'building...'
  system('CGO_ENABLED=0 GOOS=linux GOARCH=amd64 godep go build -o sms-web sms-web.go')

  env = IO.read('.env').split("\n")[3..-1].join("\n") # remove first 3 lines
  s3 = AWS::S3.new

  puts 'uploading...'
  s3.buckets["vip-sms-#{environment}"].objects['sms-web'].write(file: 'sms-web')
  s3.buckets["vip-sms-#{environment}"].objects['sms-web-env'].write(env)

  system('rm sms-web')

  unless environment == 'development'
    tag = case environment
      when 'staging' then 'vip-sms-app-staging-web'
      when 'production' then 'vip-sms-app-web'
    end

    puts 'restarting service on instances...'
    ec2 = AWS::EC2.new
    ec2.instances.with_tag('Name', tag).each do |instance|
      next if instance.status == :terminated
      Helper.run_command('ubuntu', instance.public_ip_address, 'sudo service sms-web stop')
      Helper.run_command('ubuntu', instance.public_ip_address, 'sudo service sms-web start')
    end
  end
end

desc 'Send test request'
task :test, [:environment, :number, :message] => :aws_auth do |_, args|
  number = args.number or
    fail "You must specify a phone number: `rake test[ENVIRONMENT, NUMBER, 'MESSAGE']`"

  number =~ (/\+\d{11}$/) or
    fail "Phone number must be in the format '+15555551234' `rake test[ENVIRONMENT, NUMBER, 'MESSAGE']`"

  environment = args.environment or
    fail "You must specify an environment type (staging, or production): `rake test[ENVIRONMENT, NUMBER, 'MESSAGE']`"

  message = args.message or
    fail "You must specify a message: `rake test[ENVIRONMENT, NUMBER, 'MESSAGE']`"

  dns = case environment
        when 'production'
          elb = AWS::ELB.new
          elb.load_balancers['vip-sms-app-lb1'].dns_name
        when 'staging'
          ec2 = AWS::EC2.new
          ec2.instances.with_tag('Name', 'vip-sms-app-staging-web').collect.first.dns_name
        end

  uri = URI.parse("http://#{dns}")

  data = {
    'From' => number,
    'Body' => message,
    'AccountSid' => ENV['TWILIO_SID']
  }

  response = Net::HTTP.post_form(uri, data)
  puts "Response code: #{response.code}"
end

desc 'AWS auth config'
task :aws_auth => :dotenv do
  AWS.config(
    :access_key_id => ENV['ACCESS_KEY_ID'],
    :secret_access_key => ENV['SECRET_ACCESS_KEY']
  )
end
