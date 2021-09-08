##TodoList

startAt 6.3

- download
  - [x] download files     2 day

- prepare
  - [x] 拓扑信息终端展示    1.5 day
  - [x] 安装docker         1 day
  - [x] 安装ansible        1 day
  - [ ] k8s 集群配置检查
    - [ ] ntp 时间打印
    - [ ] ip段
  - [ ] 
- install
  - [x] k8s
  - [x] mongo
  - [x] ops
- uninstall
  - [x] k8s
  - [x] mongo
- web
  - [x] 编辑拓扑文件        4 day
  - [x] 实时拓扑状态展示     4 day
- add-node remove-node
  - [x] 增删节点

- update
  - [ ] koictl 自升级       2 day

- [x] 日志管理

- [x] download images from json
- [x] download charts latest
- [x] hare download charts latest
- [x] mongo connect 

7.12
- [x] registry chart select node
- [x] delete app ingress
- [x] set opsapp namespace
- [x] download bizcould
- [x] download all
- [x] untar on install
- [x] compress mongo k8s
- [x] ntp server
- [ ] exist mongo

- [x] test download all
- [x] test single node
- [x] test elk
- [x] test prometheus
- [x] test install all
- [x] test registry node
- [x] test hare
- [x] test ntp / storage
- [ ] mongo monitor
- [x] add metrics.image.
- [x] test ssh user/password 特殊字符 1234!@#$
- [x] download image error
- [x] health check k8s : failed pod  failed helm 
- [ ] test ops
    - [x] tenantmgr 使用private key ssh登录 mongodb和fileserver  
    - [ ] oss server
    - [x] mongo admin_password 固定 surpass_admin, bizconf_surpass 
    - [x] tenantmgr 重新部署hare时 不支持设置helm repo
- [x] multi elastic node
- [x] 拆分opsservice sgservice
- [x] download image after node add 
- [ ] skywalking
- [x] minio
  - [x] minio ingress
- [ ] https 证书
- [x] storageClass nfs
- [x] hostname prefix
- [x] 简化模板
    - [x] aliyun
      - [x] nfs改名
      - [x] 去掉registry/minio
      - [x] mongo init
    - [x] offline
      - [x] 去掉nfs
      - [x] 单节点mongo
- [x] 亲和性
- [ ] download bizcloudenterprise
- [x] ssh key登录
- [x] upgrade k8s 1.20.7
- [x] taint om node
    - [x] test minio 可用 mongo独占节点
    - [x] test om node taints
    - [x] test om 部署
 - [x] mongo test
 - [x] krew
    - [x] git