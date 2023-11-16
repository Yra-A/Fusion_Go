package db

import (
	"fmt"
	"github.com/Yra-A/Fusion_Go/pkg/constants"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormopentracing "gorm.io/plugin/opentracing"
)

var DB *gorm.DB

// Init init DB
func Init() {
	var err error
	DB, err = gorm.Open(mysql.Open(constants.MySQLDefaultDSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		fmt.Println(err)
	}

	if err = DB.Use(gormopentracing.New()); err != nil {
		fmt.Println(err)
	}

	err = DB.AutoMigrate(&UserProfileInfo{}, &Authentication{}, &Honors{})
	if err != nil {
		fmt.Println(err)
	}
}
