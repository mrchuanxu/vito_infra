package dbhelper

import (
	"context"

	"github.com/jinzhu/gorm"
)

//Register 注册数据库连接的回调和上下文信息，通常在拦截器中使用
func Register(ctx context.Context) context.Context {
	db := db(ctx, "db")
	if db == nil {
		return ctx
	}
	db = registerCallback(db)
	return context.WithValue(ctx, "db", db)
}

//DB 从context中获取数据库连接实例，通常在repository层使用
func DB(ctx context.Context) *gorm.DB {
	if db := db(ctx, "trans"); db != nil {
		return db
	}

	return db(ctx, "db")
}

//Begin 开启数据库的事务，通常在service中使用
func Begin(ctx context.Context) context.Context {
	db := db(ctx, "db")
	trans := db.Begin()
	return context.WithValue(ctx, "trans", trans)
}

//Commit 提交数据库的事务，通常在service中使用
func Commit(ctx context.Context) context.Context {
	trans := db(ctx, "trans")
	trans.Commit()
	return context.WithValue(ctx, "trans", nil)
}

//Rollback 回滚数据库的事务，通常在service中使用
func Rollback(ctx context.Context) context.Context {
	trans := db(ctx, "trans")
	trans.Rollback()
	return context.WithValue(ctx, "trans", nil)
}

//Call 注册配置库连接池实例的回调，可注册事件处理，修改连接配置等
func Call(ctx context.Context, handle func(*gorm.DB) *gorm.DB) context.Context {
	db := db(ctx, "db")
	db = handle(db)
	return context.WithValue(ctx, "db", db)
}
