package db

import (
	"gin-api/app/util/yaml"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var Db *gorm.DB

func init() {
	var err error

	conf := yaml.Conf().DataSource

	Db, err = gorm.Open(conf.Connection, conf.Username+":"+conf.Password+"@tcp("+conf.Host+")/"+conf.Database)

	if err != nil {
		log.Fatal("DataBase Connection Error:" + err.Error())
	} else {
		log.Print("DataBase Connection Succeeded")
	}
}
