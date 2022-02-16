package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"rshop/app/product/service/internal/conf"
)

type Dao struct {
	db  *gorm.DB
	log *log.Logger
}

var ProviderSet = wire.NewSet(NewDB, NewDao, NewProductRepo)

func NewDao(db *gorm.DB, log *log.Logger) *Dao {
	return &Dao{db: db, log: log}
}

func NewDB(conf *conf.Data) *gorm.DB {
	db, err := gorm.Open(mysql.Open(conf.Database.Source), &gorm.Config{})
	db.Debug()
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	if err := db.AutoMigrate(&Product{}); err != nil {
		log.Fatal(err)
	}
	return db

}
