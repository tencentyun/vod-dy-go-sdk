# 简介

欢迎使用腾讯云开发者工具套件（SDK），此 SDK 是云 API 3.0 平台的配套开发工具。

# 依赖环境

1. Go 1.9 版本及以上（如使用 go mod 需要 Go 1.14）。
2. 在腾讯云控制台 [访问管理](https://console.cloud.tencent.com/cam/capi) 页面获取密钥 SecretID 和 SecretKey，请务必妥善保管，或者使用更安全的临时安全凭证。

# 获取安装

## 通过go get安装（推荐）

推荐使用腾讯云镜像加速下载：

1. Linux 或 MacOS:

    ```bash
    export GOPROXY=https://mirrors.tencent.com/go/
    ```

2. Windows:

    ```cmd
    set GOPROXY=https://mirrors.tencent.com/go/
    ```

### 安装步骤

注意：此安装方式仅支持使用 **Go Modules** 模式进行依赖管理，即环境变量 `GO111MODULE=auto`或者`GO111MODULE=on`, 并且在您的项目中执行了 `go mod init xxx`.

1. 安装公共基础包：

    ```bash
    go get -v -u github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common
    ```

2. 安装 vod 产品包:

    ```bash
    go get -v -u github.com/tencentyun/vod-dy-go-sdk
    ```

# 相关配置

**如无特殊需要，建议您使用默认配置。**

在创建客户端前，如有需要，您可以通过修改`profile.ClientProfile`中字段的值进行一些配置。

```go
// 非必要步骤
// 实例化一个客户端配置对象，可以指定超时时间等配置
cpf := profile.NewClientProfile()
```

具体的配置项说明如下：

## 请求方式

SDK默认使用POST方法。 如果你一定要使用GET方法，可以在这里设置。**GET方法无法处理一些较大的请求**。

```go
cpf.HttpProfile.ReqMethod = "POST"
```

## 超时时间

SDK有默认的超时时间，如非必要请不要修改默认设置。
如有需要请在代码中查阅以获取最新的默认值。  
单位：秒

```go
cpf.HttpProfile.ReqTimeout = 30
```

## 签名方式

SDK默认用 `TC3-HMAC-SHA256` 进行签名，它更安全但是会轻微降低性能。

```go
cpf.SignMethod = "HmacSHA1"
```

## DEBUG模式

DEBUG模式会打印更详细的日志，当您需要进行详细的排查错误时可以开启。  
默认为 `false`

```go
cpf.Debug = true
```

## 禁用长连接(Keep-alive)

SDK 的每一个 client 默认使用长连接模式，即请求的头部 Connection 字段的值为 keep-alive.

如果您需要使用短连接，可以按照以下方式进行设置：

```go
...
    client, _ := vod.NewClient(credential, regions.Guangzhou, cpf)
    tp := &http.Transport{
        DisableKeepAlives: true,
    }
    client.WithHttpTransport(tp)
...
```

则此 client 发出的每个请求头部的 Connection 字段的值为 close

## 地域容灾

从 `v1.0.227`开始，腾讯云 GO SDK 支持地域容灾功能：

当请求满足以下条件时：

1. 失败次数 >= 5 次
2. 失败率 >= 75%

SDK 会自动将您请求的地域设置为备选地域。

相关设置如下：

```golang
    // 开启
    cpf.DisableRegionBreaker = false
    // 设置备用请求地址，不需要指定服务，SDK 会自动在头部加上服务名(如cvm)
    // 例如，设置为 ap-guangzhou.tencentcloudapi.com，则最终的请求为 vod.ap-guangzhou.tencentcloudapi.com
    cpf.BackupEndpoint = "ap-guangzhou.tencentcloudapi.com"
```

此功能仅支持单个客户端的同步请求。

## 代理

如果是有代理的环境下，需要设置系统环境变量 `https_proxy` ，否则可能无法正常调用，抛出连接超时的异常。或者自定义 Transport 指定代理，通过 client.WithHttpTransport 覆盖默认配置。

## 开启 DNS 缓存

当前 GO SDK 总是会去请求 DNS 服务器，而没有使用到 nscd 的缓存，可以通过导出环境变量`GODEBUG=netdns=cgo`，或者`go build`编译时指定参数`-tags 'netcgo'`控制读取 nscd 缓存。

## 忽略服务器证书校验

虽然使用 SDK 调用公有云服务时，必须校验服务器证书，以识破他人伪装的服务器，确保请求的安全。
但是某些极端情况下，例如测试时，你可能会需要忽略自签名的服务器证书。
以下是其中一种可能的方法：

```golang
import "crypto/tls"
...
    client, _ := vod.NewClient(credential, regions.Guangzhou, cpf)
    tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }
    client.WithHttpTransport(tr)
...
```

**再次强调，除非你知道自己在做什么，并明白由此带来的风险，否则不要尝试关闭服务器证书校验。**
