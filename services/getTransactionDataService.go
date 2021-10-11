package services

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/yuanyu90221/ethereum_blockchain_services/entities"
	"github.com/yuanyu90221/ethereum_blockchain_services/ethjsonrpc"
)

func GetTransactionDataByHash(txHash common.Hash, rpcClient *ethjsonrpc.Client, ctx context.Context) (entities.TxData, error) {
	var result entities.TxData
	receipt, err := rpcClient.GetTransactionRecept(ctx, txHash)
	if err != nil {
		return result, err
	}
	transaction, _, err := rpcClient.GetTransactionByHash(ctx, txHash)
	if err != nil {
		return result, err
	}
	msg, err := transaction.AsMessage(types.NewEIP155Signer(transaction.ChainId()), transaction.GasFeeCap())
	if err != nil {
		return result, err
	}
	result = entities.TxData{TxHash: txHash, From: msg.From(), To: *(msg.To()), Data: transaction.Data(),
		Value: transaction.Value(), Logs: receipt.Logs}
	return result, nil
}
