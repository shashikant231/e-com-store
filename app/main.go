package main

import (
	"e-commerce-store/domain"
	_storeHttpDelivery "e-commerce-store/store/delivery/http"
	_storeRepository "e-commerce-store/store/repository"
	_storeUsecase "e-commerce-store/store/usecase"
	"fmt"

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
	viper.SetConfigFile(`config.yml`)
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
	dsn := "root:@tcp(localhost:3306)/e-store?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", dbHost, dbUser, dbPass, dbName, dbPort, "require", "Asia/Kolkata")
	dbConn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error("err in database connection message", err)
	}
	sqlDB, err := dbConn.DB()
	if err != nil {
		panic("Failed to get database instance!")
	}
	defer sqlDB.Close()
	if err := sqlDB.Ping(); err != nil {
		panic("Failed to ping database!")
	}

	fmt.Println("Connected to database!")
	// Auto-migrate database
	err = dbConn.AutoMigrate(&domain.Category{})
	if err != nil {
		panic(err)
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
