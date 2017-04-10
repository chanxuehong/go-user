package userid

import (
	"github.com/chanxuehong/go-user/config"
)

// 集群环境下不能重复
func getSnowflakeWorkerId() (int, error) {
	return config.ConfigData.SnowflakeWorkerId, nil
}
