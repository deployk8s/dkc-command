#### 说明

dkc项目旨在centos 7.9.2009上使用kubespray(版本2.16.0)剧本离线安装kubernetes v1.20.7

dkc-command是dkc项目中的命令行工具, 提供下载离线文件,安装kubernetes等功能

#### 使用方式

1. 下载离线文件,
```shell script
./dkc-command download cache
```
2. 编辑拓扑文件,放在inventory目录

3. 安装kubernetes
```shell script
./dkc-command install k8s
```
4. 增加节点
```shell script
./dkc-command node add --hostname <hostname> --ip <ip>
```
5. 删除节点
```shell script
./dkc-command node del --hostname <hostname> 
```
6. 检查kubernetes
```shell script
./dkc-command status k8s
```

7. 删除kubernetes
```shell script
./dkc-command uninstall k8s
```