package main

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "time"
)

type User struct {
    ID       uint   `gorm:"primaryKey"`
    Username string `gorm:"unique;not null"`
    Email    string `gorm:"unique;not null"`
    Password string `gorm:"not null"`
    CreatedAt time.Time
}

func main() {
    dsn := "user:password@tcp(mysql:3306)/mydb?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
    
    db.AutoMigrate(&User{}) 
}