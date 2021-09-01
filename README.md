# vito_infra 是一个基础包
集合配置信息，数据库链接，redis链接，与日志打印<br>

使用方式
```
go get -u github.com/VitoChueng/vito_infra
```

因为要用到etcd，go mod 需要添加

```
replace google.golang.org/grpc v1.29.1 => google.golang.org/grpc v1.26.0

replace github.com/coreos/bbolt v1.3.4 => go.etcd.io/bbolt v1.3.4
```
