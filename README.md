# go-sqlmap

## 介绍

- sqlmap：渗透测试界的神器，这是一个简单的sqlmap
- 使用Golang重写的原因：高效、生成可执行文件直接运行、无需搭建环境等
- 测试通过sqli-lab前八关，传入一个有注入的url即可直接脱库
- 目前可以做到UnionSelect回显注入、Updatexml和Polygon的报错注入、布尔盲注

## 使用说明

- 使用`-u http://xxx/index.php?id=1`进行注入（默认会探测版本和所有数据库名）
- 使用`-D security`指定具体的数据库进行注入获得所有表（比如security数据库）
- 使用`-D security -T users`指定数据库和表名获得所有字段名（比如security数据库的users表）
- 使用`-D security -T users -C id,username,password`获得users表这三个字段的所有数据
- 使用`--technique U`指定使用Union注入（适用于sqli-lab前4关）
- 使用`--technique E`指定使用报错注入（使用于sqli-lab第5-6关）
- 使用`--technique B`指定使用布尔盲注（使用于sqli-lab第8关）
- 使用`--beta`参数可以激活Polygon报错函数（默认使用Updatexml因为更稳定）

## 图片

- Union注入

![](https://raw.githubusercontent.com/EmYiQing/github.com/EmYiQing/go-sqlmap/master/img/02.png)

- 报错注入

![](https://raw.githubusercontent.com/EmYiQing/github.com/EmYiQing/go-sqlmap/master/img/01.png)

- 布尔盲注

![](https://raw.githubusercontent.com/EmYiQing/github.com/EmYiQing/go-sqlmap/master/img/03.png)

## 快速开始

[下载地址](https://github.com/EmYiQing/github.com/EmYiQing/go-sqlmap/releases)

```shell
github.com/EmYiQing/go-sqlmap.exe -u http://sqlilab-ip/Less-1/?id=1 -D security -T users -C id,username,password --technique U
```


