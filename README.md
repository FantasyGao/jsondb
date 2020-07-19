# jsondb 
Json file database based on golang

## 介绍
基于golang的json文件数据库，免安装，直接用

## 安装

```bash
go get https://github.com/FantasyGao/jsondb
```

## 用法
```golang
import "github.com/FantasyGao/jsondb"

// 存储的文件位置~/Desktop/db.json
var fileName = "~/Desktop/db.json"

// 创建一个实例
db := jsondb.Create(fileName)

// 写入key为x,值为hello的进缓存内
db.Write("x", "hello")

/**
 * @description 写入key为y,值为world的数据进缓存内
 **/ 
db.Write("y", "world")
  .Write("z", 123456)

/**
 * @description 写入key为m,值为jsondb的，且保存当前的数据进入json文件
 **/ 
db.Write("m", "jsondb").Save()

/**
 * @description 查询key为y的值
 **/ 
val := db.Read("y")

/**
 * @description 查询所有数据，返回为map[string]interface{}类型
 **/ 
all := db.ReadAll()


/**
 * @description 删除key为x的值并且保存进json文件
 **/ 
db.Del("x").Save()
```

## License
This code is distributed under the MIT License.
