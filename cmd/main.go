package main

import (
	apiv1 "github.com/sulo1337/cleanarch-go/internal/api/v1"
	"github.com/sulo1337/cleanarch-go/internal/service"
	"github.com/sulo1337/cleanarch-go/internal/store"
	appLogger "github.com/sulo1337/cleanarch-go/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func main() {
	logger := appLogger.NewColorLogger(appLogger.LoggingOptions{
		Level:      appLogger.Info,
		ShowTime:   true,
		ShowCaller: true,
	})

	gormLoggerAdapter := appLogger.NewGormLoggerAdapter(logger)

	dsn := "root:changeme@tcp(127.0.0.1:3306)/cleanarch"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: gormLoggerAdapter,
	})
	if err != nil {
		panic(err)
	}

	logger.Errorf("This is an error message %v", time.Now())
	logger.Warnf("This is a warning message %v", time.Now())
	logger.Infof("This is an info message %v", time.Now())
	logger.Debugf("This is a debug message %v", time.Now())

	userStore := store.NewUserStore(db)
	postStore := store.NewPostStore(db)
	store := store.NewStore(userStore, postStore)

	userService := service.NewUserService(store)
	postService := service.NewPostService(store)
	service := service.NewService(userService, postService)

	api := apiv1.NewAPI(service)

	http.ListenAndServe(":8080", api.GetBaseRouter())
}
