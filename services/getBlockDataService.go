package services

import (
	"context"
	"math/big"
	"sort"

	"github.com/yuanyu90221/ethereum_blockchain_services/entities"
	"github.com/yuanyu90221/ethereum_blockchain_services/ethjsonrpc"
)

func GetBlockTxDataById(id *big.Int, rpcClient *ethjsonrpc.Client, ctx context.Context) (entities.BlockTxData, error) {
	var blockTxData entities.BlockTxData
	blockData, err := rpcClient.GetBlockByNumber(ctx, id)
	if err != nil {
		return blockTxData, err
	}
	blockTxData = entities.BlockTxData{BlockNumber: blockData.Number(), BlockTime: blockData.Time(), BlockHash: blockData.Hash(), ParentHash: blockData.ParentHash(), Transactions: blockData.Transactions()}
	return blockTxData, err
}

func GetLastNBlockData(limit int, rpcClient *ethjsonrpc.Client, ctx context.Context) ([]entities.BlockData, error) {
	blockData := []entities.BlockData{}
	latestBlockNumber, err := rpcClient.GetBlockNumber(ctx)
	if err != nil {
		return blockData, err
	}
	dataChannel := make(chan entities.BlockData, limit)
	for i := 0; i < limit; i++ {
		latestBlockNumber = big.NewInt(0).Sub(latestBlockNumber, big.NewInt(int64(1)))
		go GetBlockRoutine(rpcClient, latestBlockNumber, ctx, dataChannel)
	}
	for response := range dataChannel {
		blockData = append(blockData, response)
		if len(blockData) == limit {
			close(dataChannel)
		}
		sort.Slice(blockData, func(i, j int) bool {
			return (blockData[i].BlockNumber.Cmp(blockData[j].BlockNumber)) >= 0
		})
	}
	return blockData, nil
}

func GetBlockRoutine(rpcClient *ethjsonrpc.Client, blockNumber *big.Int, ctx context.Context,
	ch chan entities.BlockData) {
	block, err := rpcClient.GetBlockByNumber(ctx, blockNumber)
	var blockData entities.BlockData
	if err != nil {
		ch <- blockData
		return
	}
	blockData = entities.BlockData{BlockNumber: block.Number(), BlockHash: block.Hash(),
		BlockTime: block.Time(), ParentHash: block.ParentHash()}
	ch <- blockData
}
