package dbhelper

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

func Create(db *gorm.DB, m interface{}) error {
	err := db.Create(m).Error
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func Save(db *gorm.DB, m interface{}) error {
	err := db.Save(m).Error
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func First(db *gorm.DB, m interface{}, filter ...interface{}) error {
	err := db.First(m, filter...).Error
	if err == gorm.ErrRecordNotFound {
		return err
	}
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func Pluck(db *gorm.DB, m interface{}, c string, v interface{}, filter ...interface{}) error {
	if len(filter) > 0 {
		db = db.Where(filter[0], filter[1:]...)
	}
	err := db.Model(m).Pluck(c, v).Error
	if err == gorm.ErrRecordNotFound {
		return err
	}
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func Update(db *gorm.DB, m interface{}) error {
	type attributer interface {
		Attributes() map[string]interface{}
		Refresh()
	}
	var err error
	if attr, ok := m.(attributer); ok {
		attributes := attr.Attributes()
		if len(attributes) == 0 {
			return nil
		}
		err = db.Model(m).Updates(attributes).Error
		if err != nil {
			return errors.WithStack(err)
		}
		attr.Refresh()
	} else {
		err = db.Updates(m).Error
		if err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}

func Count(db *gorm.DB, v interface{}, filter ...interface{}) error {
	if len(filter) > 0 {
		db = db.Where(filter[0], filter[1:]...)
	}
	err := db.Count(v).Error
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func Delete(db *gorm.DB, m interface{}, filter ...interface{}) error {
	err := db.Delete(m, filter...).Error
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func Find(db *gorm.DB, m interface{}, filter ...interface{}) error {
	err := db.Find(m, filter...).Error
	if err == gorm.ErrRecordNotFound {
		return err
	}
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func Scan(db *gorm.DB, v interface{}) error {
	err := db.Scan(v).Error
	if err == gorm.ErrRecordNotFound {
		return err
	}
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
