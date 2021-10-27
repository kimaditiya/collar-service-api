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

	dsn := "root:Ad200393@tcp(20.100.1.132:3306)/db_new_collar?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	//init process, parsing db connection
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	// userByEmail, err := userRepository.FindByEmail("arkans@gandul.com")

	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// if userByEmail.ID == 0 {
	// 	fmt.Println("User Not Found")
	// } else {

	// 	fmt.Println(userByEmail.FullName)
	// }

	// input := user.LoginInput{
	// 	NIK:      "1234567",
	// 	Password: "password",
	// }

	// user, err := userService.Login(input)
	// if err != nil {
	// 	fmt.Println("Terjadi Kesalahan")
	// 	fmt.Println(err.Error())
	// }

	// fmt.Println(user.NIK)

	//routing
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()

	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)

	router.Run()

}
