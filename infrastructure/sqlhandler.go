package infrastructure

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	// PostgreSQL driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/maaaaakoto35/PostUpAPI/interfaces/database"
)

// SQLHandler this struct has gorn.DB.
type SQLHandler struct {
	Conn *gorm.DB
}

// NewMySQLDb this func is initializing MySQL db.
func NewMySQLDb() database.SQLHandler {

	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABAE"),
	)

	conn, err := gorm.Open(os.Getenv("DB_DRIVER"), connectionString)
	if err != nil {
		panic(err)
	}

	sqlHandler := new(SQLHandler)
	sqlHandler.Conn = conn

	return sqlHandler
}

// Find this func is selecting somw rows.
func (handler *SQLHandler) Find(out interface{}, where ...interface{}) *gorm.DB {
	return handler.Conn.Find(out, where...)
}

// Count this func is a number of recordes.
func (handler *SQLHandler) Count(out interface{}, where ...interface{}) (int, error) {
	var count int = 0
	if err := handler.Conn.Find(out, where...).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// Create this func is inserting value.
func (handler *SQLHandler) Create(value interface{}) *gorm.DB {
	return handler.Conn.Create(value)
}

// Save this func is updating a row.
func (handler *SQLHandler) Save(value interface{}) *gorm.DB {
	return handler.Conn.Save(value)
}

// SaveValue this func is updating some columns.
func (handler *SQLHandler) SaveValue(in interface{}, set string, value string) *gorm.DB {
	return handler.Conn.Model(in).Update(set, value)
}

// Delete this func is deleting a row.
func (handler *SQLHandler) Delete(value interface{}) *gorm.DB {
	return handler.Conn.Delete(value)
}

// Delete this func is deleting a row.
func (handler *SQLHandler) Order(order string) *gorm.DB {
	return handler.Conn.Order(order)
}
