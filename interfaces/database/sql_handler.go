package database

import "github.com/jinzhu/gorm"

// SQLHandler this interface is connecting SQLHandler.
type SQLHandler interface {
	Find(interface{}, ...interface{}) *gorm.DB
	Count(interface{}, ...interface{}) (int, error)
	Create(interface{}) *gorm.DB
	Save(interface{}) *gorm.DB
	Delete(interface{}) *gorm.DB
}
