/**
 * @Author: lixiang
 * @Date: 2025-03-22 14:52
 * @Description: token.go
 */

package models

type TokenInfo struct {
	Id           int    `gorm:"column:id;primaryKey"`
	Logo         string `json:"logo" gorm:"column:logo"`
	Token        string `json:"token" gorm:"column:token"`
	Symbol       string `json:"symbol" gorm:"column:symbol"`
	ChainId      string `json:"chain_id" gorm:"column:chain_id"`
	Price        string `json:"price" gorm:"column:price"`
	Decimals     int    `json:"decimals" gorm:"column:decimals"`
	AbiFileExist int    `json:"abi_file_exist" gorm:"column:abi_file_exist"`
	CreatedAt    string `json:"created_at" gorm:"column:created_at"`
	UpdatedAt    string `json:"updated_at" gorm:"column:updated_at"`
}
type RedisTokenInfo struct {
	Logo    string `json:"logo" gorm:"column:logo"`
	Token   string `json:"token" gorm:"column:token"`
	Symbol  string `json:"symbol" gorm:"column:symbol"`
	ChainId string `json:"chain_id" gorm:"column:chain_id"`
	Price   string `json:"price" gorm:"column:price"`
}

type BorrowToken struct {
	BorrowFee  string `json:"borrowFee"`
	TokenLogo  string `json:"tokenLogo"`
	TokenName  string `json:"tokenName"`
	TokenPrice string `json:"tokenPrice"`
}

type LendToken struct {
	LendFee    string `json:"lendFee"`
	TokenLogo  string `json:"tokenLogo"`
	TokenName  string `json:"tokenName"`
	TokenPrice string `json:"tokenPrice"`
}
