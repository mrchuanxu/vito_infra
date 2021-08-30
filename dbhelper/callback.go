package dbhelper

import (
	"context"
	"time"

	"github.com/jinzhu/gorm"
)

func registerCallback(db *gorm.DB) *gorm.DB {
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	db.BlockGlobalUpdate(true)
	return db
}

// updateTimeStampForCreateCallback sets `CreateTime`, `ModifyTime` when creating.
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		now := time.Now()

		if createTimeField, ok := scope.FieldByName("CreatedOn"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(now)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(now)
			}
		}
	}
}

// updateTimeStampForUpdateCallback sets `ModifyTime` when updating.
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.Search.Omit("created_on", "created_by", "deleted_state", "deleted_by", "createdOn", "createdBy", "deletedState", "deletedBy")

		now := time.Now()

		if modifyTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
			if modifyTimeField.IsBlank {
				scope.SetColumn("ModifiedOn", now)
			}
		}
	}

}

func db(ctx context.Context, key string) *gorm.DB {
	db, ok := ctx.Value(key).(*gorm.DB)
	if ok {
		return db
	}

	return nil
}
