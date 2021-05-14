# Go-Sqlmap

## 介绍

- 效果
![](https://raw.githubusercontent.com/EmYiQing/go-sqlmap/master/img/1.png)

- sqlmap是渗透测试界的神器，于是想尝试写一个简单的sqlmap，深入理解sql注入
- 使用Golang重写的原因：高效、生成可执行文件直接运行、无需搭建环境等
- 目前测试在sqlilab的前3关可以成功，主要是针对mysql的回显注入功能，最终目标是可以用它通关sqlilab

## 快速开始

[可执行文件下载地址](https://github.com/EmYiQing/go-sqlmap/releases)

```shell
go-sqlmap.exe -u http://sqlilab-ip/Less-1/?id=1
```
![](https://raw.githubusercontent.com/EmYiQing/go-sqlmap/master/img/0.png)


