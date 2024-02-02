package rest

import (
	_ "example/docs"
	"example/internal/app"

	"github.com/gin-gonic/gin"
	swaggerfile "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	productService app.ProductService
	userService    app.UserService
}

func NewHandler(productService app.ProductService, userService app.UserService) *Handler {
	return &Handler{
		productService: productService,
		userService:    userService,
	}
}

func (h *Handler) InitRouters() *gin.Engine {
	router := gin.Default()

	authGroup := router.Group("/auth")

	{
		authGroup.POST("/login", h.LoginUserHandler)
		authGroup.POST("/sign-up", h.SignUpUserHandler)
	}

	apiGroup := router.Group("/api", h.userIdentity)
	{
		userGroup := apiGroup.Group("/users")
		{
			userGroup.GET("/", h.GetPagesUser)
		}

		productGroup := apiGroup.Group("/products")
		{
			productGroup.POST("", h.CreateProduct)
			productGroup.GET("/:id", h.GetProductByID)
			productGroup.GET("/page", h.GetPagesProduct)
			productGroup.DELETE("/:id", h.DeleteProductByID)
			productGroup.PUT("/:id", h.UpdateProductByID)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfile.Handler))
	return router
}
