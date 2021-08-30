package conn_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	conn "github.com/vito_infra/conn"
)

func Test_Conn(t *testing.T) {
	db, err := conn.MySQLConn(context.Background())
	assert.Nil(t, err)
	err = db.DB().Ping()
	assert.Nil(t, err)
}
