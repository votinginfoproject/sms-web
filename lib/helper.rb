require 'net/ssh'

module Helper
  extend self

  def run_command(user, host, command)
    Net::SSH.start(host, user) do |ssh|
      channel = ssh.open_channel do |cha|
        puts "----------------------------------------------------"
        puts "running command: #{command}"
        puts "----------------------------------------------------"
        cha.exec(command) do |ch, success|
          raise "could not execute command" unless success

          ch.on_data do |c, data|
            STDOUT.write data
          end

          ch.on_extended_data do |c, type, data|
            STDERR.write data
          end

          ch.on_close do
            puts "----------------------------------------------------"
            puts "finished running command: #{command}"
            puts "----------------------------------------------------"
          end
        end
      end
    end
  end
end
