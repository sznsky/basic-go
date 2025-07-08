package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"time"
)

func main() {

	// 1.生成公钥和私钥对
	privateKey, publicKey := generateRSAKeyPair()
	fmt.Println("RSA密钥生成成功")

	// 2.执行工作量证明,找到需要的hash值
	nickname := "再出发"
	data, hash := performPOW(nickname, "0000", 4)
	fmt.Printf("\n POW结果：\n输出内容，%s\n哈希值：%s\n", data, hash)

	// 3.使用私钥对数据进去签名
	signature := signData(privateKey, data)

	// 4.使用公钥验签签名
	valid := verifySignature(publicKey, data, signature)
	if valid {
		fmt.Printf("\n 签名验证成功")
	} else {
		fmt.Printf("\n 签名验证失败")
	}
}

// 生成私钥和公钥的方法
// 1.生成RSA密钥对,2048是推荐的密钥长度 2.两个返回值，指向rsa对象的两个公钥和私钥的指针
func generateRSAKeyPair() (*rsa.PrivateKey, *rsa.PublicKey) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatal("密钥生成失败：", err.Error())
	}
	// 通过私钥获取公钥
	return privateKey, &privateKey.PublicKey
}

// 执行工作量证明(找到指定数量0开头的哈希，这里是4个0)
func performPOW(nickname, target string, zeros int) (string, string) {
	start := time.Now()
	nonce := 0
	var hashString string
	for {
		// 组合昵称
		data := fmt.Sprintf("%s%d", nickname, nonce)
		// 计算sha256哈希值
		hash := sha256.Sum256([]byte(data))
		// 将哈希值进行编码, hash[:] 将数组转为切片
		hashString = hex.EncodeToString(hash[:])

		// 检查是否满足目标条件
		if hashString[:zeros] == target {
			elapsed := time.Since(start)
			fmt.Printf("\n找到 %d 个0开头的哈希：\n", zeros)
			fmt.Printf("花费的时间：%v\n", elapsed)
			fmt.Printf("输入的内容：%s\n", data)
			fmt.Printf("哈希值：%s\n", hashString)
			fmt.Printf("尝试次数：%d\n", nonce)
			return data, hashString
		}
		nonce++
	}
}

// 使用私钥签名数据
func signData(privateKey *rsa.PrivateKey, data string) []byte {
	// 将数据进行hash256 hash计算
	hashed := sha256.Sum256([]byte(data))
	// 获取这个hash值的私钥签名（hashed[:] 将整个数组转为切片）
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
	if err != nil {
		log.Fatal("签名失败：", err.Error())
	}
	return signature
}

// 使用公钥验证签名
func verifySignature(publicKey *rsa.PublicKey, data string, signature []byte) bool {
	// 将数据进行hash256 hash计算
	hashed := sha256.Sum256([]byte(data))
	err := rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashed[:], signature)
	return err == nil
}
