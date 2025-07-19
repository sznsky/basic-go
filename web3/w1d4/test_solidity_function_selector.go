package main

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
)

// 测试solidity生成函数选择器代码
func main() {
	// 比如函数transfer(address,uint256) 获取函数签名选择器
	//functionName := "transfer(address,uint256)"
	//hash := crypto.Keccak256([]byte(functionName))
	//var hashString = hex.EncodeToString(hash[:])
	//fmt.Println(hashString)

	// 1.函数签名字符串
	inputString := "transfer(address,uint256)"
	// 将字符串转为字节数组
	inputBytes := []byte(inputString)

	// 2.调用Keccak256 计算hash
	hashBytes := crypto.Keccak256(inputBytes)

	// 3.将hash结果转为16进制
	hashString := hex.EncodeToString(hashBytes)

	// 4.在以太坊中，函数选择器就是hash的前4个字节
	functionSelector := hashString[:8]
	fmt.Printf("输入字符串: %s\n", functionSelector)
	fmt.Printf("该函数签名的函数选择器 (Function Selector) 是: 0x%s\n", functionSelector)
}
