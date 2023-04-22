package server

import (
	"context"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Engine struct {
	SQL           sql.Service
	Car           Car.Service
	CarController Car.Controller
	logger        logger.Logger
}

func NewEngine(db *gorm.DB, logger logger.Logger) *Engine {
	ctx := context.TODO()
	/*topicARN := viper.GetString("")
	urlBucket := viper.GetString("")
	awsRegion := viper.GetString("")

	//httpClient := &http.Client{}*/

	sqlClient := sql.NewClient(ctx)
	sqlService := sql.NewService(sqlClient, logger)

	carRepository := car.NewClient(db)
	carService := car.NewService(carRepository, logger)
	carController := car.NewController(carService, logger)

	return &Engine{
		SQL:           sqlService,
		Car:           carService,
		CarController: carController,

		logger: logger,
	}
}

func (e *Engine) Run() error {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.RedirectTrailingSlash = true
	r.Use(gin.Recovery())

	e.registerRoutes(&r.RouterGroup)

	//port := viper.GetString("server.port")
	return r.Run(port)
}

func (e *Engine) registerRoutes(rg *gin.RouterGroup) {

	// Car
	c := rg.Group("/car")
	c.GET("/:car", e.CarController.GetApplicationByID)
	c.POST("/", e.carController.CreateCar)
	c.DELETE("/:car", e.carController.DeleteCar)
	c.PATCH("/status/:car", e.CarController.ChangeStatus)

	u := rg.Group("/user")
	u.GET("/:car", e.userController.GetApplicationByID)
	u.POST("/", e.userController.CreateCar)
	u.DELETE("/:car", e.userController.DeleteCar)
	u.PATCH("/status/:car", e.CarController.ChangeStatus)

	cp := rg.Group("/carPayment")
	cp.GET("/:car", e.carPaymentController.GetApplicationByID)
	cp.POST("/", e.carPaymentController.UpdatePayment)
	cp.DELETE("/:car", e.carPaymentController.DeletePayment)
	cp.PATCH("/status/:car", e.carPaymentController.ChangeStatus)

	wP := rg.Group("/washPayment")
	wp.GET("/:car", e.washPaymentController.GetApplicationByID)
	wp.POST("/", e.washPaymentController.UpdatePayment)
	wp.DELETE("/:car", e.washPaymentController.DeletePayment)
	wp.PATCH("/status/:car", e.washPaymentController.ChangeStatus)

	/*wP := rg.Group("/locate")
	wp.GET("/:car", e.washPaymentController.GetApplicationByID)
	wp.POST("/", e.washPaymentController.UpdatePayment)
	wp.DELETE("/:car", e.washPaymentController.DeletePayment)
	wp.PATCH("/status/:car", e.washPaymentController.ChangeStatus)*/

}
