/**
 * @Author: lixiang
 * @Date: 2025-03-20 10:46
 * @Description: poolRequest.go
 */

package request

type PoolRequest struct {
	ChainId int `json:"chainId" binding:"required"`
}
