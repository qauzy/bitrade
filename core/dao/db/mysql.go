package db

import (
	"bitrade/core/config"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DB struct {
	w *gorm.DB
	r *gorm.DB
}

// TybDbConf tydb
type Config struct {
	Addr        string `ini:"Addr"`
	LogEnable   bool   `ini:"LogEnable"`
	MaxConnect  int    `ini:"MaxConnect"`
	IdleConnect int    `ini:"IdleConnect"`
	MaxLifeTime int    `ini:"MaxLifeTime"`
}

func NewDB() (db *DB) {
	dbw := NewDBWrite()
	dbr := NewDBRead()

	db = &DB{
		w: dbw,
		r: dbr,
	}
	return
}

func (db *DB) DBWrite() *gorm.DB {
	return db.w
}

func (db *DB) DBRead() *gorm.DB {
	return db.r
}

func (db *DB) Close() {
	db.w.Close()
	db.r.Close()
}

//初始化写的数据库
func NewDBWrite() *gorm.DB {
	var conf Config
	err := config.GetConfig().Section("DbConf").MapTo(&conf)
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open("mysql", conf.Addr)
	if err != nil {
		tmpStr := fmt.Sprintf("Connet the dao=%s , err: %v", conf.Addr, err)
		panic(tmpStr)
	}

	db.DB().SetMaxOpenConns(conf.MaxConnect)
	db.DB().SetMaxIdleConns(conf.IdleConnect)
	db.DB().SetConnMaxLifetime(time.Duration(conf.MaxLifeTime) * time.Second)

	if err = db.DB().Ping(); err != nil {
		panic(err)
	}

	db.LogMode(conf.LogEnable)
	db.SingularTable(true)

	return db
}

func NewDBRead() *gorm.DB {

	var conf Config
	err := config.GetConfig().Section("ReDbConf").MapTo(&conf)
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open("mysql", conf.Addr)
	if err != nil {
		tmpStr := fmt.Sprintf("Connet the dao=%s , err: %v", conf.Addr, err)
		panic(tmpStr)
	}

	db.DB().SetMaxOpenConns(conf.MaxConnect)
	db.DB().SetMaxIdleConns(conf.IdleConnect)
	db.DB().SetConnMaxLifetime(time.Duration(conf.MaxLifeTime) * time.Second)

	if err = db.DB().Ping(); err != nil {
		panic(err)
	}

	db.LogMode(conf.LogEnable)
	db.SingularTable(true)

	return db
}
