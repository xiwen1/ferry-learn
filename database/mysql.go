package database

import (
	orm "ferry-learn/global"
	"log"

	"github.com/jinzhu/gorm"
)

type Mysql struct{}

func (m *Mysql) Setup() {
	var err error
	orm.MysqlConn = m.GetConn()
	orm.Eloquent, err = m.Open("mysql", orm.MysqlConn)
	if err != nil {
		log.Fatalf("mysql connect error %v", err)
	} else {
		log.Println("mysql connect successfully!")
	}

	if orm.Eloquent.Error != nil {
		log.Fatalf("database error %v", orm.Eloquent.Error)
	}
	
	//todo: 配置orm的设置信息
}

func (m *Mysql) Open(dbType string, conn string) (*gorm.DB, error) {
	return gorm.Open(dbType, conn)
}

func (m *Mysql) GetConn() string {
	//todo: add config to avoid hard code
	return "ferry:zkw030813@tcp(127.0.0.1:3306)/ferry?charset=utf8&parseTime=True&loc=Local&timeout=10000ms"
}