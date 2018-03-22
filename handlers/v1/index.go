package v1

import "github.com/julienschmidt/httprouter"

func RegisterHandlers(router *httprouter.Router) {
	router.GET("/v1/categories", listCategories)
	router.POST("/v1/categories", createCategory)
	router.GET("/v1/categories/:id", retrieveCategory)
}
