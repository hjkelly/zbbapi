package v1

import "github.com/julienschmidt/httprouter"

func RegisterHandlers(router *httprouter.Router) {
	router.GET("/v1/categories", listCategories)
	router.POST("/v1/categories", createCategory)
	router.GET("/v1/categories/:id", retrieveCategory)
	router.PUT("/v1/categories/:id", updateCategory)
	router.DELETE("/v1/categories/:id", deleteCategory)
}
