package main

import (
	"log"
	"startcare/user"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// func main() {
// 	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
// 	// 	dsn := "root:1n73GR!ya@tcp(127.0.0.1:3306)/gotrial?charset=utf8mb4&parseTime=True&loc=Local"
// 	// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 	// 	if err != nil {
// 	// 		log.Fatal(err.Error())
// 	// 	}

// 	// 	fmt.Println("Database connected")

// 	// 	var users []user.User
// 	// 	db.Find(&users)

// 	// 	for _, user := range users {
// 	// 		fmt.Println(user.Name)
// 	// 		fmt.Println(user.Email)
// 	// 		fmt.Println("================")
// 	// 	}
// 	router := gin.Default()
// 	router.GET("/handler", handler)
// 	router.Run()
// }

func main() {
	dsn := "root:1n73GR!ya@tcp(127.0.0.1:3306)/gotrial?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	userInput := user.RegisterUserInput{}
	userInput.Name = "save from service"
	userInput.Email = "dagum@mailinator.com"
	userInput.Occupation = "qa"
	userInput.Password = "password"

	userService.RegisterUser(userInput)

	// input
	// handler mapping input to struct
	// service mapping struct User
	// repository save struct User to db
	// db
}
