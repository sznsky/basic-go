package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

func main() {
	// 输入交易数据（来自某笔交易）
	txData := "0xa9059cbb0000000000000000000000005494befe3ce72a2ca0001fe0ed0c55b42f8c358f000000000000000000000000000000000000000000000000000000000836d54c"

	// ERC20 ABI（包含 transfer 方法）
	const abiJSON = `[
		{
			"type":"function",
			"name":"transfer",
			"inputs":[
				{"name":"recipient","type":"address"},
				{"name":"amount","type":"uint256"}
			],
			"outputs":[{"name":"","type":"bool"}]
		}
	]`

	// 去除 0x 并解码 hex 数据
	dataBytes, err := hex.DecodeString(strings.TrimPrefix(txData, "0x"))
	if err != nil {
		log.Fatalf("交易数据解码失败: %v", err)
	}

	// 前4字节是函数选择器
	methodID := dataBytes[:4]

	// 解析 ABI
	parsedABI, err := abi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		log.Fatalf("ABI 解析失败: %v", err)
	}

	// 查找匹配的函数
	var method *abi.Method
	for name, m := range parsedABI.Methods {
		if hex.EncodeToString(m.ID) == hex.EncodeToString(methodID) {
			mCopy := m // ✅ 必须复制变量，避免指针引用错误
			method = &mCopy
			fmt.Printf("📌 匹配函数: %s\n", name)
			break
		}
	}

	if method == nil {
		log.Fatalf("找不到匹配的函数 ID: %x", methodID)
	}

	// 剩余字节是参数部分
	argsData := dataBytes[4:]
	unpacked, err := method.Inputs.Unpack(argsData)
	if err != nil {
		log.Fatalf("参数解包失败: %v", err)
	}

	// 输出参数信息
	fmt.Println("📥 参数列表:")
	for i, input := range method.Inputs {
		val := unpacked[i]
		fmt.Printf("  - 名称: %s\n    类型: %s\n", input.Name, input.Type.String())

		switch v := val.(type) {
		case common.Address:
			fmt.Printf("    值: %s\n", v.Hex())
		case *big.Int:
			fmt.Printf("    值: %s\n", v.String())
		default:
			fmt.Printf("    值: %v\n", v)
		}
	}

}
