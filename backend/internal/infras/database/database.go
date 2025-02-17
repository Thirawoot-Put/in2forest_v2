package database

import (
	"fmt"
	"log"
	"thirawoot/in2forest_shop_backend/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(env *config.ConfigEnv) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Bangkok",
		env.DbHost,
		env.DbUser,
		env.DbPwd,
		env.DbName,
		env.DbPort,
		env.DbSsl,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	migrate(db)

	return db
}

func migrate(db *gorm.DB) {
	err := db.AutoMigrate()
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}
