## gin-http3
一个gin使用HTTP/3的案例。

同时启动了向下兼容的，HTTP/2+TLS，和HTTP。

证书和key，使用mkcert完成创建，后拷贝到项目目录的。

gin作为HTTP服务的Handler使用。

如果只需要测试HTTP3，则仅需要`http3.ListenAndServeQUIC`部分即可。

运行：

```shell
go mod tidy
go run .
```

```go
gin-http3> go run .
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /                         --> main.main.func1 (3 handlers)

```



使用curl访问测试：

HTTP3，curl --http3 选项指示发出H3请求。需要新版的curl支持。同时需要指定证书。

```go
>curl --http3 --cacert ./localhost.pem https://localhost/
Hello HTTP/3.0
```



HTTP2，仅需要指定证书即可。此时curl发出H2请求：

```go
>curl --cacert ./localhost.pem https://localhost/
Hello HTTP/2.0
```



HTTP/1，此处使用http协议演示。

```shell
>curl http://localhost/
Hello HTTP/1.1
```



以上三个例子对比，我们的服务程序，同时支持了H3，H2，H1三种协议，其中H2和H1是通用的，优先使用了H2，都支持http和https，而H3仅仅支持https。

```shell
>curl --http3 http://localhost/
Hello HTTP/1.1
```