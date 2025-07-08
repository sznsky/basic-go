package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

// 使用go语言的工作量证明
func main() {
	nickname := "再出发"  // 替换为你的昵称
	target4 := "0000"  // 4个0开头的目标
	target5 := "00000" // 5个0开头的目标

	// 寻找4个0开头的哈希
	findTarget(nickname, target4, 4)

	// 寻找5个0开头的哈希
	findTarget(nickname, target5, 5)
}

// 计算字符串的hash值，一直计算到目标前缀的出现
func findTarget(nickname, target string, zeros int) {
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
			return
		}
		nonce++
	}
}
