# Install etcd

## 安装 `cfssl`

``` bash
apt-get install golang-cfssl
```

## 生成证书

使用 [Makefile](./tls-setup/Makefile) 生成证书
```
make all infra0=etcd0 infra1=etcd1 infra2=etcd2
```

## 单节点

## HA

使用 [Makefile](./Makefile) 部署高可用集群
```
make all
```

## 使用实践

- cluster operator
- lease
- snapshot
- alarm
- defrag