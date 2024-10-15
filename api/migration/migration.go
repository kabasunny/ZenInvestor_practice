package main

import (
    ""fmt""
    ""gorm.io/driver/mysql""
    ""gorm.io/gorm""
    ""os""
    ""time""
)

type User struct {
    ID       uint   gorm:"primaryKey"
    Username string gorm:"unique;not null"
    Email    string gorm:"unique;not null"
    Password string gorm:"not null"
    CreatedAt time.Time
}

func main() {
    dsn := fmt.Sprintf(""%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local"", 
                      os.Getenv(""MYSQL_USER""), 
                      os.Getenv(""MYSQL_PASSWORD""), 
                      os.Getenv(""MYSQL_HOST""), 
                      os.Getenv(""MYSQL_DB""))
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic(""failed to connect database"")
    }

    db.AutoMigrate(&User{})
}