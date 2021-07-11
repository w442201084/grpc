
### 一、拉取(micro)镜像
```
docker pull micro/micro

// go env -w GOPROXY="http://mirrors.aliyun.com/goproxy/,direct"


```

### 二、创建一个module
```
--rm        :       保证每次结束运行容器之后删除
-v          :       挂在数据卷
-w          :       指定工作目录

docker run --rm -v $(pwd):$(pwd) -w $(pwd) \
micro/micro new user[模块名称]
```

### 三、Gorm
```
// file head need to add ...
_ "github.com/jinzhu/gorm/dialects/mysql"

# git包
go get -u github.com/jinzhu/gorm
go get -u github.com/go-sql-driver/mysql

```
