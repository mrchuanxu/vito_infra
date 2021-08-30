// Package conn 数据库配置链接以及客户端
package conn

import (
	"context"
	"errors"

	"google.golang.org/grpc"

	"github.com/vito_infra/common"

	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"
	"github.com/vito_infra/config"
	"github.com/vito_infra/logger"
	"google.golang.org/grpc/metadata"
)

const (
	our2bTest     = "our2b_test"
	provider      = "mysql"
	redisTest     = "redis_addr"
	redisTestPass = "redis_pass"
	maxOpenConn   = 10
	maxIdleConn   = 5
)

const (
	dbParse = "?charset=utf8mb4&parseTime=True"
)

func DBHandleGrpcInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		return MySqlDBHandleGrpc(ctx, req, info, handler)
	}
}

// MySQLConn 获取mysql链接
func MySQLConn(ctx context.Context) (*gorm.DB, error) {
	kvRsp, err := common.EtcdClientKV.Get(ctx, config.GetString("dbConnStr", our2bTest))
	if err != nil {
		logger.TransLogger.Sugar().Errorf("MySQLConn has err[%v] with conn [%s]", err, our2bTest)
		return nil, err
	}
	dbName := getDBName(ctx)
	if dbName == "" {
		return nil, errors.New("sql object is nil")
	}
	dbClient, err := PrepareGorm(Provider(config.GetString("provider", provider)),
		ConnectStr(string(kvRsp.Kvs[0].Value)+dbName+dbParse),
		MaxConns(config.GetInt("maxOpenConn", maxOpenConn)),
		MaxIdleConns(config.GetInt("maxIdleConn", maxIdleConn)),
	)
	if err != nil {
		logger.TransLogger.Sugar().Errorf("MySQLConn  PrepareGorm has err[%v] with conn [%s]", err, our2bTest)
	}
	return dbClient, err
}

func getDBName(ctx context.Context) string {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		v := md.Get("db_code")
		if len(v) > 0 {
			return v[0]
		}
	}
	return ""
}

// RedisConn 获取redis客户端
func RedisConn(ctx context.Context) (*redis.Client, error) {
	kvAddrRsp, err := common.EtcdClientKV.Get(ctx, config.GetString("redisConn", redisTest))
	if err != nil {
		logger.TransLogger.Sugar().Errorf("RedisConn  has err[%v] with conn [%s]", err, redisTest)
		return nil, err
	}
	kvPasRsp, err := common.EtcdClientKV.Get(ctx, config.GetString("redisConnPass", redisTestPass))
	if err != nil {
		logger.TransLogger.Sugar().Errorf("RedisConn  has err[%v] with connPass [%s]", err, redisTestPass)
		return nil, err
	}
	redisClient, err := PrepareRedis(ctx,
		RedisAddr(string(kvAddrRsp.Kvs[0].Value)),
		RedisPassword(string(kvPasRsp.Kvs[0].Value)),
	)
	if err != nil {
		logger.TransLogger.Sugar().Errorf("RedisConn PrepareRedis  has err[%v]", err)
	}
	return redisClient, err
}

func MySqlDBHandleGrpc(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	db, err := MySQLConn(ctx)
	if err != nil {
		logger.TransLogger.Sugar().Errorf("MySqlDBHandleGrpc has err [%v]", err)
		return nil, errors.New("GetMysqlConn err")
	}
	defer db.Close()
	ctx = context.WithValue(ctx, "db", db)
	return handler(ctx, req)
}
