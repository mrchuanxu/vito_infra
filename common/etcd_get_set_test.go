package common_test

import (
	"context"
	"testing"

	"github.com/mrchuanxu/vito_infra/common"
	"github.com/stretchr/testify/assert"
)

func Test_GetSet(t *testing.T) {
	_, err := common.EtcdClientKV.Put(context.Background(), "redis_addr", "")
	assert.Nil(t, err)
	kvrsp, err := common.EtcdClientKV.Get(context.Background(), "dbConnStr")
	t.Log(string(kvrsp.Kvs[0].Value))

	assert.Nil(t, err)
}

func Test_GetFuck(t *testing.T) {
	kvrsp, err := common.EtcdClientKV.Get(context.Background(), "hello")
	assert.Nil(t, err)
	t.Log(string(kvrsp.Kvs[0].Value))
}
