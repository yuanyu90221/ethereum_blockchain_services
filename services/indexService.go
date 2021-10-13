package services

import (
	"context"
	"log"
	"math/big"

	"github.com/yuanyu90221/ethereum_blockchain_services/configs"
	"github.com/yuanyu90221/ethereum_blockchain_services/ethjsonrpc"
)

func IndexService(ctx context.Context, rpcClient *ethjsonrpc.Client) {
	jobChan := make(chan int64, 10)
	initIdx := int64(configs.GetEnvConfig().INIT_INDEX)
	defer close(jobChan)
	go worker(ctx, jobChan, rpcClient)
	for {
		// get current block Number
		latestBlockNumber, err := rpcClient.GetBlockNumber(ctx)
		if err != nil {
			log.Fatal(err)
		}
		for job := initIdx; job <= latestBlockNumber.Int64(); job++ {
			jobChan <- job
		}
	}
}
func worker(ctx context.Context, jobChan <-chan int64, rpcClient *ethjsonrpc.Client) {
	for {
		select {
		case <-ctx.Done():
			return
		case job := <-jobChan:
			Process(job, rpcClient, ctx)
		}
	}
}

func Process(job int64, rpcClient *ethjsonrpc.Client, ctx context.Context) {
	blockData, err := GetBlockTxDataById(big.NewInt(job), rpcClient, ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%v\n", blockData)
	// TODO: add UPSERT DATABASE logic
}
