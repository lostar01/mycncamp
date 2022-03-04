#使用说明：

#### 1. export TAG 变量，通常作为代码发布 版本
export tag=1.0
#### 2. 登录docker hub
docker login -u <username>
#### 3. 进行打包推送镜像
make push
#### 4. 在任何地方都可以跑httpserver
docker run -d --name test -p 8080:8080 lostar01/mycncamp-httpserver:1.0
#### 5. 查看容器IP配置
  root@kali:~# lsns -t net
        NS TYPE NPROCS   PID USER     NETNSID NSFS                           COMMAND
4026531992 net     238     1 root  unassigned                                /sbin/init
4026532319 net       1   550 root  unassigned                                /usr/sbin/haveged --Foreground --verbose=1 -w 1024
4026532387 net       1 16788 root           0 /run/docker/netns/766b5d6a1f7a /app/http_server
4026532441 net       1   741 rtkit unassigned                                /usr/lib/rtkit/rtkit-daemon
root@kali:~# nsenter -t 16788 -n ip add
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
9: eth0@if10: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default 
    link/ether 02:42:ac:11:00:02 brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet 172.17.0.2/16 brd 172.17.255.255 scope global eth0
       valid_lft forever preferred_lft forever
