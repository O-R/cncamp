扩缩容 横弹 纵弹
回收request请求的资源

锁死cpu资源 cpuset,绑核

numa node?cpu拓扑

bestEffort requests  
Burstable requests   request < limit
Guaranteed requests  request = limit

应用炸弹怎么快速定位到是哪个pod出现问题


gang scheduling : 全成功或全失败的调度机制

statefulset: 顺序号 1、2、3增长，3、2、1删除

k8s不做顺序依赖，B->A。A、B一起启动，B失败直到A成功

有状态服务：需要保证拓扑状态和存储状态


20：29 deamonset statefulset 不用replicatset实现多副本管理，controllerrevision 

20:45 什么场景需要 delete --cascade=orphan

20:53 lease 实现 kube-scheduler controller-manager 高可用

20:56 课间
21:01 kubelet
todo:cgroup深入学习

21:19 pause 
sandbox container (pause): sleep infinity，没有资源开销、稳定。
作用：挂载网络配置、

cri csi kubelet 启动顺序？

21:33 cri

11/11
