## RSA加密过程
1. A生成一对密钥（公钥和私钥），私钥不公开，公钥公开。
2. A传递自己的公钥给B，B用A的公钥对消息进行加密。
3. A接收到B加密的密文消息，利用A自己的私钥对消息进行解密。

## RSA签名过程
1. A生成一对密钥（公钥和私钥），私钥不公开，公钥公开。
2. A用自己的私钥对消息加签，形成签名，并将加签的消息和消息本身一起传递给B。（A先对消息用hash算法做一次处理，得到一个字符串哈希值，再用A的私钥对哈希值做一次加密得到一个签名，然后把签名和消息（原文）一起发送给B。）
3. B收到消息后，在获取A的公钥进行验签，如果验签出来的内容与消息本身一致，证明消息是A回复的（B用A的公钥对签名做解密处理，得到了哈希值a，然后用同样的hash算法对消息许做一次哈希处理，得到另一个哈希值b，对比a和b，如果两个值是相同的，那么可以确认消息是A发出来的）。

## 使用示例
生成公钥和私钥
```go
package main

import (
	"fmt"
	s
	"github.com/revevide/gtools/crypto/rsa"
)

func main() {
	bits := 2048
	prvKey, pubKey, err := rsa.GenKey(bits)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(prvKey))
	fmt.Println(string(pubKey))
}
```

加密解密
```go
package main

import (
	"fmt"
	
	"github.com/revevide/gtools/crypto/rsa"
)

func main() {
	...
	// 加密
	cipher, err := rsa.Encrypt([]byte(testPubKey), []byte(testPassword))
	if err != nil {
		panic(err)
	}
	// 解密
	plain, err := rsa.Decrypt([]byte(testPrvKey), cipher)
	if err != nil {
		panic(err)
	}
	...
}
```
