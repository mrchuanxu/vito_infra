package common_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/vito_infra/common"
	"testing"
)

func Test_GetSet(t *testing.T){
	_,err :=common.EtcdClientKV.Put(context.Background(),"redis_addr","jinmao_cn:2Nc^JeNS7Jjm0755@tcp(pc-2zet0b02gcf3ylb32-out.mysql.polardb.rds.aliyuncs.com:3306)/")
	assert.Nil(t, err)
	kvrsp,err := common.EtcdClientKV.Get(context.Background(),"dbConnStr")
	t.Log(string(kvrsp.Kvs[0].Value))

	assert.Nil(t, err)
}


func Test_GetFuck(t *testing.T){
	kvrsp,err := common.EtcdClientKV.Get(context.Background(),"hello")
	assert.Nil(t, err)
	t.Log(string(kvrsp.Kvs[0].Value))
}