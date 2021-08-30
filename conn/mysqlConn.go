package conn

import (
	_ "github.com/go-sql-driver/mysql" // 初始化mysql驱动
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

// GormConfigs gorm连接配置
type GormConfigs struct {
	// 数据库类型
	Provider     string
	ConnectStr   string
	MaxConns     int
	MaxIdleConns int
	// EnableLogger bool
	// Logger       logger.Logger
}

// GormConfig gorm配置
type GormConfig func(*GormConfigs)

// PrepareGorm 准备gorm连接
func PrepareGorm(confs ...GormConfig) (db *gorm.DB, err error) {
	arg := &GormConfigs{}
	// 优雅传参
	for _, conf := range confs {
		conf(arg)
	}
	db, err = gorm.Open(arg.Provider, arg.ConnectStr)
	if err != nil {
		err = errors.Wrapf(err, "Open DB has error")
		return
	}
	// Disable plular
	db.SingularTable(true)
	if arg.MaxConns > 0 && arg.MaxIdleConns > 0 {
		db.DB().SetMaxOpenConns(arg.MaxConns)
		db.DB().SetMaxIdleConns(arg.MaxIdleConns)
	}
	return
}

// Provider 数据库提供者
func Provider(provider string) GormConfig {
	return func(arg *GormConfigs) {
		arg.Provider = provider
	}
}

// ConnectStr 数据库地址
func ConnectStr(connectStr string) GormConfig {
	return func(arg *GormConfigs) {
		arg.ConnectStr = connectStr
	}
}

// MaxConns 最大连接数
func MaxConns(maxConns int) GormConfig {
	return func(arg *GormConfigs) {
		arg.MaxConns = maxConns
	}
}

// MaxIdleConns 最大id数
func MaxIdleConns(maxIdleConns int) GormConfig {
	return func(arg *GormConfigs) {
		arg.MaxIdleConns = maxIdleConns
	}
}
