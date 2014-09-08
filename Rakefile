require 'aws-sdk'
require 'dotenv/tasks'
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
      Helper.run_command('ubuntu', instance.public_ip_address, 'sudo service sms-web stop')
      Helper.run_command('ubuntu', instance.public_ip_address, 'sudo service sms-web start')
    end
  end
end

desc 'AWS auth config'
task :aws_auth => :dotenv do
  AWS.config(
    :access_key_id => ENV['ACCESS_KEY_ID'],
    :secret_access_key => ENV['SECRET_ACCESS_KEY']
  )
end
