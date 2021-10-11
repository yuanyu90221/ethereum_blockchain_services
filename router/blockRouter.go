package router

import (
	"fmt"
	"math/big"
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum"
	"github.com/gin-gonic/gin"
	"github.com/yuanyu90221/ethereum_blockchain_services/ethjsonrpc"
	"github.com/yuanyu90221/ethereum_blockchain_services/services"
)

func SetBlockRouter(rootRouter *gin.Engine) {
	// blocks?limit=n
	rootRouter.GET("/blocks", func(ctx *gin.Context) {
		limit := ctx.Query("limit")
		limitNumber, _ := strconv.Atoi(limit)
		result, err := services.GetLastNBlockData(limitNumber, ethjsonrpc.MyRpcClient, ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error":   err,
				"message": "get block by id failed, please check input",
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"block": result,
		})
	})
	// blocks/:id
	rootRouter.GET("/blocks/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		blockNumber, _ := strconv.Atoi(id)
		result, err := services.GetBlockTxDataById(big.NewInt(int64(blockNumber)), ethjsonrpc.MyRpcClient, ctx)
		if err != nil {
			if err.Error() == ethereum.NotFound.Error() {
				ctx.JSON(http.StatusNotFound, gin.H{
					"error":   fmt.Sprintf("id:%s %s\n", id, err.Error()),
					"message": "get transaction by txHash, please try another input",
				})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error":   err,
				"message": "get block by id failed, please check input",
			})
			return
		}
		ctx.JSON(http.StatusOK, result)
	})
}
