package routes

import (
	"net/http"

	middleware "github.com/RolvinNoronha/fileupload-backend/internal/middlewares"
	"github.com/RolvinNoronha/fileupload-backend/internal/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func NewRouter(db *gorm.DB) http.Handler {
	g := gin.Default()

	v1 := g.Group("/api/v1");

	userRoutes := v1.Group("/user")
	{
		userRepo := user.NewRepository(db);
		userService := user.NewService(userRepo);
		userHandler := user.NewHandler(userService);

		userRoutes.POST("/register", userHandler.CreateUser);
		userRoutes.POST("/login", userHandler.Login);
	}

	authGroup := v1.Group("/files");
	authGroup.Use(middleware.AuthMiddleWare());


	return g;
}
