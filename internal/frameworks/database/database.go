package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"storage/internal/domain/entities"
)

func NewSession() *gorm.DB {
	dns := getDNS()
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database - err %v", err)
	}

	autoMigrate(db)

	return db
}

func autoMigrate(db *gorm.DB) {
	_ = db.AutoMigrate(entities.Product{}, entities.Audit{})
}

func getDNS() string {
	var (
		user        = os.Getenv("DB_USER")
		pass        = os.Getenv("DB_PASS")
		host        = os.Getenv("DB_HOST")
		port        = os.Getenv("DB_PORT")
		name        = os.Getenv("DB_NAME")
		dnsTemplate = os.Getenv("DB_CONN_STR") //"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	)
	return fmt.Sprintf(dnsTemplate, user, pass, host, port, name)
}
