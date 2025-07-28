package models

import (
	"encoding/json" // 新增导入
	"math/big"
	"time"
)

// Transaction represents an ERC20 token transfer
type Transaction struct {
	// ID 是主键，通常由数据库自动生成
	ID int `json:"id" gorm:"primaryKey"`

	// 使用 gorm:"column:tx_hash;uniqueIndex" 确保数据库列名为 tx_hash 且有唯一索引
	TxHash string `json:"txHash" gorm:"column:tx_hash;uniqueIndex;type:varchar(66)"`

	// from_address 和 to_address 是常见的数据库列名，避免使用 Go 关键字 'from' 和 'to'
	From string `json:"from" gorm:"column:from_address;type:varchar(42)"`
	To   string `json:"to" gorm:"column:to_address;type:varchar(42)"`

	// value_wei 用于存储大整数，通常以字符串形式存储以避免精度问题
	Value *big.Int `json:"value" gorm:"column:value_wei;type:varchar(78)"` // 根据实际可能的最大值调整长度

	// block_num 和 log_index 使用下划线命名规范
	BlockNum uint64 `json:"blockNum" gorm:"column:block_num"`
	LogIndex uint   `json:"logIndex" gorm:"column:log_index"` // 修正 json 标签重复错误

	Timestamp time.Time `json:"timestamp" gorm:"column:timestamp"`

	// token_address 使用下划线命名规范
	TokenAddress string `json:"tokenAddress" gorm:"column:token_address;type:varchar(42)"`
}

// MarshalJSON 方法自定义 Transaction 结构体如何被编码为 JSON
// 确保 *big.Int 类型的 Value 字段被编码为 JSON 字符串
func (t Transaction) MarshalJSON() ([]byte, error) {
	// 修正：LogIndex 的 json tag 在原始代码中重复了，这里假设正确的是 "logIndex"
	// 这里的 `json:"json:"blockNum"` 看起来也有问题，修正为 `json:"blockNum"`

	// 定义一个与 Transaction 相同的匿名结构体，用于避免递归调用 MarshalJSON
	// 但其 value 字段会是 string 类型，用于 JSON 序列化
	type Alias Transaction
	return json.Marshal(&struct {
		Alias
		Value string `json:"value"` // 将 Value 字段声明为字符串类型，用于JSON编码
	}{
		Alias: (Alias)(t),       // 将原始 Transaction 的所有字段（除了被重写的）复制过来
		Value: t.Value.String(), // 将 *big.Int 转换为字符串
	})
}
