K8S集群部署说明
--

##### 前提说明
1. koictl命令可以对k8s-offline-install进行离线文件整理,简化部署操作,
   同样也导致K8S集群的配置信息被隐藏.
2. k8s-offline-install项目是基于开源项目[kubespray](https://github.com/kubernetes-sigs/kubespray)开发的.
   kubespray项目支持的操作都适用于k8s-offline-install,举例
   - 通过ansible进行安装K8S,在koictl根路径:
   ```
   $ ansible-playbook -i inventory/hosts.yaml k8s-offline-install/kubespray/prepare.yml -v
   $ ansible-playbook -i inventory/hosts.yaml k8s-offline-install/kubespray/cluster-offline.yml -v
   ```

##### K8S配置说明

1. K8S集群相关的配置在inventory/group_vars和k8s-offline-install/kubespray/roles中各模块下.
   可以通过编辑inventory/group_vars进行调整配置参数

##### 组件启动方式
| 组件 | 启动方式 | 配置文件
--- | --- | ---
kube-apiserver | docker | |
controller-manager|docker||
kubelet| systemctl | |
etcd | docker | |


##### 默认配置

ingress-nginx-controller:

|配置项 | 值 |
--- | ---
|域名 | www.test.com
|http 端口| hostPort 80
|https端口| hostPort 443


服务网址:

| | 网址 | 登录
--- | --- | ---
kibana | http://www.test.com/kibana | kibana/Bizconf1
grafana | http://www.test.com/grafana | admin/admin
alertmanager | http://www.test.com/alertmanager |
prometheus | http://www.test.com/prometheus |

数据路径：

| | 存储模式 | 路径
--- | --- | ---
| kube conf | local path | /etc/kubernetes|
| etcd | local path |master节点/var/lib/etcd |


mongodb:

| 配置项 | 值 |
--- | ---
|mongo 版本 | v4.0.23
| rs0端口 | 30003
| rs0 admin 用户名| admin
|rs0 admin 密码 | super_admin_123
|suprass_ops 用户名 | suprass
|suprass_ops 密码 | surpass_private
|存储路径 | /mongodb

