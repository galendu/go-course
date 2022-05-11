# 基于CA证书的双向TLS认证


## 凭证泄露问题





## TLS介绍

在不安全信道上构建安全信道，这是SSL的核心，所谓安全包括
+ 身份认证
+ 数据完整性
+ 数据加密性。

而非对称算法在TLS中的运用就是为了协商一个密钥，密钥的目的就是为了后续数据能够被加密，而加密密钥有且只有通信双方知道


## 准备证书

### 根证书

根证书（root certificate）是属于根证书颁发机构（CA）的公钥证书。我们可以通过验证CA的签名从而信任CA ，任何人都可以得到CA的证书（含公钥），用以验证它所签发的证书（客户端、服务端）

它包含的文件如下：
+ 公钥
+ 密钥

```sh
# 生成私钥
openssl genrsa -out ca.key 2048

# 生成公钥
openssl req -new -x509 -days 7200 -key ca.key -out ca.pem
Country Name (2 letter code) []:
State or Province Name (full name) []:
Locality Name (eg, city) []:
Organization Name (eg, company) []:
Organizational Unit Name (eg, section) []:
Common Name (eg, fully qualified host name) []:go-grpc-example
Email Address []:
```

生成完后会得到根证书的公钥和私钥
+ 私钥: ca.key
+ 公钥: ca.pem

### Server证书

使用我们自己建立的CA为Server端签发证书

```sh
# 生成私钥
openssl genrsa -out server_key.pem 4096

# 证书签署申请
openssl req -new                                    \
  -key server_key.pem                               \
  -days 3650                                        \
  -out server_csr.pem                               \
  -subj /C=US/ST=CA/L=SVL/O=gRPC/CN=test-server1/   


# CA签署证书
openssl x509 -req           \
  -in server_csr.pem        \
  -CAkey ../ca.key         \
  -CA ../ca.pem           \
  -days 3650                \
  -set_serial 1000          \
  -out server_cert.pem      \
  -sha256

# CA校验证书有效性
openssl verify -verbose -CAfile ../ca.pem  server_cert.pem
server_cert.pem: OK
```



### Client证书

```sh
# 生成私钥
openssl genrsa -out client_key.pem 4096

# 证书签署申请


# CA签署证书


```


## Grpc TLS双向认证


### Server





### Client








## 参考

+ [TLS/SSL 协议详解 (30) SSL中的RSA、DHE、ECDHE、ECDH流程与区别](https://blog.csdn.net/mrpre/article/details/78025940)
+ [Grpc examples create.sh](https://github.com/grpc/grpc-go/tree/master/examples/data/x509)