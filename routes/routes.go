package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	router.GET("/events", getEvents)
	router.POST("/events", createEvents)
	router.GET("/events/:id", getEvent)
	router.PUT("/events/:id", updateEvent)
}
