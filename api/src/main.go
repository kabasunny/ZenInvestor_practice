package main

import (
    "fmt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "os"
    "time"
)

type User struct {
    ID        uint `gorm:"primaryKey"`
    Username  string `gorm:"unique;not null"`
    Email     string `gorm:"unique;not null"`
    Password  string `gorm:"not null"`
    CreatedAt time.Time
}

func main() {
    dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        os.Getenv("MYSQL_USER"),
        os.Getenv("MYSQL_PASSWORD"),
        os.Getenv("MYSQL_HOST"),
        os.Getenv("MYSQL_DB"))

    var db *gorm.DB
    var err error

    // 繝・・繧ｿ繝吶・繧ｹ謗･邯壹・繝ｪ繝医Λ繧､
    for i := 0; i < 10; i++ {
        db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
        if err == nil {
            break
        }
        fmt.Println("Database connection failed. Retrying in 5 seconds...")
        time.Sleep(5 * time.Second)
    }

    if err != nil {
        panic("failed to connect database")
    }

    db.AutoMigrate(&User{})
}