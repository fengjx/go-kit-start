package rediscli

import (
	"github.com/redis/go-redis/v9"

	"github.com/fengjx/go-kit-start/common/config"
)

var cliMap = make(map[string]*redis.Client)
var defaultCli *redis.Client

func Init() {
	for key, c := range config.GetConfig().Redis {
		cli := redis.NewClient(&redis.Options{
			ClientName: key,
			Addr:       c.Addr,
			Password:   c.Password,
			DB:         c.DB,
		})
		cliMap[key] = cli
	}
	defaultCli = cliMap["default"]
}

func GetDefaultClient() *redis.Client {
	return defaultCli
}

func GetClient(name string) *redis.Client {
	return cliMap[name]
}
