# Go-Sqlmap

## 介绍

- sqlmap：渗透测试界的神器，这是一个简单的sqlmap
- 使用Golang重写的原因：高效、生成可执行文件直接运行、无需搭建环境等
- 测试可以支持sqli-lab前六关，传入一个有注入的url即可直接脱库
- 目前可以做到UnionSelect回显注入、Updatexml和Polygon的报错注入

## 图片

- 使用`--technique U`指定使用Union注入（适用于sqli-lab前四关）
- 指定具体的数据库`-D security`，即可得到这个数据库所有表
- 再指定具体的表`-D security -T users`，即可得到这个表所有字段
- 指定具体的字段`-C id,username,password`，加上之前的条件，可以直接脱库

![](https://raw.githubusercontent.com/EmYiQing/go-sqlmap/master/img/02.png)

- 使用`--technique E`指定使用报错注入（适用于sqli-lab第5、6关）

![](https://raw.githubusercontent.com/EmYiQing/go-sqlmap/master/img/01.png)

## 快速开始

[下载地址](https://github.com/EmYiQing/go-sqlmap/releases)

```shell
go-sqlmap.exe -u http://sqlilab-ip/Less-1/?id=1
```


