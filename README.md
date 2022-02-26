# airport-proxy

科学上网机场平台提供的订阅连接通常包含倍率节点，平台和客户端基本不提供屏蔽节点功能，导致自动切换状态下切换到倍率节点，消耗大量套餐流量。

## 原理
使用airport-proxy作为订阅代理，在`config.json`中添加过滤关键字，实现屏蔽功能。基本原理为：
* 获取请求中的身份信息
* 携带身份信息将请求转发到机场平台，解析机场响应
* 根据过滤关键字过滤机场响应
* 将过滤后的内容返回给客户端


一般机场会根据请求头的`User-Agent`判断请求客户端，来选择不同的响应格式。
例如使用 [ClashX](https://github.com/yichengchen/clashX) 客户端，机场返回yaml格式；除了yaml格式，还有直接返回节点信息并编码的格式。

## 使用
```shell
git clone https://github.com/ch0ngsheng/airport-proxy.git
cd airport-proxy
go build -o proxy main.go
```
### 修改配置文件
```shell
vim config.json
```
在`keywords`中添加需要屏蔽的关键字

### 启动代理
```shell
./proxy
```
默认监听http的9090端口，可使用`--port`参数自定义端口。
### 客户端添加配置
在客户端中添加托管配置，url格式为（已nicecloud机场为例）
```text
http://{IP}:9090/v1/ap/filter?ap=nicecloud&token={token}
```
其中token为机场提供的身份信息，可从机场官网的订阅链接中获取

## 支持平台
目前已验证的机场有：
* [Nice Cloud](http://nicecloud.me/#/register?code=iJkpNGUm)


目前已验证的客户端有：
* [ClashX](https://github.com/yichengchen/clashX)