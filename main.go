package main

import (
	"GoPolyglot/libs/common"
	"GoPolyglot/libs/configs"
	"GoPolyglot/libs/dbs"
	"GoPolyglot/libs/logger"
	"GoPolyglot/router"
	"flag"
	_ "github.com/gin-gonic/gin"
	"log"
)

func init() {

	var (
		err   error
		envCl string
	)
	/* 获取命令行参数 */
	flag.StringVar(&envCl, "e", "default", "环境变量默认是default ")
	flag.Parse()

	common.CommandParameterAdd(
		"env", envCl,
	)

	/* 初始化配置 */
	if err = configs.InitConfigs(envCl); err != nil {
		log.Fatal(err.Error())
	}
	///* 初始化mysql */
	if err = dbs.InitMysql(); err != nil {
		log.Fatal(err.Error())
	}
	///* 初始化redis */
	//if err = dbs.InitRedis(); err != nil {
	//	log.Fatal(err.Error())
	//}
	/* 初始化logger */
	if err = logger.InitLogger(); err != nil {
		log.Fatal(err.Error())
	}

}

func main() {
	router.InitRouter()
}
