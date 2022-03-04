#使用说明：

1. export TAG 变量，通常作为代码发布 版本
export tag=1.0
2. 登录docker hub
docker login -u <username>
3. 进行打包推送镜像
make push
