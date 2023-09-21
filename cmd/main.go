package main

import (
	apiv1 "github.com/sulo1337/cleanarch-go/internal/api/v1"
	"github.com/sulo1337/cleanarch-go/internal/service"
	"github.com/sulo1337/cleanarch-go/internal/store"
	"net/http"
)

func main() {
	//dsn := "root:changeme@tcp(127.0.0.1:3306)/cleanarch"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//if err != nil {
	//	panic(err)
	//}

	userStore := store.NewUserStore(nil)
	postStore := store.NewPostStore(nil)
	store := store.NewStore(userStore, postStore)

	userService := service.NewUserService(store)
	postService := service.NewPostService(store)
	service := service.NewService(userService, postService)

	api := apiv1.NewAPI(service)

	http.ListenAndServe(":8080", api.GetBaseRouter())
}
