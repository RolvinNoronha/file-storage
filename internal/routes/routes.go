package routes

import (
	"net/http"

	"github.com/RolvinNoronha/fileupload-backend/internal/file"
	"github.com/RolvinNoronha/fileupload-backend/internal/folder"
	middleware "github.com/RolvinNoronha/fileupload-backend/internal/middlewares"
	"github.com/RolvinNoronha/fileupload-backend/internal/user"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB, client *s3.Client) http.Handler {
	g := gin.Default()

	v1 := g.Group("/api/v1")

	userRoutes := v1.Group("/user")
	{
		userRepo := user.NewRepository(db)
		userService := user.NewService(userRepo)
		userHandler := user.NewHandler(userService)

		userRoutes.POST("/register", userHandler.CreateUser)
		userRoutes.POST("/login", userHandler.Login)
	}

	// authGroup := v1.Group("/files");
	// authGroup.Use(middleware.AuthMiddleWare());

	fileRoutes := v1.Group("/file")
	fileRoutes.Use(middleware.AuthMiddleWare())
	{
		fileRepo := file.NewRepository(db)
		fileService := file.NewService(fileRepo, client)
		fileHandler := file.NewHandler(fileService)

		fileRoutes.POST("/create", fileHandler.CreateFile)
		fileRoutes.GET("/files/:userId", fileHandler.GetFileByUserID)
		fileRoutes.GET("/files/:userId/:folderId", fileHandler.GetFileByUserIDFolderID)
	}

	folderRoutes := v1.Group("/folder")
	folderRoutes.Use(middleware.AuthMiddleWare())
	{
		folderRepo := folder.NewRepository(db)
		folderService := folder.NewService(folderRepo)
		folderHandler := folder.NewHandler(folderService)

		folderRoutes.POST("/create", folderHandler.CreateFolder)
		folderRoutes.GET("/folders/:userId", folderHandler.GetFoldersByUserID)
	}

	return g
}
