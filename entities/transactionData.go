package entities

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type TxData struct {
	TxHash common.Hash    `json:"tx_hash"`
	From   common.Address `json:"from"`
	To     common.Address `json:"to"`
	Nonce  uint64         `json:"nonce"`
	Data   []byte         `json:"data"`
	Value  *big.Int       `json:"value"`
	Logs   []*types.Log   `json:"logs"`
}
