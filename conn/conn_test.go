package conn_test

import (
	"context"
	"github.com/vito_infra/util"
	"google.golang.org/grpc/metadata"
	"testing"

	"github.com/stretchr/testify/assert"
	conn "github.com/vito_infra/conn"
)

func Test_Conn(t *testing.T) {
	ctx := metadata.NewIncomingContext(context.Background(),metadata.MD{})
	ctx = util.SetDBCodeCtx(ctx,"myscrm_jmadmin")
	db, err := conn.MySQLConn(ctx)
	assert.Nil(t, err)
	err = db.DB().Ping()
	assert.Nil(t, err)
}
