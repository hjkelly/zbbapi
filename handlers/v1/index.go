package v1

import "github.com/julienschmidt/httprouter"

// RegisterHandlers links all current route handlers to the router provided.
func RegisterHandlers(router *httprouter.Router) {
	router.GET("/v1/categories", listCategories)
	router.POST("/v1/categories", createCategory)
	router.GET("/v1/categories/:id", retrieveCategory)
	router.PUT("/v1/categories/:id", updateCategory)
	router.DELETE("/v1/categories/:id", deleteCategory)

	router.GET("/v1/plans", listPlans)
	router.POST("/v1/plans", createPlan)
	router.GET("/v1/plans/:id", retrievePlan)
	router.PUT("/v1/plans/:id", updatePlan)
	router.DELETE("/v1/plans/:id", deletePlan)

	router.POST("/v1/plans/:id/conversions", createConversion)

	router.GET("/v1/budgets", listBudgets)
	router.POST("/v1/budgets", createBudget)
	router.GET("/v1/budgets/:id", retrieveBudget)
	router.PUT("/v1/budgets/:id", updateBudget)
	router.DELETE("/v1/budgets/:id", deleteBudget)

	router.GET("/v1/jobs", listJobs)
	router.POST("/v1/jobs", createJob)
	router.GET("/v1/jobs/:id", retrieveJob)
	router.PUT("/v1/jobs/:id", updateJob)
	router.DELETE("/v1/jobs/:id", deleteJob)
}
