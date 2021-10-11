package router

import (
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/yuanyu90221/ethereum_blockchain_services/ethjsonrpc"
	"github.com/yuanyu90221/ethereum_blockchain_services/services"
)

func SetTransactionRouter(rootRouter *gin.Engine) {
	// transaction/:txHash
	rootRouter.GET("/transaction/:txHash", func(ctx *gin.Context) {
		txHash := ctx.Param("txHash")
		txHashData := common.HexToHash(txHash)
		result, err := services.GetTransactionDataByHash(txHashData, ethjsonrpc.MyRpcClient, ctx)
		if err != nil {
			if err.Error() == ethereum.NotFound.Error() {
				ctx.JSON(http.StatusNotFound, gin.H{
					"error":   fmt.Sprintf("txHash:%s %s\n", txHash, err.Error()),
					"message": "get transaction by txHash, please try another input",
				})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error":   err.Error(),
				"message": "get transaction by txHash, please try another input",
			})
			return
		}
		ctx.JSON(http.StatusOK, result)
	})
}
