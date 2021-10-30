package main

import (
	"fmt"
	"log"
	"net/http"
	"startcare/auth"
	"startcare/campaign"
	"startcare/handler"
	"startcare/helper"
	"startcare/user"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
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
	campaignRepository := campaign.NewRepository(db)

	userService := user.NewService(userRepository)
	campaignService := campaign.NewService(campaignRepository)
	authService := auth.NewService()

	campaigns, _ := campaignService.FindCampaigns(2)
	fmt.Println(len(campaigns))
	// token, err := authService.ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo0fQ.5uLQPYy4C1VKJ9cr05xswjgH0gwk8uTSiGltkAdHhVM")
	// if err != nil {
	// 	fmt.Println("ERROR")
	// }

	// if token.Valid {
	// 	fmt.Println("VALID")
	// } else {
	// 	fmt.Println("INVALID")
	// }

	// fmt.Println(authService.GenerateToken(1001))

	// userService.SaveAvatar(1, "images/1-profile.png")

	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)

	router.Run()
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// Bearer hardcode
		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["user_id"].(float64))

		user, err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		c.Set("currentUser", user)
	}
}
