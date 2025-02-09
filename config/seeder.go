package config

import (
	"fmt"
	"inventory-book/models"
)

func SeedBooks() {
	var count int64
	DB.Model(&models.Book{}).Count(&count)

	if count == 0 {
		books := []models.Book{
			{Title: "Golang for Beginners", Author: "John Doe", Stock: 301},
			{Title: "Advanced Golang", Author: "Jane Smith", Stock: 2},
			{Title: "Web Development with Golang", Author: "Alice Johnson", Stock: 11},
			{Title: "Concurrency in Go", Author: "Rob Pike", Stock: 50},
			{Title: "Go Design Patterns", Author: "Mario Castro", Stock: 75},
			{Title: "Go Programming Blueprints", Author: "Mat Ryer", Stock: 100},
			{Title: "The Go Programming Language", Author: "Alan Donovan", Stock: 200},
			{Title: "Go in Practice", Author: "Matt Butcher", Stock: 150},
			{Title: "Go Web Programming", Author: "Sau Sheong Chang", Stock: 80},
			{Title: "Mastering Go", Author: "Mihalis Tsoukalos", Stock: 60},
		}

		for _, book := range books {
			result := DB.Create(&book)
			if result.Error != nil {
				fmt.Println("Failed added books:", result.Error)
			}
		}

		fmt.Println("Seeder successfully: added books!")
	} else {
		fmt.Println("Books table is not empty, skipping seeder!")
	}
}
