# 模块八作业：

- 构建本地镜像，并将镜像推送至 Docker 官方镜像仓库
  
    ```
    make push tag=httpserver-v3
    ```

- 优雅启动：[查看代码](https://github.com/hzhhong/cncamp/blob/main/module8/httpserver/httpserver.yaml#L64-L75)
    - `startupProbe` : 当容器创建后开始检测，每隔10s检测一次，调大 failureThreshold=30  
    - `readinessProbe` : startupProbe 成功后开始检测，每隔5s检测一次，2次成功后ready

- 探活：[查看代码](https://github.com/hzhhong/cncamp/blob/main/module8/httpserver/httpserver.yaml#L57-L63)
- 资源需求和QoS保证：[查看代码](https://github.com/hzhhong/cncamp/blob/main/module8/httpserver/httpserver.yaml#L39-L45)
- 优雅终止：[查看代码](https://github.com/hzhhong/gap/blob/v0.0.1/app.go#L67-L104)
    - 监听进程退出信号
    - 结合 [errorgroup](https://pkg.go.dev/golang.org/x/sync@v0.0.0-20210220032951-036812b2e83c/errgroup?utm_source=gopls) 实现多Server优雅退出
- 日志等级：[查看代码](https://github.com/hzhhong/gap/blob/v0.0.1/log/log.go#L22-L32)
    
    ```
    INFO method=getCfgValue server.http1.name=httpserver1
    INFO method=getCfgValue server.http1.addr=0.0.0.0:8081
    INFO method=getCfgValue server.http2.name=httpserver2
    INFO method=getCfgValue server.http2.addr=0.0.0.0:8082
    2021/11/21 02:03:31 App Started
    INFO msg=[HTTP] server [httpserver1] listening on: 0.0.0.0:8081
    INFO msg=[HTTP] server [httpserver2] listening on: 0.0.0.0:8082
    INFO TimeStamp=2021-11-21T02:03:36+08:00 server=httpserver1 clientIp=[::1]:63586 path=/readinesshealthz1 statuscode=200 latency=0
    INFO TimeStamp=2021-11-21T02:03:36+08:00 server=httpserver1 clientIp=[::1]:63586 path=/favicon.ico statuscode=200 latency=0
    INFO TimeStamp=2021-11-21T02:03:50+08:00 server=httpserver1 clientIp=[::1]:63586 path=/a statuscode=200 latency=0
    INFO TimeStamp=2021-11-21T02:03:51+08:00 server=httpserver1 clientIp=[::1]:63586 path=/favicon.ico statuscode=200 latency=0
    INFO TimeStamp=2021-11-21T02:04:03+08:00 server=httpserver1 clientIp=[::1]:63586 path=/a statuscode=200 latency=0
    INFO TimeStamp=2021-11-21T02:04:03+08:00 server=httpserver1 clientIp=[::1]:63586 path=/favicon.ico statuscode=200 latency=7.39e-05
    INFO TimeStamp=2021-11-21T02:04:12+08:00 server=httpserver1 clientIp=[::1]:63586 path=/a statuscode=200 latency=0
    INFO TimeStamp=2021-11-21T02:04:12+08:00 server=httpserver1 clientIp=[::1]:63586 path=/favicon.ico statuscode=200 latency=0.0005307
    INFO msg=Server [httpserver2] Exited Properly
    INFO msg=Server [httpserver1] Exited Properly
    2021/11/21 02:07:22 App Exited Properly
    ```
- 配置和代码分离
    - 读取配置代码：[点击查看](https://github.com/hzhhong/cncamp/blob/main/module8/httpserver/main.go#L49-L61)
    - configMap：[点击查看](https://github.com/hzhhong/cncamp/blob/main/module8/httpserver/httpserver.yaml#L1-L13)
    - 挂载配置：[点击查看1](https://github.com/hzhhong/cncamp/blob/main/module8/httpserver/httpserver.yaml#L77-L83)、[点击查看2](https://github.com/hzhhong/cncamp/blob/main/module8/httpserver/httpserver.yaml#L46-L49)
    
- 创建Sevice供集群内使用 [点击查看](https://github.com/hzhhong/cncamp/blob/main/module8/httpserver/httpserver.yaml#L85-L97)
  - 调整 deployment 的 replicas=3 保证高可用
- 配置Ingress供集群外使用 [点击查看](https://github.com/hzhhong/cncamp/blob/main/module8/httpserver/httpserver.yaml#L99-L129)
  
  
  
  