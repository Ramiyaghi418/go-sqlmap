# Go-Sqlmap

## 介绍

- sqlmap是渗透测试界的神器，于是想尝试写一个简单的sqlmap，深入理解sql注入
- 使用Golang重写的原因：高效、生成可执行文件直接运行、无需搭建环境等

## 图片

效果如图，目前能够完成基于报错的注入，sqli-lab前三关没问题

- 直接使用
![](https://raw.githubusercontent.com/EmYiQing/go-sqlmap/master/img/01.png)
  
- 指定数据库

![](https://raw.githubusercontent.com/EmYiQing/go-sqlmap/master/img/02.png)

- 指定表名

![](https://raw.githubusercontent.com/EmYiQing/go-sqlmap/master/img/03.png)

- 指定字段进行脱裤

![](https://raw.githubusercontent.com/EmYiQing/go-sqlmap/master/img/04.png)

## 快速开始

[下载地址](https://github.com/EmYiQing/go-sqlmap/releases)

```shell
go-sqlmap.exe -u http://sqlilab-ip/Less-1/?id=1
```


