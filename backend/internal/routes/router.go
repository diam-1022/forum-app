package routes

import (
    "forum-app/backend/internal/handlers"
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
)

func Setup(handler *handlers.Handler) *gin.Engine {
    router := gin.Default()
    router.Use(cors.New(cors.Config{
        AllowOrigins: []string{"*"},
        AllowMethods: []string{"GET", "POST", "OPTIONS"},
        AllowHeaders: []string{"Origin", "Content-Type"},
    }))

    api := router.Group("/api")
    {
        api.GET("/health", handler.Health)
        api.GET("/boards", handler.ListBoards)
        api.GET("/topics", handler.ListTopics)
        api.GET("/posts", handler.ListPosts)
        api.GET("/posts/:id", handler.GetPost)
        api.POST("/posts", handler.CreatePost)
        api.POST("/posts/:id/comments", handler.CreateComment)
    }

    return router
}

