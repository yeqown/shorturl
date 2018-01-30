package shorturl

import (
	"bufio"
	"encoding/json"
	"os"
)

var _instance Config

type RedisConfig struct {
	Addr     string `json:"Addr"`
	DB       int    `json:"DB"`
	Pwd      string `json:"Pwd"`
	PoolSize int    `json:"PoolSize"`
}
type Config struct {
	MySql string       `json:"MySql"`
	Redis *RedisConfig `json:"Redis"`
}

func GetInstance() *Config {
	return &_instance
}

func LoadConfig(filepath string) error {
	file, err := os.Open(filepath)
	defer file.Close()
	if err != nil {
		return err
	}

	data, _ := bufio.NewReader(file).ReadBytes(0)
	if err := json.Unmarshal(data, &_instance); err != nil {
		return err
	}
	return nil
}
