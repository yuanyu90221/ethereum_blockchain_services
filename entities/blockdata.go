package entities

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type BlockData struct {
	BlockNumber *big.Int    `json:"block_number"`
	BlockHash   common.Hash `json:"block_hash"`
	BlockTime   uint64      `json:"block_time"`
	ParentHash  common.Hash `json:"parant_hash"`
}

type BlockTxData struct {
	BlockNumber  *big.Int           `json:"block_number"`
	BlockHash    common.Hash        `json:"block_hash"`
	BlockTime    uint64             `json:"block_time"`
	ParentHash   common.Hash        `json:"parant_hash"`
	Transactions types.Transactions `json:"transactions"`
}
