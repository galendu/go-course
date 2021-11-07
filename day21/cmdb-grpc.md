# gRPC版CMDB


## 自定义Struct Tag

安装插件: 
```sh
go install github.com/favadi/protoc-go-inject-tag@latest
```

定义protobuf:
```protobuf
message IP {
  string Address = 1; // @gotags: valid:"ip" yaml:"ip" json:"overrided"
}
```

使用插件:
```
protoc-go-inject-tag -input=./test.pb.go
```