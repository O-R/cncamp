# 1009模块三作业：

- 构建本地镜像

    ```
    docker build -t huangzhihong/cncamp:httpserver-v1 .
    ```

- 编写 Dockerfile 将练习 2.2 编写的 httpserver 容器化（请思考有哪些最佳实践可以引入到 Dockerfile 中来）

    > [Dockerfile](./Dockerfile)
    1. 使用 Multi-stage build 减少层数
    2. 合并Run命令 减少层数
   

- 将镜像推送至 Docker 官方镜像仓库
  
    ```
    docker push huangzhihong/cncamp:httpserver-v1
    ```

- 通过 Docker 命令本地启动 httpserver
    
    ```
    docker run -d --name httpserver -p 80:8080 huangzhihong/cncamp:httpserver-v1
    ```

- 通过 nsenter 进入容器查看 IP 配置
  
  ```
  nsenter -t $(docker inspect -f "{{.State.Pid}}" httpserver) -n ip a
  ```