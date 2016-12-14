package db

import (
		"github.com/tkc/GoogleAppEngineDeploy-Sandbox/src/conf"
		"github.com/jinzhu/gorm"
		_ "github.com/go-sql-driver/mysql"
		_ "google.golang.org/appengine/cloudsql"
)

func Connect() *gorm.DB {
		db, err := gorm.Open("mysql", conf.DnSetting)
		if err != nil {
				panic(err)
		}
		return db
}
