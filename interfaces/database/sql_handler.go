package database

import "github.com/jinzhu/gorm"

// SQLHandler this interface is connecting SQLHandler.
type SQLHandler interface {
	Find(interface{}, ...interface{}) *gorm.DB
	Create(interface{}) *gorm.DB
	Save(interface{}) *gorm.DB
	SaveValue(interface{}, string, string) *gorm.DB
	Delete(interface{}) *gorm.DB
}
