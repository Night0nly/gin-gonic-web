package infrastructure

import (
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/mediadotech/FY2018_2H_fp_team_training_hung/gin-gonic-web/adapter/controller"
	"github.com/mediadotech/FY2018_2H_fp_team_training_hung/gin-gonic-web/adapter/persistent/repository"
	"github.com/mediadotech/FY2018_2H_fp_team_training_hung/gin-gonic-web/adapter/persistent/service"
	service2 "github.com/mediadotech/FY2018_2H_fp_team_training_hung/gin-gonic-web/infrastructure/service"
)

type Router struct {
	router *gin.Engine
}

func NewRouter() *Router {
	router := gin.Default()

	//CORS setup
	corsConfig := cors.DefaultConfig()
	router.Use(cors.New(corsConfig))

	return &Router{
		router: router,
	}
}

func (s *Router) Run() {
	s.router.Use(static.Serve("/", static.LocalFile("./frontend/dist", true)))

	collectorList := service.EmptyCollectorList()
	collectorList.Add(service2.NewGoogleDriveCollector())
	collectorList.Add(service2.NewChiasenhacCollector())
	songController := controller.NewSongController(repository.NewSongResultPool(*collectorList))

	api := s.router.Group("/api")

	api.GET("/base-data", func(context *gin.Context) {
		songController.GetAllSong(context)
	})

	api.GET("/song", func(context *gin.Context) {
		songController.SearchSongByName(context)
	})

	//Google
	googleLoginController := controller.NewGoogleLoginController()
	api.GET("/login/google", func(context *gin.Context) {
		googleLoginController.RedirectToGoogleProvider()
	})
	api.GET("/login/google/callback", func(context *gin.Context) {
		googleLoginController.HandleProviderCallBack()
	})

	s.router.Run(":8088")
}
