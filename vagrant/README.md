### vagrant 测试环境

1. 安装virtualbox
2. virtualbox 创建natnetwork. Name随意, 网段192.168.56.0/24
3. 安装vagrant
4. 准备工作目录
    ```shell script
    mkdir centos7
    cd centos7
    cp 1-1-1/Vagrantfile .
    ```
5. 创建虚拟机
    ```shell script
    vagrant up
    ```

6. 准备dkc

    ```shell script
    vagrant ssh vagrant-master-1
    
    #进入vagrant-master-1之后
    curl -o /tmp/dkc.tar.gz https://github.com/deployKubernetesInCHINA/dkc-command/releases/download/release/dkc.tar.gz
    tar xf /tmp/dkc.tar.gz -C $(pwd)
    cd dkc
    
    #将1-1-1/hosts.yaml 复制到inventory目录
    #scp root@192.168.56.1:/path-to-1-1-1/hosts.yaml inventory/
    ```

7. 安装

    ```shell script
    ./dkc download cache
    ./dkc install k8s -y
    ```