package common

import (
	"context"
	"time"

	stdErr "github.com/pkg/errors"
	"github.com/vito_infra/config"

	"github.com/vito_infra/logger"
	"go.etcd.io/etcd/clientv3"
)

const (
	clientIP = "localhost:2379"
)

var (
	EtcdClientKV clientv3.KV
)

// GetClient 获取配置中心的
func init() {
	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints:         []string{config.GetString("etcd_client", clientIP)},
		DialTimeout:       5 * time.Second,
		DialKeepAliveTime: time.Duration(time.Minute),
	})
	if err != nil {
		err = stdErr.Wrapf(err, "GetClient has err[%v] with clientIP[%s]", err, clientIP)
		logger.TransLogger.Sugar().Errorf("setting is wrong with err[%v]", err)
	}
	EtcdClientKV = clientv3.NewKV(etcdClient)
}

func GetToken(ctx context.Context) string {
	kvRsp, err := EtcdClientKV.Get(ctx, config.GetString("tushare.token", ""))
	if err != nil {
		logger.TransLogger.Sugar().Panic("tushare token is empty")
	}
	return string(kvRsp.Kvs[0].Value)
}
