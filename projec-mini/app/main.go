package main

import (
	_middleware "projectour/app/middlewares"

	userUseCase "projectour/businesses/users"
	userController "projectour/controllers/users"
	userRepo "projectour/drivers/repository/users"

	placeUseCase "projectour/businesses/places"
	placeController "projectour/controllers/places"
	placeRepo "projectour/drivers/repository/places"

	placeimageUseCase "projectour/businesses/placeimages"
	placeimageController "projectour/controllers/placeimages"
	placeimageRepo "projectour/drivers/repository/placeimages"

	savedplaceUseCase "projectour/businesses/savedplace"
	savedplaceController "projectour/controllers/savedplace"
	savedplaceRepo "projectour/drivers/repository/savedplace"

	ratedplaceUseCase "projectour/businesses/ratedplace"
	ratedplaceController "projectour/controllers/ratedplace"
	ratedplaceRepo "projectour/drivers/repository/ratedplace"

	"projectour/app/routes"
	"projectour/drivers/mysql"

	"log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	if viper.GetBool(`debug`) {
		log.Println("Service Run on Debug mode")
	}
}

func DBMigrate(DB *gorm.DB) {
	DB.AutoMigrate(&userRepo.User{})
	DB.AutoMigrate(&placeRepo.Place{})
	DB.AutoMigrate(&placeimageRepo.PlaceImage{})
	DB.AutoMigrate(&savedplaceRepo.SavedPlace{})
	DB.AutoMigrate(&ratedplaceRepo.RatedPlace{})
}

func main() {
	ConfigDB := mysql.ConfigDB{
		DB_Username: viper.GetString("database.user"),
		DB_Password: viper.GetString("database.pass"),
		DB_Host:     viper.GetString("database.host"),
		DB_Port:     viper.GetString("database.port"),
		DB_Database: viper.GetString("database.name"),
	}

	configJWT := _middleware.ConfigJWT{
		SecretJWT:       viper.GetString("JWT.secretKey"),
		ExpiresDuration: viper.GetInt("JWT.expired_time"),
	}

	DB := ConfigDB.InitialDB()
	DBMigrate(DB)
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	e := echo.New()
	_middleware.Log(e)
	e.Use(middleware.CORS())
	userRepoInterface := userRepo.NewUserRepo(DB)
	userUseCaseInterface := userUseCase.NewUseCase(userRepoInterface, timeoutContext, &configJWT)
	userUseControllerInterface := userController.NewUserController(userUseCaseInterface)

	placeRepoInterface := placeRepo.NewPlaceRepo(DB)
	placeUseCaseInterface := placeUseCase.NewUseCase(placeRepoInterface, timeoutContext)
	placeUseControllerInterface := placeController.NewPlaceController(placeUseCaseInterface)

	placeimageRepoInterface := placeimageRepo.NewPlaceImageRepo(DB)
	placeimageUseCaseInterface := placeimageUseCase.NewUseCase(placeimageRepoInterface, timeoutContext)
	placeimageUseControllerInterface := placeimageController.NewPlaceImageController(placeimageUseCaseInterface)

	savedplaceRepoInterface := savedplaceRepo.NewSavedPlaceRepo(DB)
	savedplaceUseCaseInterface := savedplaceUseCase.NewUseCase(savedplaceRepoInterface, timeoutContext)
	savedplaceUseControllerInterface := savedplaceController.NewSavedPlaceController(savedplaceUseCaseInterface)

	ratedplaceRepoInterface := ratedplaceRepo.NewRatedPlaceRepo(DB)
	ratedplaceUseCaseInterface := ratedplaceUseCase.NewUseCase(ratedplaceRepoInterface, timeoutContext)
	ratedplaceUseControllerInterface := ratedplaceController.NewRatedPlaceController(ratedplaceUseCaseInterface)

	routesInit := routes.RouteControllerList{
		UserController:       *userUseControllerInterface,
		PlaceController:      *placeUseControllerInterface,
		PlaceImageController: *placeimageUseControllerInterface,
		SavedPlaceController: *savedplaceUseControllerInterface,
		RatedPlaceController: *ratedplaceUseControllerInterface,
		JWTMiddleware:        configJWT.Init(),
	}
	routesInit.RouteRegister(e)
	log.Fatal(e.Start(viper.GetString("server.address")))

}
