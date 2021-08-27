package infrastructure

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SqlHandler struct {
	DB *gorm.DB
}

func NewSqlHandler() *SqlHandler {
	dsn := "gorm:gorm@tcp(db:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.DB = db
	return sqlHandler
}
