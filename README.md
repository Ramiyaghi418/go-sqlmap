# go-sqlmap

## 介绍

- sqlmap：渗透测试界的神器，这是一个简单的sqlmap（可以把它当成玩具）
- 使用Golang重写的原因：高效、生成可执行文件直接运行、无需搭建环境等
- 测试通过sqli-lab（https://github.com/Audi-1/sqli-labs）部分关卡
- 目前仅支持GET请求和MySQL数据库，支持Union注入、报错注入和布尔盲注
- 检测是否存在SQL注入的部分不准确，希望有大佬可以帮忙改进

## Introduce

- Sqlmap is a famous tool in penetration testing field, and this is a simple sqlmap
- The reasons for using golang rewriting are: high efficiency, generating executable files to run directly, etc
- At present, sqli-labs(https://github.com/Audi-1/sqli-labs) can be successfully injected into the first eight levels
- Union Select Injection,Updatexml and Polygon Error Based Injection and Bool Blind Injection are currently supported

## Instructions

- Use`-u http://xxx/index.php?id=1`Do Injection(By default, the version and all database names are probed)
- Use`-D security`Specify the specific database for injection to obtain all tables(such as security database)
- Use`-D security -T users`Specify database and table names to get all field names(such as users table of security database)
- Use`-D security -T users -C id,username,password`Get all the data of the three fields in the users table
- Use`--technique U`Specifies to use union select injection(sqli-labs 1,2,3,4)
- Use`--technique E`Specifies to use error based injection(sqli-lbs 5,6)
- Use`--technique B`Specifies to use bool blind injection(sqli-labs 8)
- Use`--beta`Parameter to activate the polygon error function(use updatexml by default because it is more stable)

## 图片

- Use Sqli-Labs For Test: [sqli-labs github](https://github.com/Audi-1/sqli-labs)

![](https://raw.githubusercontent.com/EmYiQing/go-sqlmap/master/img/demo.gif)

- Union Select Injection

![](https://raw.githubusercontent.com/EmYiQing/go-sqlmap/master/img/02.png)

- Error Based Injection

![](https://raw.githubusercontent.com/EmYiQing/go-sqlmap/master/img/01.png)

- Bool Blind Injection

![](https://raw.githubusercontent.com/EmYiQing/go-sqlmap/master/img/03.png)

## Quick Start

[Download](https://github.com/EmYiQing/go-sqlmap/releases)

```shell
go-sqlmap.exe -u http://sqlilab-ip/Less-1/?id=1 -D security -T users -C id,username,password --technique U
```

## API

```shell
go get https://github.com/EmYiQing/go-sqlmap
```

```go
package main

import (
	sqlmap "github.com/EmYiQing/go-sqlmap/api"
	"github.com/EmYiQing/go-sqlmap/input"
)

func main() {
	opts := input.Input{
		Beta:      false,
		Url:       "http://192.168.222.129:81/Less-1/?id=1",
		Database:  "security",
		Table:     "users",
		Columns:   []string{"id", "username", "password"},
		Technique: []string{"U"},
		Param:     "id",
	}
	instance := sqlmap.NewScanner(opts)
	instance.Run()
}
```


