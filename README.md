<h1 align="center">
<img src="https://nekobox-public.oss-cn-hangzhou.aliyuncs.com/images/Neko.png" width=100px/>

狼的提问箱
</h1>

<p align="center">
匿名提问箱 / Anonymous Question Box

## 截图

### QRcode
![QRCode](./dev/qrcode.png#gh-light-mode-only)

### Commit with picture
![CWP](./dev/commitwithpic.png#gh-light-mode-only)

## 安装

### 需求

* [Go](https://golang.org/dl/) (v1.19 或更高版本)
* [MySQL](https://www.mysql.com/downloads/) (v5.7 或更高版本)
* [Redis](https://redis.io/download/) (v6.0 或更高版本)

### 从源码编译

```bash
git clone https://github.com/tamakyi/TamaBox.git

cd TamaBox

go build -o TamaBox
```

### 编辑配置文件

```bash
cp conf/app.sample.ini conf/app.ini
```

### 运行

```bash
./TamaBox web
```

## 架构

![Architecture](./dev/nekobox-arch-light.png#gh-light-mode-only)
![Architecture](./dev/nekobox-arch-dark.png#gh-dark-mode-only)

NekoBox 使用 GitHub Actions 进行持续集成和部署。

当用户访问 NekoBox 时，请求将会被发送至阿里云 CDN，CDN 的访问日志将会被实时推送到阿里云日志服务
(SLS)。日志数据将在 SLS 中存储 180 天，用于审计。

用户的信息、提问和回答将被存储在 MySQL 数据库中。

用户的会话、CSRF 令牌和电子邮件验证令牌将被暂时存储在 Redis 中。

用户的整个请求和响应链路将被上传到 Uptrace 用于调试。这些数据将被储存 30 天。管理员可以使用用户提供的 `TraceID`
来追踪查询指定的请求。

当用户提交提问时，问题的内容将被发送到七牛文本审查服务进行审查。如果提问内容存在问题，该内容将被发送到阿里云文本审查服务进行二次审查。
如果内容审查仍未通过，该提问将被拒绝发送。这是由于七牛文本审查服务不是很准确，一些非冒犯性的内容可能被七牛文本审查误报。

当用户收到新的提问时，阿里云邮件服务（DM）会向用户的邮箱发送一封邮件。

你可以在主页查看 NekoBox 的更新日志，也欢迎访问赞助页面来打钱支持 NekoBox。 更新日志和赞助商名单存储在独立部署的
Pocketbase 服务中。

## 开源协议

MIT License

## 备注

* 基于原版[Nekobox](https://github.com/NekoWheel/NekoBox)修改，基本与原版无异，主要是添加了个人主页二维码和首页统计数，并修改了部分适合自己的参数。