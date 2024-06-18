package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sergeysergeevru/go-blog-service/internal/app"
	"github.com/sergeysergeevru/go-blog-service/internal/config"
	"github.com/sergeysergeevru/go-blog-service/internal/controllers/oapi"
	"github.com/sergeysergeevru/go-blog-service/internal/storage"
)

func CreateHandler(cfg *config.Cfg) *gin.Engine {
	router := gin.Default()
	blogService := app.NewBlogService(storage.NewPostInMemory())
	postController := NewPost(blogService)
	strictHandler := oapi.NewStrictHandler(postController, nil)
	oapi.RegisterHandlersWithOptions(router, strictHandler, oapi.GinServerOptions{BaseURL: cfg.Server.Prefix})

	return router
}
