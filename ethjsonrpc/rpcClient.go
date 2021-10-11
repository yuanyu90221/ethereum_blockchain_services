package ethjsonrpc

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/yuanyu90221/ethereum_blockchain_services/configs"
)

type Client struct {
	rpcClient *rpc.Client
	ethClient *ethclient.Client
}

var MyRpcClient *Client

func Connect(host string) (*Client, error) {
	rpcClient, err := rpc.Dial(host)

	if err != nil {
		return nil, err
	}

	ethClient := ethclient.NewClient(rpcClient)
	return &Client{rpcClient, ethClient}, nil
}

func (ec *Client) Close() {
	ec.ethClient.Close()
}
func DoConnect() {
	rpc_client_endpoint := configs.GetEnvConfig().RPC_CLIENT_ENDPOINT
	tempClient, err := Connect(rpc_client_endpoint)
	if err == nil {
		MyRpcClient = tempClient
	}
}
func init() {
	DoConnect()
}
func GetConnect() *Client {
	if MyRpcClient == nil {
		DoConnect()
	}
	return MyRpcClient
}
func (ec *Client) GetBlockNumber(ctx context.Context) (*big.Int, error) {
	var result hexutil.Big
	err := ec.rpcClient.CallContext(ctx, &result, "eth_blockNumber")
	return (*big.Int)(&result), err
}

func (ec *Client) GetNetworkId(ctx context.Context) (*big.Int, error) {
	return ec.ethClient.NetworkID(ctx)
}

func (ec *Client) GetChainId(ctx context.Context) (*big.Int, error) {
	return ec.ethClient.ChainID(ctx)
}
func (ec *Client) GetBlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error) {
	return ec.ethClient.BlockByNumber(ctx, number)
}

func (ec *Client) GetBlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error) {
	return ec.ethClient.BlockByHash(ctx, hash)
}

func (ec *Client) GetTransactionByHash(ctx context.Context, hash common.Hash) (*types.Transaction, bool, error) {
	return ec.ethClient.TransactionByHash(ctx, hash)
	// var result types.Transaction
	// err := ec.rpcClient.CallContext(ctx, &result, "eth_getTransactionByHash", hash)
	// return (*types.Transaction)(&result), err
}

func (ec *Client) GetTransactionRecept(ctx context.Context, hash common.Hash) (*types.Receipt, error) {
	fmt.Printf("hash=%v\n", hash)
	return ec.ethClient.TransactionReceipt(ctx, hash)
}
