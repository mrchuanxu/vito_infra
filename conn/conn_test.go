package conn_test

import (
	"context"
	"github.com/vito_infra/util"
	"google.golang.org/grpc/metadata"
	"testing"

	"github.com/stretchr/testify/assert"
	conn "github.com/vito_infra/conn"
)

func Test_MySqlConn(t *testing.T) {
	ctx := metadata.NewIncomingContext(context.Background(),metadata.MD{})
	ctx = util.SetDBCodeCtx(ctx,"myscrm_jmadmin")
	db, err := conn.MySQLConn(ctx)
	assert.Nil(t, err)
	err = db.DB().Ping()
	assert.Nil(t, err)
}


func Test_MysqlParamsConn(t *testing.T){
	ctx := metadata.NewIncomingContext(context.Background(),metadata.MD{})
	ctx = util.SetDBCodeCtx(ctx,"myscrm_jmadmin")
	db, err := conn.MysqlParamsConn(ctx,"jinmao_cn:2Nc^JeNS7Jjm0755@tcp(pc-2zet0b02gcf3ylb32-out.mysql.polardb.rds.aliyuncs.com:3306)/")
	assert.Nil(t, err)
	err = db.DB().Ping()
	assert.Nil(t, err)
}

func Test_RedisParamsConn(t *testing.T){
	redisC,err := conn.RedisParamsConn(context.Background(),"127.0.0.1:6379","")
	assert.Nil(t, err)
	cmd := redisC.Ping()
	v,err := cmd.Result()
	t.Log(v)
	assert.Nil(t, err)
}