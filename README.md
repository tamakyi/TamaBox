<h1 align="center">
<img src="https://nekobox-public.oss-cn-hangzhou.aliyuncs.com/images/Neko.png" width=100px/>

狼的提问箱
</h1>

<p align="center">
匿名提问箱 / Anonymous Question Box

## 说明
* 基于原版[Nekobox](https://github.com/NekoWheel/NekoBox)修改，基本与原版无异，请支持原作者。
* 修改的地方在于添加了几个~~加了没加都不影响体验的~~js，形成了**二维码分享**和引入waline并屏蔽了部分框架的**态度面板**。具体如截图所示。

## 截图

### QRcode
![QRCode](./dev/erweima.png#gh-light-mode-only)

### 发表态度
![CWP](./dev/taidu.png#gh-light-mode-only)

## 安装

### 需求

* [Go](https://golang.org/dl/) (v1.19 或更高版本)
* [MySQL](https://www.mysql.com/downloads/) (v5.7 或更高版本)
* [Redis](https://redis.io/download/) (v6.0 或更高版本)

### 从源码编译

```bash
git clone -b dev-qrcode  https://github.com/tamakyi/TamaBox.git

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

## 开源协议

MIT License

