package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// 实现一个最小的区块链

// Transaction 交易结构体
type Transaction struct {
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
	Amount    int    `json:"amount"`
}

// Block 区块结构体
type Block struct {
	Index        int           `json:"index"`
	Timestamp    int64         `json:"timestamp"`
	Transactions []Transaction `json:"transactions"`
	Proof        int64         `json:"proof"`
	PreviousHash string        `json:"previous_hash"`
}

// Blockchain 区块链结构体
type Blockchain struct {
	Chain []Block
}

// CreateBlock 创建新区块
func (bc *Blockchain) CreateBlock(proof int64, previousHash string) Block {
	block := Block{
		Index:        len(bc.Chain) + 1,
		Timestamp:    time.Now().Unix(),
		Transactions: []Transaction{},
		Proof:        proof,
		PreviousHash: previousHash,
	}
	// 在原区块的后面增加一个区块，形成新的区块
	bc.Chain = append(bc.Chain, block)
	return block
}

// GetLastBlock 获取区块链最后一个区块
func (bc *Blockchain) GetLastBlock() Block {
	return bc.Chain[len(bc.Chain)-1]
}

// ProofOfWork 工作量证明(POW)
func (bc *Blockchain) ProofOfWork(lastProof int64) int64 {
	var proof int64 = 0
	for !bc.ValidProof(lastProof, proof) {
		proof++
	}
	return proof
}

// ValidProof 验证工作量证明
func (bc *Blockchain) ValidProof(lastProof, proof int64) bool {
	// 将最后一个工作证明值和当前工作证明值，合并计算哈希值
	guess := fmt.Sprintf("%d%d", lastProof, proof)
	guessHash := sha256.Sum256([]byte(guess))
	hexHash := hex.EncodeToString(guessHash[:])
	return strings.HasPrefix(hexHash, "0000")
}

// Hash 计算区块哈希值
func (bc *Blockchain) Hash(block Block) string {
	blockBytes, _ := json.Marshal(block)
	hash := sha256.Sum256(blockBytes)
	return hex.EncodeToString(hash[:])
}

// CreateGenesisBlock 创建创世区块
func CreateGenesisBlock() Blockchain {
	bc := Blockchain{}
	genesisBlock := Block{
		Index:        1,
		Timestamp:    time.Now().Unix(),
		Transactions: []Transaction{},
		Proof:        100,
		PreviousHash: "0",
	}
	bc.Chain = append(bc.Chain, genesisBlock)
	return bc
}

// AddTransaction 添加交易
func (bc *Blockchain) AddTransaction(sender, recipient string, amount int) {
	lastBlock := bc.GetLastBlock()
	newTransaction := Transaction{
		Sender:    sender,
		Recipient: recipient,
		Amount:    amount,
	}
	bc.Chain[len(bc.Chain)-1].Transactions = append(lastBlock.Transactions, newTransaction)
}

func main() {
	// 初始化区块链
	blockchain := CreateGenesisBlock()

	// 添加交易
	blockchain.AddTransaction("lilei", "lucy", 11)

	// 获取最后一个区块
	lastBlock := blockchain.GetLastBlock()
	lastProof := lastBlock.Proof
	proof := blockchain.ProofOfWork(lastProof)

	// 获取上一个区块的哈希值
	previousHash := blockchain.Hash(lastBlock)

	// 创建新区块链
	blockchain.CreateBlock(proof, previousHash)

	// 打印区块链信息
	fmt.Println("区块链的长度：", len(blockchain.Chain))
	// 遍历整个区块链
	for i, block := range blockchain.Chain {
		fmt.Printf("\n区块 %d:\n", i+1)
		fmt.Println("索引:", block.Index)
		fmt.Println("时间戳:", block.Timestamp)
		fmt.Println("交易数量:", len(block.Transactions))
		fmt.Println("工作量证明:", block.Proof)
		fmt.Println("前一区块哈希:", block.PreviousHash)
		fmt.Println("当前区块哈希:", blockchain.Hash(block))
	}
}
