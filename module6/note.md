# kube-apiserver

## 限流



1. APF(API Priority and Fairness)
- FlowSchema
    可以有相同name不同的distinguisherMethod(sa、？)
- PriorityLevelConfigurationq
    QueueSet(pool),针对每个FS限制一定数量的Queue。每个FS（相同name不同distinguisherMethod）可以分配到相同的Queue??