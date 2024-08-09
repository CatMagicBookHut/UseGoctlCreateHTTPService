package model

import (
	"crud_api/utils"
	"fmt"
	"os"
	"strconv"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"gorm.io/driver/mysql"
)

var db *gorm.DB
var err error

func InitDb() {
	dbHost := utils.DbHost
	if mysqlHost := os.Getenv("MYSQL_HOST"); mysqlHost != "" {
		dbHost = mysqlHost
	}

	dbInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassword,
		dbHost,
		strconv.Itoa(utils.DbPort),
		utils.DbName,
	)
	db, err = gorm.Open(mysql.Open(dbInfo), &gorm.Config{
		// gorm日志模式:silent
		Logger: logger.Default.LogMode(logger.Silent),
		// 外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
		// 禁用默认事务
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			// 单数表名
			SingularTable: true,
		},
	})

	if err != nil {
		fmt.Println("DataBase connect failed.", err)
		os.Exit(1)
	}

	// 自动迁移数据
	db.AutoMigrate(&Cat{})

	sqlDB, _ := db.DB()
	// 设置连接池最大闲置连接数量
	sqlDB.SetMaxIdleConns(100)

	// 设置数据库与最大连接量
	sqlDB.SetMaxOpenConns(1000)

	// 设置连接最大可复用时间
	sqlDB.SetConnMaxLifetime(30 * time.Second)
}
