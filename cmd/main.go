package main

import (
	apiv1 "github.com/sulo1337/cleanarch-go/internal/api/v1"
	"github.com/sulo1337/cleanarch-go/internal/service"
	"github.com/sulo1337/cleanarch-go/internal/store"
	appLogger "github.com/sulo1337/cleanarch-go/pkg/logger"
	"go.uber.org/fx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

func main() {
	fx.New(
		fx.Provide(
			Logger,
			appLogger.NewGormLoggerAdapter,
			DB,
			store.NewUserStore,
			store.NewPostStore,
			service.NewPostService,
			service.NewUserService,
			apiv1.NewAPI,
		),
		fx.Invoke(func(logger appLogger.Logger, api *apiv1.API) {
			logger.Info("started server on port 8080")
			err := http.ListenAndServe(":8080", api.GetBaseRouter())
			if err != nil {
				logger.Fatal(err)
			}
		})).Run()
}

func Logger() appLogger.Logger {
	logger := appLogger.NewColorLogger(appLogger.LoggingOptions{
		Level:      appLogger.Info,
		ShowTime:   true,
		ShowCaller: true,
	})
	return logger
}

func DB(gormLoggerAdapter *appLogger.GormLoggerAdapter) *gorm.DB {
	dsn := "root:changeme@tcp(127.0.0.1:3306)/cleanarch"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: gormLoggerAdapter,
	})
	if err != nil {
		panic(err)
	}
	return db
}
