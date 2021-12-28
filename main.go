package main

import (
	"collar-service-api/assign"
	"collar-service-api/collectiongroup"
	"collar-service-api/customers"
	"collar-service-api/dkh"
	"collar-service-api/handler"
	"collar-service-api/listarea"
	"collar-service-api/listbranch"
	"collar-service-api/menu"
	"collar-service-api/reasons"
	"collar-service-api/user"
	"collar-service-api/wilayah"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	/**
		koneksi to mysql
		*dsn := "root:Ad200393@tcp(20.100.1.132:3306)/db_new_collar?charset=utf8mb4&parseTime=True&loc=Local" // prod
		*dsn := "root:Ad200393@tcp(20.100.1.32:3306)/db_new_collar?charset=utf8mb4&parseTime=True&loc=Local" // dev
	**/
	dsn := "root:Ad200393@tcp(20.100.1.32:3306)/db_new_collar?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	customerRepository := customers.NewRepository(db)
	customerService := customers.NewService(customerRepository)
	menusRepository := menu.NewRepository(db)
	menuService := menu.NewService(menusRepository)
	wilayahRepository := wilayah.NewRepository(db)
	wilayahService := wilayah.NewService(wilayahRepository)
	listAreaNewRepository := listarea.NewRepository(db)
	listAreaNewService := listarea.NewService(listAreaNewRepository)
	listBranchNewRepository := listbranch.NewRepository(db)
	listBranchNewService := listbranch.NewService(listBranchNewRepository)
	collectionGroupNewRepository := collectiongroup.NewRepository(db)
	collectionGroupNewService := collectiongroup.NewService(collectionGroupNewRepository)
	reasonsRepository := reasons.NewRepository(db)
	reasonsService := reasons.NewService(reasonsRepository)
	assignRepository := assign.NewRepository(db)
	assignService := assign.NewService(assignRepository)
	dkhRepository := dkh.NewRepository(db)
	dkhService := dkh.NewService(dkhRepository)

	/**
		*routing
	**/
	userHandler := handler.NewUserHandler(userService)
	customerHandler := handler.NewCustomerHandler(customerService)
	menuHandler := handler.NewMenuHandler(menuService)
	listAreaNewHandler := handler.NewListAreaHandler(listAreaNewService)
	listBranchNewHandler := handler.NewListBranchHandler(listBranchNewService)
	collectionGroupNewHandler := handler.NewCollectionGroupHandler(collectionGroupNewService)
	WilayahHandler := handler.ListWilayahHandler(wilayahService)
	ReasonsHandler := handler.ListReasonsHandler(reasonsService)
	AssignHandler := handler.ListAssignHandler(assignService)
	DkhHandler := handler.ListDkhHandler(dkhService)
	router := gin.Default()
	router.Use(cors.Default())
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/createUserRoles", userHandler.CreateUserRoles)
	api.GET("/getAllListUser", userHandler.GetAllListUser)
	api.GET("/customers", customerHandler.ListOfCustomer)
	api.POST("/customersByProduct", customerHandler.FindListOfCustomerByProduct)
	api.POST("/customersByOvdDays", customerHandler.FindListOfCustomerByOvdDays)
	api.POST("/customersByDueDate", customerHandler.FindListOfCustomerByDueDate)
	api.POST("/addMenu", menuHandler.MenuNew)
	api.GET("/getAllMenu", menuHandler.GetAll)
	api.GET("/getAllMenuParent", menuHandler.GetAllMenuParent)
	api.GET("/getAllListArea", listAreaNewHandler.GetAllListArea)
	api.GET("/getAllListBranch", listBranchNewHandler.GetAllListBranch)
	api.POST("/postListBranchByArea", listBranchNewHandler.FindListOfBranchByAreaID)
	api.GET("/getAllListBranchPos", listBranchNewHandler.GetAllListBranchPos)
	api.POST("/postFindPostByPostID", listBranchNewHandler.FindPostByBranch)
	api.GET("/getAllListPrivelege", collectionGroupNewHandler.ListOfPrivelege)
	api.GET("/getAllListCollectionGroup", collectionGroupNewHandler.ListCollectionGroup)
	api.POST("/postCreateCollectionGroup", collectionGroupNewHandler.CollectionGroupNew)
	api.POST("/postUpdateCollectionGroup", collectionGroupNewHandler.CollectionGroupUpdate)
	api.GET("/customersColl", customerHandler.ListOfCustomerColl)
	api.GET("/getAllListWilayah", WilayahHandler.GetAllListWilayah)
	api.GET("/getAllListWilayahAssign", WilayahHandler.GetAllListWilayahAssign)
	api.POST("/postcreatewilayahtagih", WilayahHandler.WilayahTagihNew)
	api.GET("/getAllListReasons", ReasonsHandler.GetAllListReasons)
	api.GET("/getAllListAssign", AssignHandler.GetAllListAssign)
	api.POST("/postcreateassign", AssignHandler.AssignNew)
	api.GET("/getAllListDkh", DkhHandler.GetAllListDkh)
	api.POST("/postcreatedkh", DkhHandler.DkhNew)

	router.Run()

}
