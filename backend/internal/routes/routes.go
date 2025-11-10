package routes

import (
	"net/http"

	"github.com/RolvinNoronha/fileupload-backend/internal/file"
	"github.com/RolvinNoronha/fileupload-backend/internal/folder"
	middleware "github.com/RolvinNoronha/fileupload-backend/internal/middlewares"
	"github.com/RolvinNoronha/fileupload-backend/internal/user"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/elastic/go-elasticsearch/v9"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(ps *gorm.DB, es *elasticsearch.Client, s3 *s3.Client) http.Handler {
	g := gin.Default()

	v1 := g.Group("/api/v1")

	userRoutes := v1.Group("/user")
	{
		userRepo := user.NewRepository(ps)
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
		fileRepo := file.NewRepository(ps, es)
		fileService := file.NewService(fileRepo, s3)
		fileHandler := file.NewHandler(fileService)

		fileRoutes.POST("/create", fileHandler.CreateFile)
		fileRoutes.GET("/files", fileHandler.GetFileByUserID)
		fileRoutes.GET("/files/:folderId", fileHandler.GetFileByUserIDFolderID)
		fileRoutes.GET("/url/:fileId", fileHandler.GetFileUrl)
		fileRoutes.GET("/search", fileHandler.SearchFilesHandler)
	}

	folderRoutes := v1.Group("/folder")
	folderRoutes.Use(middleware.AuthMiddleWare())
	{
		folderRepo := folder.NewRepository(ps)
		folderService := folder.NewService(folderRepo)
		folderHandler := folder.NewHandler(folderService)

		folderRoutes.POST("/create", folderHandler.CreateFolder)
		folderRoutes.GET("/folders", folderHandler.GetFolders)
		folderRoutes.GET("/folders/:folderId", folderHandler.GetFoldersByFolderId)
	}

	return g
}
