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
	Id                     int    `json:"id" gorm:"column:id;primaryKey"`
	PoolID                 int    `json:"pool_id" gorm:"column:pool_id;"`
	AutoLiquidateThreshold string `json:"autoLiquidateThreshold" gorm:"column:auto_liquidata_threshold;"`
	BorrowSupply           string `json:"borrowSupply" gorm:"column:borrow_supply;"`
	BorrowToken            string `json:"borrowToken" gorm:"column:pool_id;"`
	BorrowTokenInfo        string `json:"borrowTokenInfo" gorm:"column:borrow_token_info;"`
	EndTime                string `json:"endTime" gorm:"end_time;"`
	InterestRate           string `json:"interestRate" gorm:"column:interest_rate;"`
	JpCoin                 string `json:"jpCoin" gorm:"column:jp_coin;"`
	LendSupply             string `json:"lendSupply" gorm:"column:lend_supply;"`
	LendToken              string `json:"lendToken" gorm:"column:lend_token;"`
	LendTokenInfo          string `json:"lendTokenInfo" gorm:"column:lend_token_info;"`
	MartgageRate           string `json:"martgageRate" gorm:"column:martgage_rate;"`
	MaxSupply              string `json:"maxSupply" gorm:"column:max_supply;"`
	SettleTime             string `json:"settleTime" gorm:"column:settle_time;"`
	SpCoin                 string `json:"spCoin" gorm:"column:sp_coin;"`
	State                  string `json:"state" gorm:"column:state;"`
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
