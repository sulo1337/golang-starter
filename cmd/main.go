package main

import (
	"github.com/sulo1337/cleanarch-go/internal/api"
	"github.com/sulo1337/cleanarch-go/internal/service"
	"github.com/sulo1337/cleanarch-go/internal/store"
	appLogger "github.com/sulo1337/cleanarch-go/pkg/logger"
	"go.uber.org/fx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log/slog"
	"net/http"
)

func main() {
	fx.New(
		fx.NopLogger,
		fx.Provide(
			appLogger.NewLogger,
			appLogger.NewGormLoggerAdapter,
			DB,
			store.NewUserStore,
			store.NewPostStore,
			service.NewPostService,
			service.NewUserService,
			api.NewAPI,
		),
		fx.Invoke(func(logger *slog.Logger, api *api.API) {
			logger.Info("server started", "port", "8080")
			err := http.ListenAndServe(":8080", api.GetBaseRouter())
			if err != nil {
				panic(err)
			}
		})).Run()
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
