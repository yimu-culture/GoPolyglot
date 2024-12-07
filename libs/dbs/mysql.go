package dbs

import (
	"GoPolyglot/libs/configs"
	"fmt"
	"github.com/outreach-golang/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var GMysql map[string]*gorm.DB

func init() {
	GMysql = make(map[string]*gorm.DB)
}

func InitMysql() error {
	for _, v := range configs.GConfig.Databases.Mysql {
		cli, err := gorm.Open(
			mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&timeout=5s&loc=Asia%%2FShanghai",
				v.Username, v.Password, v.Address, v.Dbname)),
			&gorm.Config{
				QueryFields: true,
			},
		)
		if err != nil {
			return err
		}

		err = cli.Use(&logger.TracePlugin{})
		if err != nil {
			return err
		}

		db, _ := cli.DB()

		db.SetMaxOpenConns(v.MaxOpenConns)
		db.SetMaxIdleConns(10)
		db.SetConnMaxLifetime(time.Hour)

		GMysql[v.Asname] = cli
	}

	return nil

}
