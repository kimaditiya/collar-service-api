package main

import (
	"collar-service-api/handler"
	"collar-service-api/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:@tcp(127.0.0.1:3306)/db_new_collar?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	//init process, parsing db connection
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	//routing
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()

	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)

	router.Run()

	// userInput := user.RegisterUserInput{}
	// userInput.FullName = "Test From Service"
	// userInput.NIK = "23456"
	// userInput.Email = "test@smartfinance.co.id"
	// userInput.Password = "asdf1234"

	// userService.RegisterUserInput(userInput)

	// userRepository.Save(user)

	// for _, users := range users {
	// 	fmt.Println(users.NIK)
	// 	fmt.Println(users.FullName)
	// 	fmt.Println(users.Email)
	// 	fmt.Println("=================")

	// }

	// router := gin.Default()
	// router.GET("/", handler)
	// router.Run()

}

// func handler(c *gin.Context) {
// 	dsn := "root:@tcp(127.0.0.1:3306)/db_new_collar?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	fmt.Println("Connect to database establish successfully")

// 	var users []user.User

// 	db.Find(&users)

// 	c.JSON(http.StatusOK, users)
// }
