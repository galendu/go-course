# 密码学简介与Golang的加密库Crypto的使用

据记载，公元前400年，古希腊人发明了置换密码。1881年世界上的第一个电话保密专利出现。在第二次世界大战期间，德国军方启用“恩尼格玛”密码机，密码学在战争中起着非常重要的作用, 这段历史很有趣,建议看看[恩格玛机破解历史](https://www.zhihu.com/question/28397034)。

随着信息化和数字化社会的发展，人们对信息安全和保密的重要性认识不断提高，于是在1997年，美国国家标准局公布实施了“美国数据加密标准（DES）”，民间力量开始全面介入密码学的研究和应用中，采用的加密算法有DES、RSA、SHA等。随着对加密强度需求的不断提高，近期又出现了AES、ECC等。


## 密码学的目的

+ 保密性：防止用户的标识或数据被读取。
+ 数据完整性：防止数据被更改。
+ 身份验证：确保数据发自特定的一方。

## 密码学的应用

随着密码学商业应用的普及，公钥密码学受到前所未有的重视。除传统的密码应用系统外，PKI系统以公钥密码技术为主，提供加密、签名、认证、密钥管理、分配等功能。

+ 保密通信：保密通信是密码学产生的动因。使用公私钥密码体制进行保密通信时，信息接收者只有知道对应的密钥才可以解密该信息。
+ 数字签名：数字签名技术可以代替传统的手写签名，而且从安全的角度考虑，数字签名具有很好的防伪造功能。在政府机关、军事领域、商业领域有广泛的应用环境。
+ 秘密共享：秘密共享技术是指将一个秘密信息利用密码技术分拆成n个称为共享因子的信息，分发给n个成员，只有k(k≤n)个合法成员的共享因子才可以恢复该秘密信息，其中任何一个或m(m≤k)个成员合作都不知道该秘密信息。利用秘密共享技术可以控制任何需要多个人共同控制的秘密信息、命令等。
+ 认证功能：在公开的信道上进行敏感信息的传输，采用签名技术实现对消息的真实性、完整性进行验证，通过验证公钥证书实现对通信主体的身份验证。
+ 密钥管理：密钥是保密系统中更为脆弱而重要的环节，公钥密码体制是解决密钥管理工作的有力工具；利用公钥密码体制进行密钥协商和产生，保密通信双方不需要事先共享秘密信息；利用公钥密码体制进行密钥分发、保护、密钥托管、密钥恢复等

接下来我们将依次介绍:
+ 散列算法
+ 对称加密算法
+ 非对称加密算法
+ 秘钥交换算法

## 散列算法

散列是信息的提炼，通常其长度要比信息小得多，且为一个固定长度。加密性强的散列一定是不可逆的，这就意味着通过散列结果，无法推出任何部分的原始信息。任何输入信息的变化，哪怕仅一位，都将导致散列结果的明显变化，这称之为雪崩效应。散列还应该是防冲突的，即找不出具有相同散列结果的两条信息。具有这些特性的散列结果就可以用于验证信息是否被修改。常用于保证数据完整性

单向散列函数一般用于产生消息摘要，密钥加密等，常见的有:
+ MD5(Message Digest Algorithm 5): 是RSA数据安全公司开发的一种单向散列算法。
+ SHA(Secure Hash Algorithm): 可以对任意长度的数据运算生成一个160位的数值

### MD5
MD5即Message-Digest Algorithm 5（信息-摘要算法5），用于确保信息传输完整一致。是计算机广泛使用的杂凑算法之一（又译摘要算法、哈希算法），主流编程语言普遍已有MD5实现。将数据（如汉字）运算为另一固定长度值，是杂凑算法的基础原理，MD5的前身有MD2、MD3和MD4

由于MD5已经被破解了(中国山东大学的王小云教授破解)

```go
import (
	"crypto/md5"
	"fmt"
)

func main() {
    // 最基础的使用方式: Sum 返回数据的MD5校验和
	fmt.Printf("%x\n", md5.Sum([]byte("测试数据")))
}
```

### SHA-1
在1993年，安全散列算法（SHA）由美国国家标准和技术协会(NIST)提出，并作为联邦信息处理标准（FIPS PUB 180）公布；1995年又发布了一个修订版FIPS PUB 180-1，通常称之为SHA-1。SHA-1是基于MD4算法的，并且它的设计在很大程度上是模仿MD4的。现在已成为公认的最安全的散列算法之一，并被广泛使用。

SHA-1是一种数据加密算法，该算法的思想是接收一段明文，然后以一种不可逆的方式将它转换成一段（通常更小）密文，也可以简单的理解为取一串输入码（称为预映射或信息），并把它们转化为长度较短、位数固定的输出序列即散列值（也称为信息摘要或信息认证代码）的过程。
该算法输入报文的最大长度不超过264位，产生的输出是一个160位的报文摘要。输入是按512 位的分组进行处理的。SHA-1是不可逆的、防冲突，并具有良好的雪崩效应。

sha1是SHA家族的五个算法之一(其它四个是SHA-224、SHA-256、SHA-384，和SHA-512)

```go
package main

import (
	"crypto/sha1"
	"fmt"
	"io"
)
// sha1散列算法
func sha1Hash(msg string) (hashData []byte) {
	h := sha1.New()
	io.WriteString(h, msg)
	hashData = h.Sum(nil)
	return
}

func main() {
	msg := "This is the message to hash!"
	// sha1
	sha1Data := sha1Hash(msg)
	fmt.Printf("SHA1: %x\n", sha1Data)
}
```

SHA-1与MD5的比较:

因为二者均由MD4导出，SHA-1和MD5彼此很相似。相应的，他们的强度和其他特性也是相似，但还有以下几点不同：

+ 对强行供给的安全性：最显著和最重要的区别是SHA-1摘要比MD5摘要长32 位。使用强行技术，产生任何一个报文使其摘要等于给定报摘要的难度对MD5是2128数量级的操作，而对SHA-1则是2160数量级的操作。这样，SHA-1对强行攻击有更大的强度。
+ 对密码分析的安全性：由于MD5的设计，易受密码分析的攻击，SHA-1显得不易受这样的攻击。
+ 速度：在相同的硬件上，SHA-1的运行速度比MD5慢

### HMac

Hmac算法也是一种哈希算法，它可以利用MD5或SHA1等哈希算法。不同的是，Hmac还需要一个密钥, 只要密钥发生了变化，那么同样的输入数据也会得到不同的签名，因此，可以把Hmac理解为用随机数“增强”的哈希算法

```go
package main

import (
	"crypto/hmac"
	"fmt"
	"io"
)

// 使用sha1的Hmac散列算法
func hmacHash(msg string, key string) (hashData []byte) {
	k := []byte(key)
	mac := hmac.New(sha1.New, k)
	io.WriteString(mac, msg)
	hashData = mac.Sum(nil)
	return
}

func main() {
	msg := "This is the message to hash!"
	// hmac
	hmacData := hmacHash(msg, "The key string!")
	fmt.Printf("HMAC: %x\n", hmacData)
}
```

### 你觉得上面这些算法好吗？
如果你使用的是MD5算法来加密你的口令，如果你的口令长度只有小写字母再加上数字，假设口令的长度是6位，那么在目前一台比较新一点的PC机上，穷举所有的口令只需要40秒钟。几乎有90%以上的用户只用小写字母和数字来组织其口令。对于6位长度的密码只需要最多40秒就可以破解了，这可能会吓到你

因为MD5，SHA的算法速度太快了, 当用于消息摘要，还是很不错的, 但是用于password hash就不行了
，有没有适合password hash的算法喃?

### bcrypt

bcrypt是这样的一个算法，因为它很慢，对于计算机来说，其慢得有点BT了，但却慢得刚刚好！对于验证用户口令来说是不慢的，对于穷举用户口令来说，其会让那些计算机变得如同蜗牛一样

bcrypt采用了一系列各种不同的Blowfish加密算法，并引入了一个work factor，这个工作因子可以让你决定这个算法的代价有多大。因为这些，这个算法不会因为计算机CPU处理速度变快了，而导致算法的时间会缩短了。因为，你可以增加work factor来把其性能降下来

同时bcrypt也是一种加盐的Hash方法，MD5 Hash时候，同一个报文经过hash的时候生成的是同一个hash值，在大数据的情况下，有些经过md5 hash的方法将会被破解(碰撞).使用BCrypt进行加密，同一个密码每次生成的hash值都是不相同的。每次加密的时候首先会生成一个随机数就是盐，之后将这个随机数与报文进行hash，得到 一个hash值

那一个被bcrypt hash过后的结果长啥样喃:

![bcrypt](../image/bcrypt.png)

Bcrypt有四个变量：

+ saltRounds: 正数，代表hash杂凑次数，数值越高越安全，默认10次。
+ myPassword: 明文密码字符串。
+ salt: 盐，一个128bits随机字符串，22字符
+ myHash: 经过明文密码password和盐salt进行hash，个人的理解是默认10次下 ，循环加盐hash10次，得到myHash

```go
package main

import (
    "fmt"

    "golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func main() {
    password := "secret"
    hash, _ := HashPassword(password) // ignore error for the sake of simplicity

    fmt.Println("Password:", password)
    fmt.Println("Hash:    ", hash)

    match := CheckPasswordHash(password, hash)
    fmt.Println("Match:   ", match)
}
```

## 对称加密算法

对称加密算法用来对敏感数据等信息进行加密，常用的算法包括：

+ DES(Data Encryption Standard): 数据加密标准，速度较快，适用于加密大量数据的场合。
+ 3DES(Triple DES): 是基于DES，对一块数据用三个不同的密钥进行三次加密，强度更高。
+ AES(Advanced Encryption Standard): 高级加密标准，是下一代的加密算法标准，速度快，安全级别高


### DES

DES全称为Data Encryption Standard，即数据加密标准，是一种使用密钥加密的块算法，1977年被美国联邦政府的国家标准局确定为联邦资料处理标准（FIPS），并授权在非密级政府通信中使用，随后该算法在国际上广泛流传开来。

AES与3DES的比较:

算法名称	算法类型	密钥长度	速度	解密时间（建设机器每秒尝试255个密钥）	资源消耗
AES	对称block密码	128、192、256位	高	1490000亿年	低
3DES	对称feistel密码	112位或168位	低	46亿年	中

```
破解历史
历史上有三次对DES有影响的攻击实验。1997年，利用当时各国 7万台计算机，历时96天破解了DES的密钥。1998年，电子边境基金会（EFF）用25万美元制造的专用计算机，用56小时破解了DES的密钥。1999年，EFF用22小时15分完成了破解工作。因此。曾经有过卓越贡献的DES也不能满足我们日益增长的需求了。
```

### AES
2000年10月，NIST(美国国家标准和技术协会)宣布通过从15种侯选算法中选出的一项新的密匙加密标准。Rijndael被选中成为将来的AES。 Rijndael是在1999年下半年，由研究员Joan Daemen和Vincent Rijmen创建的。AES正日益成为加密各种形式的电子数据的实际标准。
并于2002年5月26日制定了新的高级加密标准 (AES) 规范。

```
算法原理

AES算法基于排列和置换运算。排列是对数据重新进行安排，置换是将一个数据单元替换为另一个。AES 使用几种不同的方法来执行排列和置换运算。
AES是一个迭代的、对称密钥分组的密码，它可以使用128、192 和 256 位密钥，并且用 128 位（16字节）分组加密和解密数据。与公共密钥密码使用密钥对不同，对称密钥密码使用相同的密钥加密和解密数据。通过分组密码返回的加密数据的位数与输入数据相同。迭代加密使用一个循环结构，在该循环中重复置换和替换输入数据。
```

综上看来AES安全度最高, 基本现状就是AES已经替代DES成为新一代对称加密的标准, 下面是Golang中AES使用的栗子

```go
package main
import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)
var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}
func encrypt(plainText string, keyText string) (cipherByte []byte, err error) {
	// 转换成字节数据, 方便加密
	plainByte := []byte(plainText)
	keyByte := []byte(keyText)
	// 创建加密算法aes
	c, err := aes.NewCipher(keyByte)
	if err != nil {
		return nil, err
	}
	//加密字符串
	cfb := cipher.NewCFBEncrypter(c, commonIV)
	cipherByte = make([]byte, len(plainByte))
	cfb.XORKeyStream(cipherByte, plainByte)
	return
}
func decrypt(cipherByte []byte, keyText string) (plainText string, err error) {
	// 转换成字节数据, 方便加密
	keyByte := []byte(keyText)
	// 创建加密算法aes
	c, err := aes.NewCipher(keyByte)
	if err != nil {
		return "", err
	}
	// 解密字符串
	cfbdec := cipher.NewCFBDecrypter(c, commonIV)
	plainByte := make([]byte, len(cipherByte))
	cfbdec.XORKeyStream(plainByte, cipherByte)
	plainText = string(plainByte)
	return
}
func main() {
	plain := "The text need to be encrypt."
	// AES 规定有3种长度的key: 16, 24, 32分别对应AES-128, AES-192, or AES-256
	key := "abcdefgehjhijkmlkjjwwoew"
	// 加密
	cipherByte, err := encrypt(plain, key)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s ==> %x\n", plain, cipherByte)
	// 解密
	plainText, err := decrypt(cipherByte, key)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%x ==> %s\n", cipherByte, plainText)
}
```


## 非对称算法

非对称加密算法常用于数据加密和身份认证, 常见的非对称加密算法如下：

+ RSA: 由RSA公司发明，是一个支持变长密钥的公共密钥算法，需要加密的文件块的长度也是可变的；
+ DSA(Digital Signature Algorithm): 数字签名算法，是一种标准的DSS(数字签名标准)；
+ ECC(Elliptic Curves Cryptography): 椭圆曲线密码编码学。
+ ECDSA(Elliptic Curve Digital Signature Algorithm): 基于椭圆曲线的DSA签名算法