/**
 * @Author: lixiang
 * @Date: 2025-03-20 10:34
 * @Description: poolBaseInfo.go
 */

package models

type PoolBaseInfo struct {
	PoolID                 int             `json:"pool_id"`
	AutoLiquidateThreshold string          `json:"autoLiquidateThreshold"`
	BorrowSupply           string          `json:"borrowSupply"`
	BorrowToken            string          `json:"borrowToken"`
	BorrowTokenInfo        BorrowTokenInfo `json:"borrowTokenInfo"`
	EndTime                string          `json:"endTime"`
	InterestRate           string          `json:"interestRate"`
	JpCoin                 string          `json:"jpCoin"`
	LendSupply             string          `json:"lendSupply"`
	LendToken              string          `json:"lendToken"`
	LendTokenInfo          LendTokenInfo   `json:"lendTokenInfo"`
	MartgageRate           string          `json:"martgageRate"`
	MaxSupply              string          `json:"maxSupply"`
	SettleTime             string          `json:"settleTime"`
	SpCoin                 string          `json:"spCoin"`
	State                  string          `json:"state"`
}

// PoolBases is a struct that represents the poolbases table in the database
type PoolBases struct {
	Id                     int    `json:"-" gorm:"column:id;primaryKey;autoIncrement"`
	PoolId                 int    `json:"pool_id" gorm:"column:pool_id"`
	ChainId                string `json:"chain_id" gorm:"column:chain_id"`
	SettleTime             string `json:"settle_time" gorm:"column:settle_time"`
	EndTime                string `json:"end_time" gorm:"column:end_time"`
	InterestRate           string `json:"interest_rate" gorm:"column:interest_rate"`
	MaxSupply              string `json:"max_supply" gorm:"max_supply:"`
	LendSupply             string `json:"lend_supply" gorm:"column:lend_supply"`
	BorrowSupply           string `json:"borrow_supply" gorm:"column:borrow_supply"`
	MartgageRate           string `json:"martgage_rate" gorm:"column:martgage_rate"`
	LendToken              string `json:"lend_token" gorm:"column:lend_token"`
	LendTokenInfo          string `json:"lend_token_info" gorm:"column:lend_token_info"`
	BorrowToken            string `json:"borrow_token" gorm:"column:borrow_token"`
	BorrowTokenInfo        string `json:"borrow_token_info" gorm:"column:borrow_token_info"`
	State                  string `json:"state" gorm:"column:state"`
	SpCoin                 string `json:"sp_coin" gorm:"column:sp_coin"`
	JpCoin                 string `json:"jp_coin" gorm:"column:jp_coin"`
	LendTokenSymbol        string `json:"lend_token_symbol" gorm:"column:lend_token_symbol"`
	BorrowTokenSymbol      string `json:"borrow_token_symbol" gorm:"column:borrow_token_symbol"`
	AutoLiquidateThreshold string `json:"auto_liquidate_threshold" gorm:"column:auto_liquidate_threshold"`
	CreatedAt              string `json:"created_at" gorm:"column:created_at"`
	UpdatedAt              string `json:"updated_at" gorm:"column:updated_at"`
}

// BorrowTokenInfo 解析borrowTokenInfo字段
type BorrowTokenInfo struct {
	BorrowFee  string `json:"borrowFee"`
	TokenLogo  string `json:"tokenLogo"`
	TokenName  string `json:"tokenName"`
	TokenPrice string `json:"tokenPrice"`
}

// LendTokenInfo 解析LendTokenInfo字段
type LendTokenInfo struct {
	LendFee    string `json:"lendFee"`
	TokenLogo  string `json:"tokenLogo"`
	TokenName  string `json:"tokenName"`
	TokenPrice string `json:"tokenPrice"`
}
