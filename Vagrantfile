# -*- mode: ruby -*-
# vi: set ft=ruby :

$script = <<SCRIPT
echo "Installing content-server..."
curl -s https://packagecloud.io/install/repositories/foomo/daemon-garden/script.deb.sh | sudo bash
sudo apt-get install daemon-garden
SCRIPT

Vagrant.configure("2") do |config|
  config.vm.network :private_network, ip: "192.168.33.99"
  config.vm.network :forwarded_port, guest: 22, host: 2299

  config.vbguest.no_remote = true
  config.vbguest.auto_update = false

  config.vm.synced_folder ".", "/vagrant"

  config.vm.define 'trusty' do |instance|
    instance.vm.box = 'ubuntu/trusty64'
  end

  config.vm.define 'precise' do |instance|
    instance.vm.box = 'ubuntu/precise64'
  end

  config.vm.provision "shell", inline: $script

end
