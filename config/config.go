/**
 * @Author: lixiang
 * @Date: 2025-03-18 17:47
 * @Description: config.go
 */

package config

import (
	"github.com/BurntSushi/toml"
	"path/filepath"
	"runtime"
)

var Config config

type config struct {
	Mysql        MysqlConfig
	Redis        RedisConfig
	TestNet      TestNetConfig
	MainNet      MainNetConfig
	Token        TokenConfig
	Email        EmailConfig
	DefaultAdmin DefaultAdminConfig
	Threshold    ThresholdConfig
	Jwt          JwtConfig
	Env          EnvConfig
}
type EnvConfig struct {
	Port               string `toml:"port"`
	Version            string `toml:"version"`
	Protocol           string `toml:"protocol"`
	DomainName         string `toml:"domain_name"`
	TaskDuration       int64  `toml:"task_duration"`
	WssTimeoutDuration int64  `toml:"wss_timeout_duration"`
	TaskExtendDuration int64  `toml:"task_extend_duration"`
}

type ThresholdConfig struct {
	AstraLendTokenThresholdBnb string `toml:"astra_lend_pool_token_threshold_bnb"`
}

type EmailConfig struct {
	Username string   `toml:"username"`
	Pwd      string   `toml:"pwd"`
	Host     string   `toml:"host"`
	Port     string   `toml:"port"`
	From     string   `toml:"from"`
	Subject  string   `toml:"subject"`
	To       []string `toml:"to"`
	Cc       []string `toml:"cc"`
}

type DefaultAdminConfig struct {
	Username string `toml:"username"`
	Password string `toml:"password"`
}

type JwtConfig struct {
	SecretKey  string `toml:"secret_key"`
	ExpireTime int    `toml:"expire_time"` // duration, s
}

type TokenConfig struct {
	LogoUrl string `toml:"logo_url"`
}

type MysqlConfig struct {
	Address     string `toml:"address"`
	Port        string `toml:"port"`
	DbName      string `toml:"db_name"`
	UserName    string `toml:"user_name"`
	Password    string `toml:"password"`
	MaxOpenConn int    `toml:"max_open_conn"`
	MaxIdleConn int    `toml:"max_idle_conn"`
	MaxLifeTime int    `toml:"max_life_time"`
}

type TestNetConfig struct {
	ChainId              string `toml:"chain_id"`
	NetUrl               string `toml:"net_url"`
	AstraAddress         string `toml:"astra_address"`
	AstraLendPoolToken   string `toml:"astra_lend_pool_token"`
	AstraLendOracleToken string `toml:"astra_lend_oracle_token"`
}

type MainNetConfig struct {
	ChainId              string `toml:"chain_id"`
	NetUrl               string `toml:"net_url"`
	AstraAddress         string `toml:"astra_address"`
	AstraLendPoolToken   string `toml:"astra_lend_pool_token"`
	AstraLendOracleToken string `toml:"astra_lend_oracle_token"`
}

type RedisConfig struct {
	Address     string `toml:"address"`
	Port        string `toml:"port"`
	Db          int    `toml:"db"`
	Password    string `toml:"password"`
	MaxIdle     int    `toml:"max_idle"`
	MaxActive   int    `toml:"max_active"`
	IdleTimeout int    `toml:"idle_timeout"`
}

func init() {
	var path string
	_, fileName, _, ok := runtime.Caller(0)
	if ok {
		path = filepath.Dir(fileName)
	}
	if _, err := toml.DecodeFile(path+"/configV21.toml", &Config); err != nil {
		panic(err)
	}
}
