package main

import (
	"context"
	"fmt"
	"github.com/rs/cors"
	"log"
	"math/big"
	"net/http"
	"os"
	"strings"
	"time"

	"gitee.com/geekbang/basic-go/web3/go-backend/db"
	"gitee.com/geekbang/basic-go/web3/go-backend/handler"
	"gitee.com/geekbang/basic-go/web3/go-backend/models"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

// transferEventSigHash 是从 Etherscan 验证过的非标准 ERC20 Transfer 事件签名哈希
var transferEventSigHash = common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")

const (
	// 定义每次查询的区块范围大小，必须严格遵守 Alchemy 的 500 限制
	blockQueryRange = 499 // 尝试设置为499，确保不会超500
	// 期望获取的历史交易数量
	targetHistoricalTxs = 5
	// 连接以太坊节点的超时时间
	ethDialTimeout = 10 * time.Second
)

func main() {
	// 加载 .env 文件，如果存在
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables directly.")
	}

	// 初始化数据库连接和表
	db.InitDB()

	// 获取以太坊节点 WebSocket URL
	ethNodeURL := os.Getenv("ETH_NODE_WS_URL")
	if ethNodeURL == "" {
		log.Fatal("ETH_NODE_WS_URL environment variable not set")
	}
	fmt.Println("ETH node:", ethNodeURL)

	// 获取 ERC20 代币地址
	tokenAddressStr := os.Getenv("ERC20_TOKEN_ADDRESS")
	if tokenAddressStr == "" {
		log.Fatal("ERC20_TOKEN_ADDRESS environment variable not set")
	}
	tokenAddress := common.HexToAddress(tokenAddressStr)
	fmt.Printf("Configured Token Address: %s\n", tokenAddress.Hex())

	// --- 连接以太坊节点并添加超时处理 ---
	var client *ethclient.Client
	var clientErr error
	ctxDial, cancelDial := context.WithTimeout(context.Background(), ethDialTimeout)
	defer cancelDial()

	done := make(chan struct{})
	go func() {
		client, clientErr = ethclient.DialContext(ctxDial, ethNodeURL)
		close(done)
	}()

	select {
	case <-done:
		if clientErr != nil {
			log.Fatalf("Failed to connect to Ethereum client: %v", clientErr)
		}
		fmt.Println("Connected to Ethereum node.")
	case <-ctxDial.Done():
		log.Fatalf("Failed to connect to Ethereum client: %v (timeout after %v)", ctxDial.Err(), ethDialTimeout)
	}
	// --- 连接以太坊节点超时处理结束 ---

	// 读取 ERC20 ABI
	abiFile, err := os.ReadFile("abi/erc20.abi")
	if err != nil {
		log.Fatalf("Failed to read ABI file: %v", err)
	}
	contractAbi, err := abi.JSON(strings.NewReader(string(abiFile)))
	if err != nil {
		log.Fatalf("Failed to parse contract ABI: %v", err)
	}
	fmt.Println("ERC20 ABI loaded successfully.")

	// --- 拉取历史交易逻辑开始 ---
	log.Println("Starting to fetch historical transactions...")

	numFetched := 0
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatalf("Failed to get latest block header: %v", err)
	}
	latestBlockNum := header.Number.Uint64()

	// 从当前最新区块开始，向前查询
	// currentQueryEndBlock 会在每次循环开始时设定为上一次查询的起始区块-1
	currentQueryEndBlock := latestBlockNum

	for numFetched < targetHistoricalTxs {
		// 计算当前查询范围的起始区块
		currentQueryStartBlock := uint64(0) // 默认值，如果计算结果小于0，则为0

		// 如果 currentQueryEndBlock - blockQueryRange 小于 0，说明我们已经快到创世区块了
		if currentQueryEndBlock <= blockQueryRange {
			currentQueryStartBlock = 0
		} else {
			currentQueryStartBlock = currentQueryEndBlock - blockQueryRange
		}

		// 如果查询范围已经到达或超过0，并且还没有获取到数据，就停止
		// 避免无限循环查询 0 到 0
		if currentQueryEndBlock == 0 && currentQueryStartBlock == 0 && numFetched == 0 {
			log.Println("Reached genesis block (or block 0) and no historical transactions found.")
			break
		}

		log.Printf("Querying Transfer events from block %d to %d for token %s", currentQueryStartBlock, currentQueryEndBlock, tokenAddress.Hex())
		log.Printf("Current transferEventSigHash: %s", transferEventSigHash.Hex())

		query := ethereum.FilterQuery{
			FromBlock: big.NewInt(int64(currentQueryStartBlock)),
			ToBlock:   big.NewInt(int64(currentQueryEndBlock)),
			Addresses: []common.Address{tokenAddress},
			Topics:    [][]common.Hash{{transferEventSigHash}},
		}

		ctxFilter, cancelFilter := context.WithTimeout(context.Background(), 30*time.Second)
		logs, filterErr := client.FilterLogs(ctxFilter, query)
		cancelFilter()

		if filterErr != nil {
			log.Printf("Failed to filter historical logs from block %d to %d: %v", currentQueryStartBlock, currentQueryEndBlock, filterErr)
			// 特别注意 Alchemy 的错误信息，它告诉你允许的范围。
			// 如果频繁出现这个错误，可能需要进一步缩小 blockQueryRange，或者使用更高级的 RPC 服务。
			if strings.Contains(filterErr.Error(), "rate limit") || strings.Contains(filterErr.Error(), "timeout") {
				log.Println("Rate limit or timeout detected. Waiting 5 seconds before retrying or adjusting query range.")
				time.Sleep(5 * time.Second)
			} else if strings.Contains(filterErr.Error(), "500 block range") {
				// 如果明确提示区块范围错误，可以再次尝试缩小范围，或者检查 RPC 服务商的最新文档
				log.Println("RPC provider explicitly rejected block range, likely due to exceeding limits. Consider reducing 'blockQueryRange' constant.")
			}
			// 无论何种错误，继续尝试查询更早的区块
			// 在这里不直接 `continue` 到下一个 `i`，而是更新 `currentQueryEndBlock`，确保下一轮是正确的往前推
		} else { // 只有在成功获取到日志时才处理
			if len(logs) == 0 {
				log.Printf("No logs found for token %s in block range %d to %d. Expanding search range.", tokenAddress.Hex(), currentQueryStartBlock, currentQueryEndBlock)
			} else {
				log.Printf("FilterLogs returned %d logs in total for block range %d to %d.", len(logs), currentQueryStartBlock, currentQueryEndBlock)

				// 倒序处理日志，以获取最新的 N 条
				for j := len(logs) - 1; j >= 0 && numFetched < targetHistoricalTxs; j-- {
					vLog := logs[j]

					log.Printf("Processing log (Block: %d, TxHash: %s, LogIndex: %d, Topics count: %d)",
						vLog.BlockNumber, vLog.TxHash.Hex(), vLog.Index, len(vLog.Topics))

					if len(vLog.Topics) < 3 {
						log.Printf("Warning: Non-standard Transfer event log (TxHash: %s). Topics count: %d, skipping. Expected at least 3 topics.", vLog.TxHash.Hex(), len(vLog.Topics))
						continue
					}

					log.Printf("Raw Topics: %v", vLog.Topics)
					log.Printf("Raw Data: %x", vLog.Data)

					var transferEvent struct {
						Value *big.Int
					}

					err := contractAbi.UnpackIntoInterface(&transferEvent, "Transfer", vLog.Data)
					if err != nil {
						log.Printf("Failed to unpack historical log data (TxHash: %s): %v. Check ABI! Raw Data: %x", vLog.TxHash.Hex(), err, vLog.Data)
						continue
					}

					transferEventFrom := common.BytesToAddress(vLog.Topics[1].Bytes())
					transferEventTo := common.BytesToAddress(vLog.Topics[2].Bytes())

					block, err := client.BlockByHash(context.Background(), vLog.BlockHash)
					if err != nil {
						log.Printf("Failed to get block details for historical log (BlockHash: %s, TxHash: %s): %v", vLog.BlockHash.Hex(), vLog.TxHash.Hex(), err)
						continue
					}
					timestamp := time.Unix(int64(block.Time()), 0)

					tx := models.Transaction{
						TxHash:       vLog.TxHash.Hex(),
						From:         transferEventFrom.Hex(),
						To:           transferEventTo.Hex(),
						Value:        transferEvent.Value,
						BlockNum:     vLog.BlockNumber,
						LogIndex:     uint(vLog.Index),
						Timestamp:    timestamp,
						TokenAddress: tokenAddress.Hex(),
					}

					if err := db.InsertTransaction(tx); err != nil {
						if strings.Contains(err.Error(), "UNIQUE constraint failed") || strings.Contains(err.Error(), "Duplicate entry") {
							log.Printf("Transaction %s already exists in DB, skipping. Block: %d, LogIndex: %d", tx.TxHash, tx.BlockNum, tx.LogIndex)
						} else {
							log.Printf("Failed to insert historical transaction %s: %v. Block: %d, LogIndex: %d", tx.TxHash, err, tx.BlockNum, tx.LogIndex)
						}
					} else {
						fmt.Printf("Inserted historical TX: %s | From: %s | To: %s | Value: %s | Block: %d | LogIndex: %d\n",
							tx.TxHash, tx.From, tx.To, tx.Value.String(), tx.BlockNum, tx.LogIndex)
						numFetched++
					}
				}
			}
		}

		// 在每次循环结束时，将查询的结束区块设置为当前查询的起始区块的前一个区块
		// 这样下一轮就会继续向前查询更早的区块
		if currentQueryStartBlock == 0 { // 如果已经查询到区块 0，则停止循环
			break
		}
		currentQueryEndBlock = currentQueryStartBlock - 1
	}

	log.Println("Finished fetching historical transactions.")

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // 临时允许所有来源 (仅供测试)
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
		Debug:          true,
	})

	// --- 启动 API 服务器 ---
	router := http.NewServeMux()
	router.HandleFunc("/transactions/", handler.GetTransactionsHandler)

	// 配置 CORS
	c = cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173", "https://hoppscotch.io"}, // 允许你的前端和 Hoppscotch
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
		Debug:          true, // 开启调试日志
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Starting API server on port %s...\n", port)
	// **重要：使用被 CORS 包裹的处理器**
	log.Fatal(http.ListenAndServe(":"+port, c.Handler(router)))
}
