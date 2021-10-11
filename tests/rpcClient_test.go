package tests

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/yuanyu90221/ethereum_blockchain_services/ethjsonrpc"
)

func TestClient_GetBlockNumber(t *testing.T) {
	type fields struct {
		rpcCleint *rpc.Client
		ethClient *ethclient.Client
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *big.Int
		wantErr error
	}{
		{
			name: "networkId should be 97",
			args: args{
				ctx: context.TODO(),
			},
			want:    big.NewInt(97),
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rpc_client := "https://data-seed-prebsc-2-s3.binance.org:8545/"
			ec, error := ethjsonrpc.Connect(rpc_client)
			if error != nil {
				t.Errorf("%v", error)
				return
			}
			networkId, err := ec.GetNetworkId(tt.args.ctx)
			if err != nil {
				t.Errorf("Client.GetNetworId(context.TODO()) error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
			if networkId.Cmp(tt.want) != 0 {
				t.Errorf("Client.GetNewtworkId(context.TODO()) networkId = %v, want = %v", networkId, tt.want)
			}
		})
	}
}
