package db

import (
	"database/sql"
	"fmt"
	"gitee.com/geekbang/basic-go/web3/go-backend/models" // 确保路径正确
	"log"
	"math/big"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql" // 导入 MySQL 驱动
)

var DB *sql.DB

// InitDB initializes the database connection and ensures the transactions table exists.
func InitDB() {
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		log.Fatal("DATABASE_URL environment variable not set. Example: user:password@tcp(127.0.0.1:3306)/database_name?parseTime=true")
	}

	var err error
	DB, err = sql.Open("mysql", connStr)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	// Set connection pool parameters (recommended)
	DB.SetMaxOpenConns(20)
	DB.SetMaxIdleConns(10)
	DB.SetConnMaxLifetime(5 * time.Minute) // Connections will be reused for max 5 minutes

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	log.Println("Successfully connected to MySQL database.")

	// --- Re-introducing and refining the CREATE TABLE operation ---
	// This SQL statement must precisely match the column names defined in models.Transaction's gorm tags
	createTableSQL := `
    CREATE TABLE IF NOT EXISTS transactions (
        id INT AUTO_INCREMENT PRIMARY KEY,
        tx_hash VARCHAR(66) UNIQUE NOT NULL,      -- Matches gorm:"column:tx_hash"
        from_address VARCHAR(42) NOT NULL,        -- Matches gorm:"column:from_address"
        to_address VARCHAR(42) NOT NULL,          -- Matches gorm:"column:to_address"
        value_wei VARCHAR(78) NOT NULL,           -- Matches gorm:"column:value_wei", stores big.Int as string
        block_num BIGINT NOT NULL,                -- Matches gorm:"column:block_num"
        log_index INT NOT NULL,                   -- Matches gorm:"column:log_index"
        timestamp DATETIME NOT NULL,              -- Matches gorm:"column:timestamp"
        token_address VARCHAR(42) NOT NULL        -- Matches gorm:"column:token_address"
    );`

	_, err = DB.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("Failed to create or verify 'transactions' table: %v", err)
	}
	fmt.Println("Database initialized. 'transactions' table ensured to exist and match schema.")
}

// InsertTransaction inserts a new transaction into the database, ignoring duplicates based on tx_hash.
func InsertTransaction(tx models.Transaction) error {
	// Column names in the INSERT statement must match the actual database column names
	query := `
    INSERT IGNORE INTO transactions (tx_hash, from_address, to_address, value_wei, block_num, log_index, timestamp, token_address)
    VALUES (?, ?, ?, ?, ?, ?, ?, ?);
    `
	// MySQL uses ? as placeholder
	_, err := DB.Exec(query,
		tx.TxHash,
		tx.From,
		tx.To,
		tx.Value.String(), // Store *big.Int as string
		tx.BlockNum,
		tx.LogIndex,
		tx.Timestamp,
		tx.TokenAddress,
	)
	if err != nil {
		return fmt.Errorf("error inserting transaction: %w", err)
	}
	return nil
}

// GetTransactionsByAddress retrieves transactions related to a specific address.
func GetTransactionsByAddress(address string) ([]models.Transaction, error) {
	// Column names in the SELECT statement must match the actual database column names
	query := `
    SELECT id, tx_hash, from_address, to_address, value_wei, block_num, log_index, timestamp, token_address
    FROM transactions
    WHERE from_address = ? OR to_address = ?
    ORDER BY block_num DESC, log_index DESC;
    `
	// MySQL uses ? as placeholder
	rows, err := DB.Query(query, address, address)
	if err != nil {
		return nil, fmt.Errorf("error querying transactions: %w", err)
	}
	defer rows.Close()

	var transactions []models.Transaction
	for rows.Next() {
		var tx models.Transaction
		var valueStr string // Temporary variable to scan the string value from DB

		// Scan into struct fields, ensure order and type match the SELECT query
		err := rows.Scan(
			&tx.ID,
			&tx.TxHash,
			&tx.From,
			&tx.To,
			&valueStr, // Scan into string
			&tx.BlockNum,
			&tx.LogIndex,
			&tx.Timestamp,
			&tx.TokenAddress,
		)
		if err != nil {
			log.Printf("Error scanning transaction row: %v", err)
			continue
		}

		// Convert valueStr back to *big.Int
		tx.Value = new(big.Int)
		_, ok := tx.Value.SetString(valueStr, 10)
		if !ok {
			log.Printf("Error converting value string '%s' to big.Int for TxHash %s: Invalid number format", valueStr, tx.TxHash)
			continue
		}
		transactions = append(transactions, tx)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating transaction rows: %w", err)
	}
	return transactions, nil
}
