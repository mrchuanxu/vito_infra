package common_test

import (
	"context"
	"fmt"
	"runtime"
	"testing"

	"github.com/mrchuanxu/vito_infra/alg"
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


func Test_Stack(t *testing.T){
	TransStack := alg.Init()
	TransStack.Push(1)
	TransStack.Push(2)
	TransStack.Push(3)
	TransStack.Push("what????")

	fmt.Println(TransStack.Pop())

	fmt.Println(TransStack.Pop())

	fmt.Println(TransStack.Pop())

	fmt.Println(TransStack.Pop())
	runtime.GC()
}