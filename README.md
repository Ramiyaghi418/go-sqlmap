# Go-Sqlmap

## 介绍

- sqlmap是渗透测试界的神器，于是想尝试写一个简单的sqlmap
- 使用Golang重写的原因：高效、生成可执行文件直接运行无需环境等
- 目前测试在sqlilab的前两关可以成功，主要是针对mysql的回显注入功能，其他注入后续加入

## 快速开始

直接在github的release页面下载可执行文件：[下载地址](https://github.com/EmYiQing/go-sqlmap/releases/)

```shell
./go-sqlmap.exe -u http://sqlilab-ip/Less-1/?id=1
```

![](https://raw.githubusercontent.com/EmYiQing/go-sqlmap/master/img/0.png)

![](https://raw.githubusercontent.com/EmYiQing/go-sqlmap/master/img/1.png)
