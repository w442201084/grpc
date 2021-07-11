
### 一、pb协议生成路径
```
// pb协议
option go_package = "path;name";

// 生成文件
protoc --go_out=./ --micro_out=./ ./*.proto
```
