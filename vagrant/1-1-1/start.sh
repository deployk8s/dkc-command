
curl -o ./dkc.tar.gz https://github.com/deployKubernetesInCHINA/dkc-command/releases/download/release/dkc.tar.gz

vagrant up
vagrant upload dkc.tar.gz /root/ vagrant-master-1
vagrant ssh vagrant-master-1 -c "tar xvf dkc.tar.gz"
vagrant upload hosts.yaml /root/dkc/inventory/ vagrant-master-1
vagrant ssh vagrant-master-1
