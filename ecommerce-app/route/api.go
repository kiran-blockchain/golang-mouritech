package route

import (
	"ecommerce-app/handler"
	"ecommerce-app/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

//RunAPI ->route setup
func RunAPI(address string) error {

	userHandler := handler.NewUserHandler()
	productHandler := handler.NewProductHandler()
	orderHandler := handler.NewOrderHandler()

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Welcome to Our Mini Ecommerce")
	})

	apiRoutes := r.Group("/api")
	userRoutes := apiRoutes.Group("/user")

	{
		userRoutes.POST("/register", userHandler.AddUser)//localhost:8090/api/user/register
		userRoutes.POST("/signin", userHandler.SignInUser)
	}

	userProtectedRoutes := apiRoutes.Group("/users", middleware.AuthorizeJWT())
	{
		userProtectedRoutes.GET("/", userHandler.GetAllUser)////localhost:8090/api/users
		userProtectedRoutes.GET("/:user", userHandler.GetUser)////localhost:8090/api/users/1
		userProtectedRoutes.GET("/:user/products", userHandler.GetProductOrdered)//localhost:8090/api/users/1/products
		userProtectedRoutes.PUT("/:user", userHandler.UpdateUser)
		userProtectedRoutes.DELETE("/:user", userHandler.DeleteUser)
	}

	productRoutes := apiRoutes.Group("/products", middleware.AuthorizeJWT())
	{
		productRoutes.GET("/", productHandler.GetAllProduct)//localhost:8090/api/products
		productRoutes.GET("/:product", productHandler.GetProduct)
		productRoutes.POST("/", productHandler.AddProduct)
		productRoutes.PUT("/:product", productHandler.UpdateProduct)
		productRoutes.DELETE("/:product", productHandler.DeleteProduct)
	}

	orderRoutes := apiRoutes.Group("/order", middleware.AuthorizeJWT())
	{
		orderRoutes.POST("/product/:product/quantity/:quantity", orderHandler.OrderProduct)
	}

	fileRoutes := r.Group("/file")
	{
		fileRoutes.POST("/single", handler.SingleFile)
		fileRoutes.POST("/multi", handler.MultipleFile)
	}

	return r.Run(address)

}
