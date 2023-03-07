package main

import (
	_storeHttpDelivery "e-commerce-store/store/delivery/http"
	_storeRepository "e-commerce-store/store/repository"
	_storeUsecase "e-commerce-store/store/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	// "github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"

	"github.com/labstack/gommon/log"
)

var (
	e *echo.Echo
)

func init() {
	viper.SetConfigFile(`config.json`)
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Printf("Service RUN on DEBUG mode")
	}

}
func main() {
	// Establish data base connection
	// dbHost := viper.GetString(`database.host`)
	// dbPort := viper.GetString(`database.port`)
	// dbUser := viper.GetString(`database.user`)
	// dbPass := viper.GetString(`database.pass`)
	// dbName := viper.GetString(`database.name`)
	dsn := "root:@tcp(localhost:3306)/e-store?charset=utf8mb4"
	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", dbHost, dbUser, dbPass, dbName, dbPort, "require", "Asia/Kolkata")
	dbConn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error("err in database connection message", err)
	}
	e := echo.New()
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())
	// e.Use(middleware.CORS())

	storeRepo := _storeRepository.NewMysqlStoreRepository(dbConn)
	storeUsecase := _storeUsecase.NewStoreUseCase(storeRepo)
	_storeHttpDelivery.NewStoreHandler(e, storeUsecase)

	log.Fatal(e.Start(":" + viper.GetString(`APPLICATION_PORT`)))

}
