# 基于CA证书的双向TLS认证


## 凭证泄露问题





## TLS介绍

在不安全信道上构建安全信道，这是SSL的核心，所谓安全包括
+ 身份认证
+ 数据完整性
+ 数据加密性。

而非对称算法在TLS中的运用就是为了协商一个密钥，密钥的目的就是为了后续数据能够被加密，而加密密钥有且只有通信双方知道

## 准备证书


### 自建CA



### Server证书



### Client证书



## Grpc TLS双向认证


### Server





### Client








## 参考

+ [TLS/SSL 协议详解 (30) SSL中的RSA、DHE、ECDHE、ECDH流程与区别](https://blog.csdn.net/mrpre/article/details/78025940)