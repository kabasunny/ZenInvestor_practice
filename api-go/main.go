package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"size:100"`
	Email    string `gorm:"uniqueIndex;size:100"`
	Password string `gorm:"size:255"`
}

func main() {
	router := gin.Default()
	dsn := "zeninvestor:zenpass@tcp(localhost:3306)/zeninv?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	router.GET("/migration", func(c *gin.Context) {
		if db.Migrator().HasTable(&User{}) {
			c.JSON(200, gin.H{"message": "User table already exists. Migration skipped."})
			return
		}
		err := db.AutoMigrate(&User{})
		if err != nil {
			c.JSON(500, gin.H{"error": "failed to migrate database"})
			return
		}
		users := []User{
			{Name: "Alice", Email: "alice@example.com", Password: "password1"},
			{Name: "Bob", Email: "bob@example.com", Password: "password2"},
			{Name: "Charlie", Email: "charlie@example.com", Password: "password3"},
		}
		for _, user := range users {
			result := db.Create(&user)
			if result.Error != nil {
				c.JSON(500, gin.H{"error": "Failed to insert user: " + result.Error.Error()})
				return
			}
		}
		log.Println("Database connected and User table migrated successfully. Users added.")
		c.JSON(200, gin.H{"message": "Migration completed successfully"})
	})

	router.Run(":8086")
}
