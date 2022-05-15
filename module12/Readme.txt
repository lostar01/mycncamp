
```
 k create ns httpserver
 kubectl create -f Httpserver-Deployment.yaml -n httpserver
 kubectl create -f Httpserver-service.yaml -n httpserver
 openssl req -x509 -sha256 -nodes -days 365 -newkey rsa:2048 -subj '/O=lostar Inc./CN=*.lostar.com' -keyout httpserver.key -out httpserver.crt
 kubectl create -n istio-system secret tls httpserver --key=httpserver.key --cert=httpserver.crt
 kubectl create -f Httpserver-Gateway.yaml -n httpserver

 #启动open tracing
 1. 安装配置 jaeger
   kubectl create -f jaeger.yaml
 2. 给目的的 ns 加入service_mesh 客户端
   kubectl label ns tracing istio-injection=enabled
 3. 让相关的POD 重新生成
 4. 打开jaeger dashboard 查看
   istioctl dashboard --address 0.0.0.0 jaeger

```
